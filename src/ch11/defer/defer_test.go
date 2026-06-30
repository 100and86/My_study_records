package defer_test

import (
	"testing"
	"fmt"
)

func Clear() {
	fmt.Print("clear resources")

}
func TestDefer(t *testing.T){
	defer Clear()
	fmt.Println("kaishi")
	panic("err")
	

}
