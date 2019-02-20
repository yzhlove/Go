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

func Fetch(url string) ([]byte, error) {
	var (
		response *http.Response
		err      error
	)

	if response, err = http.Get(url); err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", response.StatusCode)
	}
	bodyReader := bufio.NewReader(response.Body)
	e := getEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
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
