package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Clock struct {
	Name    string
	Address string
}
type Update struct {
	Name  string
	Value string
}

func main() {
	if len(os.Args[1:]) == 0 {
		log.Fatal(
			"请传入时钟服务器，例如：" +
				" NewYork=127.0.0.1:8010",
		)
	}
	updates := make(chan Update)
	clocks := make([]Clock, 3, 5)
	for _, argument := range os.Args[1:] {
		clock, err := parseClock(argument)
		if err != nil {
			log.Print(err)
			continue
		}
		go readClock(clock, updates)
		clocks = append(clocks, clock)
		fmt.Printf("%-15s", clock.Name)
	}
	fmt.Println()
	latest := make(map[string]string)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case update := <-updates:
			latest[update.Name] = update.Value

		case <-ticker.C:
			printTimes(clocks, latest)
		}
	}

}

func readClock(clock Clock, updates chan<- Update) {
	conn, err := net.Dial("tcp", clock.Address)
	if err != nil {
		updates <- Update{
			Name:  clock.Name,
			Value: "连接失败",
		}
	}
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		updates <- Update{
			Name:  clock.Name,
			Value: scanner.Text(),
		}
	}
	if err := scanner.Err(); err != nil {
		updates <- Update{
			Name:  clock.Name,
			Value: "读取失败",
		}
	}

}

func parseClock(argument string) (Clock, error) {
	name, address, found := strings.Cut(argument, "=")
	if !found || name == "" || address == "" {
		return Clock{}, fmt.Errorf(
			"参数格式错误：%q，应为 名称=地址",
			argument,
		)
	}

	return Clock{
		Name:    name,
		Address: address,
	}, nil
}
func printTimes(clocks []Clock, latest map[string]string) {
	for _, clock := range clocks {
		value := latest[clock.Name]
		if value == "" {
			value = "--:--:--"
		}

		fmt.Printf("%-15s", value)
	}

	fmt.Println()
}
