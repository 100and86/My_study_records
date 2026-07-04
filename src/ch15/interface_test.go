package interface_test

//interface{}study
import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"testing"
)

func TestInterface(t *testing.T) {
	p := "I am your friends!"
	var c wordCounter
	fmt.Fprint(&c, p)
	fmt.Print(c)
}

type wordCounter int

func (c *wordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p)) //建立扫描器对象
	scanner.Split(bufio.ScanWords)                  //切换为按单词读取
	for scanner.Scan() {
		*c++
		fmt.Println(scanner.Text())
	}

	return len(p), scanner.Err()
}

func TestNewWriter(t *testing.T) {
	var buf bytes.Buffer
	w, count := CounteWriter(&buf)
	fmt.Fprint(w, "hello")
	fmt.Fprint(w, ",Tom")
	fmt.Println(buf.String())
	fmt.Println(*count)
}

type counteWriter struct {
	w io.Writer
	n *int64
}

func CounteWriter(w io.Writer) (io.Writer, *int64) {
	var count int64
	cw := &counteWriter{
		w: w,
		n: &count,
	}
	return cw, &count
}

func (c *counteWriter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	*c.n += int64(n)
	return n, err
}
func TestNewReader(t *testing.T) {
	input := `
<!doctype html>
<html>
	<head>
		<title>Hello</title>
	</head>
	<body>
		<h1>Go</h1>
		<a href="https://example.com">example</a>
	</body>
</html>`
	doc, err := html.Parse(NewReader(input))
	if err != nil {
		panic(err)
	}
	forEachNode(doc, startElement, nil)
}
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Println(n.Data)
	}
}

type StringReader struct {
	s string
	i int
}

func (str *StringReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if str.i >= len(str.s) {
		return 0, io.EOF
	}
	n := copy(p, str.s[str.i:]) //以短为主,实际在缓冲区填的字节数
	str.i += n                  //读到的位置
	return n, nil
}

//使用io.Reader做参数的上层函数的内部要做的事
// buf := make([]byte, 4096)
// for {
//     n, err := r.Read(buf)
//     if n > 0 {//一定优先处理，不然会丢掉
//         // 处理 buf[:n] 里的 HTML 字节
//     }
//     if err == io.EOF {
//         break
//     }
//     if err != nil {
//         return err
//     }
// }

func NewReader(str string) io.Reader {
	return &StringReader{s: str}
}
