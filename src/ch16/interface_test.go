package interface_test

import(
	
	"os"
	"testing"
)





func TestOs(t *testing.T) {
	os.Stdout.Write([]byte("hello,tom"))
}