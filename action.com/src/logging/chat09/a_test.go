package chat09

import (
	"net/http"
	"testing"
)

//表组测试

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {

	var urls = []struct {
		url        string
		statusCode int
	}{
		{url: "http://www.baidu.com", statusCode: http.StatusOK},
		{url: "https://cn.bing.com/", statusCode: http.StatusOK},
		{url: "http://www.google.com", statusCode: http.StatusOK},
		{url: "http://www.huaban.com", statusCode: http.StatusOK},
		{url: "http://www.xianhua.com", statusCode: http.StatusOK},
		{url: "http://www.youku.com", statusCode: http.StatusOK},
		{url: "http://www.telnet.com", statusCode: http.StatusOK},
		{url: "http://www.tudou.com", statusCode: http.StatusOK},
		{url: "http://www.iqiyi.com", statusCode: http.StatusOK},
		{url: "http://www.007ts.com", statusCode: http.StatusOK},
	}

	t.Log("Start Table Test ....")
	{
		for _, u := range urls {
			t.Logf("\tCurrent URL : %v %d \n", u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Error("\t\tURL Error:", ballotX, err)
					continue
				}
				t.Log("\t\tGet URL:", checkMark)
				defer func() {
					if err := resp.Body.Close(); err != nil {
						t.Error("\t\t\t resp Close Err:", err)
					}
				}()
				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tURL Successful:%d %v \n", u.statusCode, checkMark)
				} else {
					t.Errorf("\t\tURL Failure:%d %v %v \n", u.statusCode, ballotX, resp.StatusCode)
				}
			}
		}
	}

}
