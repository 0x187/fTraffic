package main

import (
	OpenFile "fTraffic/internal/FileOpen"
	"fTraffic/internal/send"
	"flag"
	"runtime"
	"sync"
)

func main() {
	var UserSizeINPUT int
	flag.IntVar(&UserSizeINPUT, "t", 1, "x numbers of goroutines running at time")
	flag.Parse()

	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	wg.Add(UserSizeINPUT)
	ipList := OpenFile.OpenFile("./assets/ip.txt")

	for i := 0; i < UserSizeINPUT; i++ {
		go send.Send(&wg, ipList)
	}
	wg.Wait()
}
