package persist

import (
	"context"
	"encoding/json"
	"journey/chat31_love/engine"
	"journey/chat31_love/model"
	"testing"

	elastic "gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {

	var (
		err     error
		resp    *elastic.GetResult
		client  *elastic.Client
		actual  engine.Item
		profile model.Profile
		index   = "dating_test"
	)

	excepted := engine.Item{
		URL:  "http://album.zhenai.com/u/1769167712",
		Id:   "1769167712",
		Type: "zhenai",
		Detail: model.Profile{
			Age:        "34岁",
			Height:     "162",
			Weight:     "57",
			Income:     "3000-5000元",
			Gender:     "女",
			NickName:   "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	if client, err = elastic.NewClient(elastic.SetSniff(false)); err != nil {
		panic(err)
	}
	if err = save(client, index, excepted); err != nil {
		panic(err)
	}
	if resp, err = client.Get().Index("dating_profile").
		Type(excepted.Type).Id(excepted.Id).Do(context.Background()); err != nil {
		panic(err)
	}
	t.Logf("%s \n", *resp.Source)
	if err = json.Unmarshal(*resp.Source, &actual); err != nil {
		panic(err)
	}
	//t.Logf("ToJson:%+v \n", actual)
	if profile, err = model.ToJson(actual.Detail); err != nil {
		t.Errorf("complie err %v", err)
		return
	}
	actual.Detail = profile
	if actual != excepted {
		t.Errorf("object err!")
	}

}
