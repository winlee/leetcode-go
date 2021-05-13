package geek_demo

import (
	"fmt"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	result := minimumTotal(triangle)
	fmt.Println(result)
	if result != 11 {
		t.Errorf("failed")
	}
}
