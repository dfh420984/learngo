package parser

import (
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
)

var comRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">(.*)</div>`)

//ParseProfile ...
func ParseProfile(contents []byte) engine.ParseResult {
	profile := extractString(contents, comRe)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) model.Profile {
	match := re.FindAllSubmatch(contents, -1)
	fmt.Printf("%v", match)
	profile := model.Profile{}
	return profile
}
