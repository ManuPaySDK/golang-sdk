package manupay_client

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {

	//代签名query参数
	signMap2 := map[string]interface{}{
		"platId":     "1000",
		"mchOrderNo": "P0123456789101",
		"amount":     "10000",
		"clientIp":   "192.168.0.111",
		"returnUrl":  "https://www.google.com",
		"notifyUrl":  "https://www.google.com",
		"reqTime":    "20190723141000",
		"version":    "1.0",
	}
	_, result2 := GenSign(signMap2, "EWEFD123RGSRETYDFNGFGFGSHDFGH")
	fmt.Printf("%s\n", result2)
}

func TestSign2(t *testing.T) {

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

	//代签名query参数
	signMap2 := map[string]interface{}{
		"amount":     1000,
		"appId":      "6416b86a79fc2b78402dbb3f",
		"body":       "toysDesc",
		"clientIp":   "119.57.72.252",
		"createdAt":  1694180893647,
		"currency":   "inr",
		"extParam":   string(res),
		"ifCode":     "manupay",
		"mchNo":      "M1679210601",
		"mchOrderNo": "exop20230908001",
		"payOrderId": "P1700144009406652417",
		"reqTime":    1694181364379,
		"state":      2,
		"subject":    "toys",
		"wayCode":    "SAIL_CASHIER",
		//"key":XcX2XbEUGkklmv8OREpQBoE0xBA0VMNXATigyRHFCE3NC6puxX9l8RiPudJLJ6LB4lSbnikFH2mNFiiDZypmNOlDmCnrmJeFnX44giVf8vpOqiLLr4GRIb9uJG0V0KxI,
	}
	_, result2 := GenSign(signMap2, "XcX2XbEUGkklmv8OREpQBoE0xBA0VMNXATigyRHFCE3NC6puxX9l8RiPudJLJ6LB4lSbnikFH2mNFiiDZypmNOlDmCnrmJeFnX44giVf8vpOqiLLr4GRIb9uJG0V0KxI")
	fmt.Printf("%s\n", result2)
}
