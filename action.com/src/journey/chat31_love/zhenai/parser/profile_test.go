package parser

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"
)

func TestProfile(t *testing.T) {

	const (
		url       = `http://album.zhenai.com/u/1426975040`
		userReg   = `<div class="m-btn [a-z]+" data-v-[a-zA-Z0-9]+>([^<]+)</div>`
		infoReg   = `<div class="m-btn pink" data-v-[a-zA-Z0-9]+>([^<]+)</div>`
		headTag   = `User-Agent`
		userAgent = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36`
		tempReg   = `<div class="m-btn [a-z]*" data-v-[a-zA-Z0-9]+>([^<]+)</div>`
	)
	var (
		response *http.Response
		request  *http.Request
		err      error
		bytes    []byte
	)
	if request, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		panic(err)
	}

	request.Header.Add(headTag, userAgent)
	if response, err = http.DefaultClient.Do(request); err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if bytes, err = ioutil.ReadAll(response.Body); err != nil {
		panic(err)
	}

	//t.Logf("%s", bytes)

	reg := regexp.MustCompile(tempReg)
	matchers := reg.FindAllSubmatch(bytes, -1)

	if len(matchers) == 0 {
		t.Error("match err")
	}

	for i := 0; i < len(matchers); i++ {
		t.Logf("%d %s -> %s\n", i, matchers[i][1], matchers[i][0])
	}

}

func TestParseURL(t *testing.T) {

	const url = `http://album.zhenai.com/u/([0-9]+)`
	const data = `http://album.zhenai.com/u/1426975040`
	reg := regexp.MustCompile(url)
	result := reg.FindStringSubmatch(data)
	t.Logf("result : %v ", result)
}
