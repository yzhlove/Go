package search

import (
	"encoding/json"
	"os"
)

const dataFile string = "/Users/love/WorkSpace/Go/action.com/src/sample/data/data.json"

//Feed  Rss源信息
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

//RetrieveFeeds  读取并反序列化数据源文件
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	//无论如何关闭文件流
	defer file.Close()

	//反序列化data文件
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
