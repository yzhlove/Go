package chat10

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

//模拟Http调用

const checkMark = "\u2713"
const ballotX = "\u2717"

var feed = `
<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet title="XSL_formatting" type="text/xsl" href="/shared/bsp/xsl/rss/nolsol.xsl"?>
<rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" version="2.0" xmlns:media="http://search.yahoo.com/mrss/">
    <channel>
        <title><![CDATA[BBC News - Business]]></title>
        <description><![CDATA[BBC News - Business]]></description>
        <link>https://www.bbc.co.uk/news/</link>
        <image>
            <url>https://news.bbcimg.co.uk/nol/shared/img/bbc_news_120x60.gif</url>
            <title>BBC News - Business</title>
            <link>https://www.bbc.co.uk/news/</link>
        </image>
        <generator>RSS for Node</generator>
        <lastBuildDate>Thu, 27 Dec 2018 03:31:45 GMT</lastBuildDate>
        <copyright><![CDATA[Copyright: (C) British Broadcasting Corporation, see http://news.bbc.co.uk/2/hi/help/rss/4498287.stm for terms and conditions of reuse.]]></copyright>
        <language><![CDATA[en-gb]]></language>
        <ttl>15</ttl>
        <item>
            <title><![CDATA[Boxing Day sales: Footfall down for third year, analysts say]]></title>
            <description><![CDATA[Analysts say footfall has dropped for the third year, although London appears to have bucked the trend.]]></description>
            <link>https://www.bbc.co.uk/news/business-46684639</link>
            <guid isPermaLink="true">https://www.bbc.co.uk/news/business-46684639</guid>
            <pubDate>Wed, 26 Dec 2018 18:55:22 GMT</pubDate>
            <media:thumbnail width="976" height="549" url="http://c.files.bbci.co.uk/E1E1/production/_104952875_29b5a086-400f-4cae-9bea-cc3df90f0863.jpg"/>
        </item>
        <item>
            <title><![CDATA[US stock markets rally after pre-Christmas slump]]></title>
            <description><![CDATA[The Dow Jones is up by nearly 5% and the technology-focused Nasdaq rises by nearly 6%.]]></description>
            <link>https://www.bbc.co.uk/news/business-46690452</link>
            <guid isPermaLink="true">https://www.bbc.co.uk/news/business-46690452</guid>
            <pubDate>Thu, 27 Dec 2018 00:37:33 GMT</pubDate>
            <media:thumbnail width="2048" height="1152" url="http://c.files.bbci.co.uk/3E03/production/_104957851_051340958.jpg"/>
        </item>
    </channel>
</rss>`

//模拟服务器
func mockServer() *httptest.Server {

	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		if _, err := fmt.Fprintln(w, feed); err != nil {
			log.Println("WriteErr:", err)
		}
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

//模拟服务器数据获取
func TestDownload(t *testing.T) {

	statusCode := http.StatusOK

	testServer := mockServer()
	defer testServer.Close()

	t.Log("Start Test...")
	{
		t.Logf("\tTest URL and StatusCode : %v %d \n", testServer.URL, statusCode)
		{
			resp, err := http.Get(testServer.URL)
			if err != nil {
				t.Fatalf("\t\tGet Err:%v %v\n", ballotX, err)
			}
			t.Log("\t\tGet Successful:", checkMark)

			defer func() {
				if err := resp.Body.Close(); err != nil {
					t.Error("\t\t\tClose Err:", ballotX, err)
				}
			}()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t\tReceive Err: %d status:%v %v \n", statusCode, ballotX, resp.StatusCode)
			}
			t.Logf("\t\t\tReceive Successful: %d %v \n", statusCode, checkMark)
		}
	}

}
