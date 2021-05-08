package _003

import (
	"fmt"
	"testing"
)

func TestIsValid1003(t *testing.T) {


	s := []string{
		"abcabcababcc",
		"abccba",
	}

	for _, s2 := range s {
		isValid := isValid1003(s2)
		fmt.Println(isValid)
	}
}
