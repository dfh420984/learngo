package engine

import (
	"learngo/crawler/fetcher"
	"log"
)

//Run ...
func Run(seeds ...Request) {
	requests := make([]Request, 0)
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("fetching url %s", r.URL)
		body, err := fetcher.Fetch(r.URL)
		if err != nil {
			log.Printf("fetcher.Fetch Error url : %s, err: %v", r.URL, err)
			continue
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
