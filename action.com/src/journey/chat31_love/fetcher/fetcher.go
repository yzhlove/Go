package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/text/transform"

	"golang.org/x/text/encoding/unicode"

	"golang.org/x/net/html/charset"

	"golang.org/x/text/encoding"
)

//浏览器标识
const (
	browserTag = `User-Agent`
	userAgent  = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36`
)

func Fetch(url string) ([]byte, error) {
	var (
		request  *http.Request
		response *http.Response
		err      error
		bytes    []byte
	)

	if request, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return nil, err
	}
	request.Header.Add(browserTag, userAgent)
	if response, err = http.DefaultClient.Do(request); err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", response.StatusCode)
	}
	bodyReader := bufio.NewReader(response.Body)
	e := getEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	bytes, err = ioutil.ReadAll(utf8Reader)
	return bytes, err
}

func getEncoding(read *bufio.Reader) encoding.Encoding {
	var (
		bytes []byte
		err   error
	)

	if bytes, err = read.Peek(1024); err != nil {
		log.Printf("Fetcher err: %v", err)
		return unicode.UTF8
	}
	encode, _, _ := charset.DetermineEncoding(bytes, "")
	return encode
}
