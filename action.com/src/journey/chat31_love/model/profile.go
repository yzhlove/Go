package model

import "encoding/json"

type Profile struct {
	NickName   string
	Gender     string // 性别
	Age        string
	Height     string
	Weight     string
	Income     string //收入
	Marriage   string //婚姻状况
	Education  string //教育程度
	Occupation string //职业
	HoKou      string
	Xinzuo     string
	House      string
	Car        string
}

func ToJson(data interface{}) (Profile, error) {
	var (
		bytes   []byte
		err     error
		profile Profile
	)
	if bytes, err = json.Marshal(data); err != nil {
		return profile, err
	}
	err = json.Unmarshal(bytes, &profile)
	return profile, err
}
