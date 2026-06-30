package main

//子字符串寻找
import (
	"fmt"
)

func main() {
	text := "i am study go"
	pat := "go"
	fmt.Println(SundayIndexRune(text, pat))
}

func SundayIndexRune(text string, pat string) int {
	m := len(pat)
	n := len(text)
	//边界条件处理
	if m == 0 {
		return 0
	}
	if m > n {
		return -1
	}
	//偏移表初始化,这里使用range只能完成ascll字符串的匹配，如果想实现中文按字节匹配，要使用下标进行迭代，如果想实现中文匹配，就要使用map和[]rune
	shift := [256]int{}
	for i := range shift {
		shift[i] = m + 1
	}
	for i, v := range pat {
		shift[v] = m - i
	}
	//
	for i := 0; i <= n-m; {
		j := 0
		for j < m && pat[j] == text[i+j] {
			j++
		}
		//检查匹配的字符串
		if j == m {
			fmt.Println("match!")
			return i
		}
		if i >= n-m { //==号也可以
			return -1
		}
		move := shift[text[i+m]]
		i += move

	}
	return -1 //针对其它越界条件

}
