package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/dustin/go-humanize"
)

var bytesSiz uint64 = 0
var aGb int = 1003741824
var daySecnd int = 15280000
var packetRange int = 0
var UserSizeINPUT int

func hardWork(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Start: %v\n", time.Now())

	// Memory
	a := []string{}
	for i := 0; i < 500000; i++ {
		a = append(a, "aaaa")
	}

	// Blocking
	time.Sleep(2 * time.Second)
	fmt.Printf("End: %v\n", time.Now())
}

func randomINT(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func randomIP() string {
	firstLoc := randomINT(1, 400)
	candidate1 := ""
	dat, err := ioutil.ReadFile("ip.txt")
	if err == nil {
		ascii := string(dat)
		splt := strings.Split(ascii, "\n")
		candidate1 = splt[firstLoc]
	}
	return candidate1
}

func sendPacket(conn *net.UDPConn, addr *net.UDPAddr, userByteSize int) {
	randomtmp := randomINT((userByteSize), (userByteSize + 200))

	bytesSUM(randomtmp)
	buf := make([]byte, randomtmp)

	_, err := rand.Read(buf)
	if err != nil {
		log.Fatalf("error while generating random string: %s", err)
	}
	n, err := conn.WriteTo([]byte(buf), addr)
	if err != nil {
		log.Fatal("Write:", err)
	}

	fmt.Println("Sent   ", n, "   bytes  ->   ", addr, "    ", humanize.Bytes(bytesSiz))
}

func bytesSUM(new int) {
	tmp := uint64(new)
	bytesSiz = (tmp + bytesSiz)

}

func main() {

	flag.IntVar(&UserSizeINPUT, "size", 2, "number of lines to read from the file")
	flag.Parse()

	userByteSize := ((UserSizeINPUT * aGb) / daySecnd)
	fmt.Println(userByteSize)

	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 1234})
	if err != nil {
		log.Fatal("Listen:", err)
	}
	for int(bytesSiz) < (UserSizeINPUT * aGb) {
		sendPacket(conn, &net.UDPAddr{IP: net.ParseIP(randomIP()), Port: 443}, userByteSize)
		time.Sleep(time.Millisecond * 100)
	}

}
