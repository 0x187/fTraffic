package main

import (
	OpenFile "fTraffic/internal/FileOpen"
	"fTraffic/internal/send"
	"flag"
	"net"
	"runtime"
	"sync"
)

func main() {
	var goroutineNumber int
	flag.IntVar(&goroutineNumber, "t", 1, "The number of goroutines")
	flag.Parse()

	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(goroutineNumber)
	ipList := OpenFile.OpenFile("./assets/ip.txt")
	conn, _ := net.ListenUDP("udp", &net.UDPAddr{Port: 1234})
	for i := 0; i < goroutineNumber; i++ {
		go send.Send(&wg, conn, ipList)
	}
	wg.Wait()
}
