package main

import (
	"crawlergo/engine"
	"crawlergo/parses"
	"crawlergo/scheduler"
)

func main() {

	// 例如本案例:fatal error: all goroutines are asleep - deadlock!
	// sch := make(chan string)
	// sch <- "string"
	// fmt.Println(<-sch)

	// 处理方法,使用双协程,并且接收者在发送者就绪之前,-> 运行正常
	// sch := make(chan string)
	//
	// go func() {
	// 	fmt.Println(<-sch)
	// }()
	//
	// time.Sleep(time.Second)
	// sch <- "string"
	engine.Cityurls = map[string]int{}
	e1, e2, e3 := engine.Request{}, engine.Request{}, engine.Request{}
	e1.Url = "http://www.zhenai.com/zhenghun/shenzhen/1"
	e1.ParseFunc = parses.PaserUser
	e2.Url = "http://www.zhenai.com/zhenghun/shenzhen/2"
	e2.ParseFunc = parses.PaserUser
	e3.Url = "http://www.zhenai.com/zhenghun/shenzhen/3"
	e3.ParseFunc = parses.PaserUser
	r := engine.ConEngine{WorkerCount: 2, Scheduler: &scheduler.SimpleScheduler{}}
	r.Run(e1)
}
