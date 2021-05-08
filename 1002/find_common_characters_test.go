package _002

import (
	"fmt"
	"testing"
)

func TestCommonChars(t *testing.T) {
	strs := []string{
		"bella",
		"label",
		"roller",
	}

	result := commonChars(strs)
	fmt.Println(result)
}

func TestCommonChars2(t *testing.T) {
	strs := []string{
		"bella",
		"label",
		"roller",
	}

	result := commonChars2(strs)
	fmt.Println(result)
}

