package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) //处理多个客户端的请求

	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() { //只要不结束就持续等待
		go echo(c, input.Text(), 1*time.Second) //处理单个客户端的多次输入，但是这就可能造成一件事，就是混杂输入
	}

}

// 增加超时自动结束，但是存在逻辑竞争，且两个goroutine同时操控一个timer，竞争不容易发现
func handleConnTimer1(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	readDone := make(chan struct{})
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()
	go func() {
		defer close(readDone)
		for input.Scan() { //只要不结束就持续等待
			timer.Reset(10 * time.Second)
			go echo(c, input.Text(), 1*time.Second)
		}
		readDone <- struct{}{}
	}()
	select {
	case <-timer.C:
		return
	case <-readDone:
		return
	}

}

// 优化为一个goroutine操作timer，逻辑更清晰
func handleConnTimer2(c net.Conn) {
	defer c.Close()

	input := bufio.NewScanner(c)

	text := make(chan string)
	readDone := make(chan struct{}, 1)
	defer close(readDone)

	go func() {
		defer close(text)
		for input.Scan() {
			select {
			case text <- input.Text():
			case <-readDone:
				return
			}

		}
		if err := input.Err(); err != nil {
			log.Printf("读取客户端结束：%v", err)
		}

	}()
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			fmt.Print("10s未输入，自动退出")
			return
		case line, ok := <-text:
			if !ok {
				return
			}
			timer.Reset(10 * time.Second)
			go echo(c, line, 1*time.Second)
		}
	}
}
