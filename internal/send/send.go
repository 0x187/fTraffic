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

var totalSendPaketSize uint64 = 0

func Send(wg *sync.WaitGroup, conn *net.UDPConn, ip []string) {
	rand.Seed(time.Now().Unix())
	for {
		sendPacket(conn, &net.UDPAddr{IP: net.ParseIP(ip[rand.Intn(len(ip))]), Port: 443}, 63000)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func sendPacket(conn *net.UDPConn, addr *net.UDPAddr, UserInputSize int) {
	paketSize := randomINT.RandomINT(UserInputSize-1000, UserInputSize+2506)

	paketSizeSUM(paketSize)
	buf := make([]byte, paketSize)

	_, err := rand.Read(buf)
	if err != nil {
		log.Fatalf("err while generating random string: %s", err)
	}
	n, err := conn.WriteTo(buf, addr)
	if err != nil {
		log.Fatal("Write:", err)
	}
	fmt.Println("  Sent   ", n, "   bytes  ->   ", addr, "    ", humanize.Bytes(totalSendPaketSize))
}

func paketSizeSUM(new int) {
	tmp := uint64(new)
	totalSendPaketSize = tmp + totalSendPaketSize
}
