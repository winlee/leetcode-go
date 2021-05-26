package geek_demo

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	t1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(t1)
	fmt.Println(result)
}
