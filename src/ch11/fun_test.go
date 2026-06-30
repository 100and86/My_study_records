package ch11

//可变参数练习
import (
	"fmt"
	"testing"
)

func TestFun(t *testing.T) {

	fmt.Printf("%d", sum(1, 2, 3, 4, 5))
}

func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}
