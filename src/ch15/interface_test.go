package interface_test

//interface{}study
import (
	"fmt"
	"bytes"
	"bufio"
	"testing"
	"io"

)

func TestInterface(t *testing.T) {
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

func TestNewWriter(t *testing.T){
	var buf bytes.Buffer
	w,count:=CounteWriter(&buf)
	fmt.Fprint(w,"hello")
	fmt.Fprint(w,",Tom")
	fmt.Println(buf.String())
	fmt.Println(*count)
}
type counteWriter struct {
	w io.Writer
	n *int64
}
func CounteWriter(w io.Writer)(io.Writer, *int64){
	var count int64
	cw := &counteWriter{
		w: w,
		n: &count,
	}
	return cw,&count
}

func (c *counteWriter) Write(p []byte)(int,error){
	n,err:=c.w.Write(p)
	*c.n += int64(n) 
	return n,err
}



