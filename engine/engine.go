package engine

import (
	"crawlergo/fetch"
	"fmt"
	"log"
)

var Cityurls map[string]int

func Run(seeds ...Request) {
	var requests []Request

	for _, seed := range seeds {
		requests = append(requests, seed)
	}
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]
		if _, ok := Cityurls[request.Url]; ok {
			continue
		}
		Cityurls[request.Url] = 1
		// b, err := fetch.Fetch(request.Url)
		// fmt.Println("url is :" + request.Url)
		// if err != nil {
		// 	log.Println(err.Error())
		// 	continue
		// }
		// r := request.ParseFunc(b)
		r, err := Worker(request)
		if err != nil {
			continue
		}
		requests = append(requests, r.Request...)
		fmt.Println(Cityurls)
		fmt.Println(r.Items...)
		for _, u := range r.Request {
			fmt.Println(u.Url)

		}

	}
}

func Worker(request Request) (ParseResult, error) {
	b, err := fetch.Fetch(request.Url)
	fmt.Println("url is :" + request.Url)
	if err != nil {
		log.Println(err.Error())
		return ParseResult{}, err
	}
	r := request.ParseFunc(b)
	return r, nil
}
