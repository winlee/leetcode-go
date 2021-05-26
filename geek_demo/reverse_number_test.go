package geek_demo

import (
	"fmt"
	"testing"
)

func TestReverseNumber(t *testing.T) {
	var a int64 = 1516000
	result := reverseNumber(a)
	fmt.Println(result)
}
