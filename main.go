package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/dustin/go-humanize"
)

var bytesSiz uint64 = 0
var aGb int = 1003741824
var daySeconds int = 8640000
var UserSizeINPUT int
var ipList = "ip.txt"

func randomINT(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func randomIP() string {
	//randSource := rand.NewSource(time.Now().UnixNano())
	//randGenerator := rand.New(randSource)
	//firstLoc := randGenerator.Intn(10)
	//candidate1 := ""
	//dat, err := ioutil.ReadFile(ipList)
	//if err == nil {
	//	ascii := string(dat)
	//	splt := strings.Split(ascii, "\n")
	//	candidate1 = splt[firstLoc]
	//
	//}
	//return candidate1
	file, err := os.Open(ipList)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	randsource := rand.NewSource(time.Now().UnixNano())
	randgenerator := rand.New(randsource)

	lineNum := 1
	var pick string
	for scanner.Scan() {
		line := scanner.Text()
		roll := randgenerator.Intn(lineNum)
		if roll == 0 {
			pick = line
		}
		lineNum += 1
	}
	return pick
}

func sendPacket(conn *net.UDPConn, addr *net.UDPAddr, userByteSize int) {
	randomtmp := randomINT((userByteSize - 400), (userByteSize + 400))

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
	bytesSiz = tmp + bytesSiz

}

func main() {

	flag.IntVar(&UserSizeINPUT, "size", 2, "number of lines to read from the file")
	flag.Parse()

	userByteSize := (UserSizeINPUT * aGb) / daySeconds
	fmt.Println(userByteSize)

	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 1234})
	if err != nil {
		log.Fatal("Listen:", err)
	}
	for true {
		sendPacket(conn, &net.UDPAddr{IP: net.ParseIP(randomIP()), Port: 443}, userByteSize)
		time.Sleep(time.Millisecond * 100)
	}

}
