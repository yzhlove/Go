package util

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"
)

//Sign  签名
func Sign(m map[string]string, secret string) string {
	params := make([]string, 0, len(m))
	for _, value := range []string{"_url", "h", "file"} {
		delete(m, value)
	}
	for k, v := range m {
		params = append(params, k+"="+v)
	}
	sort.Strings(params)
	str := strings.Join(params, "&")
	return encryptMD5(encryptMD5(str) + secret)
}

//MD5 加密
func encryptMD5(urlStr string) string {
	ctx := md5.New()
	ctx.Write([]byte(urlStr))
	s := ctx.Sum(nil)
	return hex.EncodeToString(s)
}
