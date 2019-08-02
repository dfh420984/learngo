package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

//Fetch ...
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("StatusOK err %d:\n", http.StatusOK)
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	e := determinEncode(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determinEncode(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("bufio.NewReader error:%v", err)
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
