package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		URL:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParserCityList,
	})
}
