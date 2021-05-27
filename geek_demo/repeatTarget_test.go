package geek_demo

import (
	"fmt"
	"testing"
)

func TestFindRepeatTarget(t *testing.T) {
	a := []int{1, 2, 3, 3, 4, 4, 4, 4, 4, 5, 5, 8, 10, 12}
	targets := []int{1, 4, 11, 12, 13}
	for _, target := range targets {
		result := findRepeatTarget(a, target)
		fmt.Println(result)
	}
}
