package map_ext

import (
	"strings"
	"testing"
	"strconv"
)

func TestMapForset(t *testing.T) {

	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(strings.Join(parts,"-"))
	// t.Logf("%T",strings.Split(s,","))
}
func TestStrconv(t *testing.T){
	s:=strconv.Itoa(10)
	t.Log("str"+s+"\n")
	if i,err:= strconv.Atoi(s);err==nil{
		t.Log(i)
	}
	
}
