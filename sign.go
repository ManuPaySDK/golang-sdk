package manupay_client

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func GenSign(params map[string]interface{}, privateSecret string) string {
	md5ctx := md5.New()
	keys := make([]string, 0, len(params))

	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, k := range keys {
		vs := params[k]
		if vs == "" {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}

		buf.WriteString(k)
		buf.WriteByte('=')
		//类型检查
		switch vv := vs.(type) {
		case string:
			buf.WriteString(vv)
		case int:
			buf.WriteString(strconv.FormatInt(int64(vv), 10))
		case int64:
			buf.WriteString(strconv.FormatInt(int64(vv), 10))
		default:
			panic(fmt.Sprintf("params type not supported, k=%s, %+v", k, vv))
		}
	}
	buf.WriteString(fmt.Sprintf("&key=%s", privateSecret))

	fmt.Printf("Sign-rawSignStr = %s\n", buf.String())
	md5ctx.Write([]byte(buf.String()))
	return strings.ToUpper(hex.EncodeToString(md5ctx.Sum(nil)))
}

// 验证签名
func VerifySign(params map[string]interface{}, privateSecret string, sign string) bool {
	//自己算一遍
	selfSign := GenSign(params, privateSecret)
	return selfSign == sign
}
