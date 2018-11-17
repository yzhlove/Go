package matchers

import (
	"regexp"
	"log"
	"fmt"
	"net/http"
	"errors"
	"encoding/xml"
	"sample/search"
)

//rss 匹配器

type (
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)


type rssMatcher struct{}

func init() {
	var matcher rssMatcher
	search.Register("rss",matcher)
}

func (m rssMatcher)retrieve(feed *search.Feed) (*rssDocument,error) {
	if feed.URI == "" {
		return nil, errors.New("No Rss feed URI!")
	}

	resp , err := http.Get(feed.URI)
	if err != nil {
		return nil,err
	}
	//无论如何，关闭连接
	defer resp.Body.Close()

	if resp.StatusCode!= 200 {
		return nil , fmt.Errorf("HTTP Response Error %d\n",resp.StatusCode)
	}

	//解析
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document,err
} 


//Search 在文档中查找特定的搜索项
func (m rssMatcher)Search(feed *search.Feed,searchTerm string) ([]*search.Result,error) {

	var results []*search.Result
	
	log.Printf("Search Feed Type[%s] Site[%s] For Uri[%s]\n",feed.Type,feed.Name,feed.URI)

	document , err := m.retrieve(feed)
	if err != nil {
		return nil,err
	}

	for _, channelItem := range document.Channel.Item {
		matched , err := regexp.MatchString(searchTerm,channelItem.Title)
		//检查标题部分是否含有搜索项
		if err != nil {
			return nil,err
		}
		//如果找到匹配的项，将其作为结果
		if matched {
			results = append(results, &search.Result{
				Field:"Title",
				Content:channelItem.Title,
			}) 
		}

		//检查描述部分是否含有搜索项
		matched , err = regexp.MatchString(searchTerm,channelItem.Description)
		if  err != nil {
			return nil , err
		}

		//如果找到匹配项，将其作为结果保存
		if matched {
			results = append(results,&search.Result{
				Field:"Description",
				Content:channelItem.Description,
			})
		}

	}
	return results,nil
}



