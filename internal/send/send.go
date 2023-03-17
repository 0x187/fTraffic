package send

import (
	"fTraffic/internal/randomINT"
	"fmt"
	"github.com/dustin/go-humanize"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

var e int64 = 0
var totalSendPaketSize uint64 = 0

//func Send(wg *sync.WaitGroup, conn *net.UDPConn, ip []string, timeSleep time.Duration, start time.Time) { //for debug

func Send(wg *sync.WaitGroup, conn *net.UDPConn, ip []string, timeSleep time.Duration) {
	rand.Seed(time.Now().Unix())
	for {
		//sendPacket(conn, &net.UDPAddr{IP: net.ParseIP(ip[rand.Intn(len(ip))]), Port: 21}, 62000, start) //for debug
		sendPacket(conn, &net.UDPAddr{IP: net.ParseIP(ip[rand.Intn(len(ip))]), Port: 21}, 62000)
		time.Sleep(timeSleep)
	}
	wg.Done()
}

// func sendPacket(conn *net.UDPConn, addr *net.UDPAddr, UserInputSize int, start time.Time) { //for debug
func sendPacket(conn *net.UDPConn, addr *net.UDPAddr, UserInputSize int) {
	paketSize := randomINT.RandomINT(UserInputSize, UserInputSize+3506)

	paketSizeSUM(paketSize)
	buf := make([]byte, paketSize)

	_, err := rand.Read(buf)
	if err != nil {
		log.Fatalf("err while generating random string: %s", err)
	}

	n, err := conn.WriteTo(buf, addr)
	if err != nil {
		e++
		time.Sleep(time.Second * 5)
		_, err := conn.WriteTo(buf, &net.UDPAddr{Port: 53, IP: net.IP{1, 1, 1, 1}})
		if err != nil {
			fmt.Println("ERROR")
			e++
		}
	}

	//elapsed := time.Since(start).Abs() //for debug
	//fmt.Printf("%-15v Send: %d bytes -> %-20s | total send: %-7s | drop: %v \n", elapsed, n, addr, humanize.Bytes(totalSendPaketSize), e) //for debug
	fmt.Printf("Send: %d bytes -> %-20s | total send: %-7s | drop: %v \n", n, addr, humanize.Bytes(totalSendPaketSize), e)
}

func paketSizeSUM(new int) {
	tmp := uint64(new)
	totalSendPaketSize = tmp + totalSendPaketSize
}
