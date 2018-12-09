package util

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"
)

//Sign 登陆请求签名算法
func Sign(m map[string]string, secret string) string {
	var params []string
	for _, v := range []string{"_url", "h", "file"} {
		delete(m, v)
	}
	for k, v := range m {
		params = append(params, k+"="+v)
	}
	sort.Strings(params)
	str := strings.Join(params, "&")
	return MD5(MD5(str) + secret)
}

//MD5 md5加密算法
func MD5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	s := ctx.Sum(nil)
	return hex.EncodeToString(s)
}
