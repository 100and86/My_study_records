package main

//interface{}study
import (
	"fmt"
	"bytes"
	"bufio"
)

func main() {
	p:="I am your friends!"
	var c wordCounter
	fmt.Fprint(&c,p)
	fmt.Print(c)
}
type wordCounter int 
func (c *wordCounter) Write(p []byte) (int,error){
	scanner:=bufio.NewScanner(bytes.NewReader(p))//建立扫描器对象
	scanner.Split(bufio.ScanWords)//切换为按单词读取
	for scanner.Scan(){
		*c++
		fmt.Println(scanner.Text())
	}
	
	return len(p),scanner.Err()
}

