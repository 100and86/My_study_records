package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

type Article struct {
	Title  string
	Author string
}

type Page struct {
	Title    string
	Articles []Article
}

const tpl = `
<h1>{{.Title}}</h1>

<ul>
{{range .Articles}}
  <li>{{.Title}} - 作者：{{.Author}}</li>
{{end}}
</ul>
`

func main() {
	page := Page{
		Title: "文章列表",
		Articles: []Article{
			{Title: "Go入门", Author: "张三"},
			{Title: "HTML学习", Author: "李四"},
			{Title: "协程学习", Author: "王老五"},
		},
	}

	t := template.Must(template.New("page").Parse(tpl))

	outputPath := filepath.Join(`D:\py file\Go_learn_file\Study\src\ch13\`, "output.html")
	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("创建失败")
		return
	}
	defer file.Close()

	err = t.Execute(file, page)
	if err != nil {
		fmt.Println("注入失败")
		return
	}
}
