package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	a := 1
	b := 1
	n := 5
	for i := 0; i < n; i++ {
	fmt.Printf("%d\t", a)
		a, b = b, a+b
	}
}
