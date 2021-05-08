package geek_demo

import (
	"fmt"
	"testing"
)

//
func TestCoinChanges(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	result := coinChange(coins, amount)
	fmt.Println(result)
	if result == 3 {
		t.Log(result)
	} else {
		t.Fail()
	}

	var a = []struct {
		coins  []int
		amount int
	}{
		{[]int{1, 2, 5}, 10},
		{[]int{1}, 10},
		{[]int{1}, 2},
		{[]int{1}, 0},
		{[]int{2}, 3},
	}

	for _, v := range a {
		fmt.Println(coinChange(v.coins, v.amount))
	}
}
