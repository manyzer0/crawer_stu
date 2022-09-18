package engine

import (
	"fmt"
)

type ConEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	SetWorkerChan(chan Request)
}

func (e *ConEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.SetWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(in, out)

	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
		Cityurls[r.Url] = 1
	}
	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("got item %v\n", item)
		}

		for _, req := range result.Request {
			if _, ok := Cityurls[req.Url]; ok {
				continue
			}
			Cityurls[req.Url] = 1
			e.Scheduler.Submit(req)
		}
	}
}

func (e *ConEngine) createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			res, err := Worker(request)
			if err != nil {
				continue
			}
			out <- res
		}
	}()
}
