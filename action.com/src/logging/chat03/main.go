package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Json 编码/解码

type (
	//文档结构
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapeUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	//文档目录
	gResponse struct {
		ErrorCode    int    `json:"error_code"`
		ErrorReason  string `json:"-"`
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func main() {

	uri := "http://jump.com/test/go/index"

	resp, err := http.Get(uri)
	if err != nil {
		log.Println("-ERROR:", err)
		return
	}

	defer resp.Body.Close()

	//Json解析
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("+ERROR:", err)
		return
	}
	fmt.Println(gr)
}
