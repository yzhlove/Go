package parser

import (
	"journey/chat31_love/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	var (
		//url  = `http://www.zhenai.com/zhenhun`
		body []byte
		err  error
	)
	const cityListNumber = 470
	//if body, err = ioutil.ReadFile("citylist_test_data.html"); err != nil {
	//	t.Error(err.Error())
	//}
	//
	////t.Logf("%s", body)

	if body, err = fetcher.Fetch("http://www.zhenai.com/zhenghun"); err != nil {
		t.Error(err)
	}

	parseRes := ParseCityList(body, "http://www.baidu.com")

	exceptRequests := []string{"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng"}

	if len(parseRes.Requests) != cityListNumber {
		t.Errorf("request %d ", len(parseRes.Requests))
	}

	for i, url := range exceptRequests {
		if parseRes.Requests[i].URL != url {
			t.Errorf("request err:%s", parseRes.Requests[i].URL)
		}
	}

	if len(parseRes.Requests) != cityListNumber {
		t.Errorf("item %d", len(parseRes.Items))
	}

}
