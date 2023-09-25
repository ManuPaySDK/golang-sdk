package manupay_client

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/parnurzeal/gorequest"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
)

// 生成账单
/*
	orderID 商户内部订单id，要求同一商户唯一
*/
func (client *ManuPayClient) PlaceUnifiedOrder(request UnifiedOrderRequest) (bool, UnifiedOrderResponse) {

	var urlResp UnifiedOrderResponse

	url := fmt.Sprintf("%s%s", client.Host, UNIFIEDORDER_PATH)

	//----------------校验额外参数(必须有如下九个字段)-------------------------
	legalParams := []string{"firstname", "lastname", "city", "address", "phone", "email", "country", "state", "postcode"}

	var res map[string]string
	err := json.Unmarshal([]byte(request.ExtParam), &res)
	if err != nil || len(res) < len(legalParams) {
		return false, UnifiedOrderResponse{}
	}
	var subset []string = lo.Without(legalParams, GetKeys(res)...)
	if len(subset) > 0 {
		return false, UnifiedOrderResponse{}
	}
	//--------------------------------------------------------------------
	if request.ExpiredTime <= 0 {
		request.ExpiredTime = 3600
	}

	//计算Body
	signForm := client.CalculateSign(request)

	//发送请求
	_, _, errs := gorequest.New().Post(url).Send(signForm.Body).EndStruct(&urlResp)
	if errs != nil {
		return false, UnifiedOrderResponse{}
	} else {
		return true, urlResp
	}
}

//---------------------------------------------------------------

type SignForm struct {
	Raw  string `json:"raw" structs:"raw"`   //签名的原始字符串
	Sign string `json:"sign" structs:"sign"` //计算的签名值
	Body string `json:"body" structs:"body"` //请求的post body
}

func (client *ManuPayClient) CalculateSign(request UnifiedOrderRequest) SignForm {

	var result2 SignForm

	//请求封装公共参数
	commonReq := CommonRequestInfo{
		MchNo: client.MchNo, //商户号
		//ReqTime: time.Now().UnixMilli(), //请求时间
	}

	//计算签名
	rawParams := structs.Map(request)
	commonParams := structs.Map(commonReq)
	maps.Copy(rawParams, commonParams)
	rawString, signVal := GenSign(rawParams, client.PrivateSecret)

	//1. 赋值: 原始字符串,签名
	result2.Raw = rawString
	result2.Sign = signVal

	//-----------------------------------------------------------
	//合并复制
	type UnifiedOrderRequestFinal struct {
		CommonRequestInfo
		UnifiedOrderRequest
	}
	commonReq.Sign = signVal
	result := UnifiedOrderRequestFinal{
		commonReq,
		request,
	}

	//构造请求body
	paramJSON, _ := json.Marshal(result)
	paramStr := string(paramJSON)

	//2. 赋值: request body
	result2.Body = paramStr
	//-----------------------------------------------------------

	return result2
}

func GetKeys(m map[string]string) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}
