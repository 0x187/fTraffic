package main

import (
	OpenFile "fTraffic/internal/FileOpen"
	"fTraffic/internal/send"
	"flag"
	"net"
	"sync"
	"time"
)

func main() {
	var goroutineNumber int
	var ipListFile string
	var timeSleep time.Duration
	flag.IntVar(&goroutineNumber, "T", 1, "The number of goroutines")
	flag.StringVar(&ipListFile, "ip", "ip.txt", "The ip list file")
	flag.DurationVar(&timeSleep, "time", time.Millisecond*53, "The send rate") // 53 Millisecond 100GB per 24h, 100ms -50Gb per 24h

	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(goroutineNumber)
	ipList := OpenFile.OpenFile(ipListFile)
	conn, _ := net.ListenUDP("udp", &net.UDPAddr{Port: 8191})

	//start := time.Now() //for debug

	for i := 0; i < goroutineNumber; i++ {
		go send.Send(&wg, conn, ipList, timeSleep)
		//go send.Send(&wg, conn, ipList, timeSleep, start) //for debug
	}
	wg.Wait()
}
