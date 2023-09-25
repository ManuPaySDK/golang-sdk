package manupay_client

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"testing"
	"time"
)

// 下单
func TestPlaceOrder(t *testing.T) {
	orderId := "exop20230908006"

	params := map[string]string{
		"firstname": "cy",
		"lastname":  "harper",
		"city":      "guangzhou",
		"phone":     "4401000001",
		"email":     "ck789@gmail.com",
		"country":   "IN",
		"address":   "baiyun district",
		"state":     "mh",
		"postcode":  "232001",
	}
	res, _ := json.Marshal(params)

	//-----------------------------------

	client := NewManuPayClient("http://127.0.0.1:9002", MchNo, PrivateSecret)
	isSucceed, response := client.PlaceUnifiedOrder(UnifiedOrderRequest{
		MchOrderNo: orderId,
		WayCode:    "SAIL_CASHIER",
		Amount:     1000,
		Currency:   "inr",
		Subject:    "toys",
		Body:       "toysDesc",
		NotifyUrl:  "https://www.jpdb001.com/notifyUrl",
		ReturnUrl:  "https://www.jpdb001.com/returnUrl",
		ExtParam:   string(res),
		ReqTime:    time.Now().UnixMilli(),
		//可选
		ExpiredTime: 3600,
	})
	fmt.Printf("result=%v\nresp=%+v\n", isSucceed, response)

	if response.Code == 0 {
		//验证返回的签名
		rawParams := structs.Map(response.Data)
		_, signVal := GenSign(rawParams, client.PrivateSecret)
		if signVal != response.Sign {
			fmt.Printf("-----sign---err---%s\n", signVal)
		} else {
			fmt.Printf("-----sign---succ---%s\n", signVal)
		}
	}
}
