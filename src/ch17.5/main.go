package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	port := flag.String("port", "8000", "服务器端口")

	flag.Parse()

	fmt.Println("os.Args[1:]:", os.Args[1:])
	fmt.Println("port:", *port)
	for _, v := range flag.Args() {
		fmt.Print(v,"\n")
	}
}
