package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	var port string
	var zone string
	flag.StringVar(&port, "port", "8000", "服务器监听")
	flag.StringVar(&zone, "zone", "Local", "时区，例如 Asia/Tokyo")
	flag.Parse()

	address := "127.0.0.1:" + port
	location, err := time.LoadLocation(zone)
	if err != nil {
		log.Fatalf("无效时区 %q：%v", zone, err)
	}
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Printf("监听%s端口中,时区为%s", port, zone)

	for {
		conn, err := listener.Accept() //阻塞时调用
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, location)

	}

}

func handleConn(conn net.Conn,location *time.Location) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().In(location).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
