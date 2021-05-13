package geek_demo

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestThreeSum2(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum2(nums)
	fmt.Println(res)
}

func TestMutexVariable(t *testing.T) {
	var i int32 = 0
	for j := 0; j < 10000; j++ {
		go func() {
			atomic.AddInt32(&i, 1)
		}()
	}

	time.Sleep(time.Second * 3)
	fmt.Println(i)
	if i == 10000 {
		t.Log("success")
	} else {
		t.Failed()
	}
}
