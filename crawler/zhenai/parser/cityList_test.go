package parser

import (
	"learngo/crawler/fetcher"
	"testing"
)

func TestCityList(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun/")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s", contents)
	result := CityList(contents)
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("expect size %d, but actual size %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("expect size %d, but actual size %d", resultSize, len(result.Items))
	}
}
