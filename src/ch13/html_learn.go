package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type User struct {
	Name string
	Age  int
}

const tpl = `
	<p>姓名：{{.Name}}</p>
	<p>年龄：{{.Age}}</p>
	`

func main() {
	user := User{
		Name: "张三",
		Age:  18,
	}

	t := template.Must(template.New("user").Parse(tpl))

	file, err := os.Create("output.html")
	if err != nil {
		fmt.Println("创建失败")
		return
	}
	defer file.Close()

	err = t.Execute(file, user)
	if err != nil {
		log.Fatal(err)
	}
}
