package geek_demo

import (
	"math"
)

//func main() {
//	number := 10
//	squareRoot := getSquareRoot(number, 0.000001, 10000)
//	if squareRoot == -1.0 {
//		fmt.Println("请输入大于1的整数")
//	} else if squareRoot == -2.0 {
//		fmt.Println("未能找到解")
//	} else {
//		fmt.Println(fmt.Printf("%d的平方根是%f", number, squareRoot))
//	}
//}

func getSquareRoot(n int, deltaThreshold float64, maxTry int) float64 {
	if n <= 1 {
		return -1.0
	}

	min, max := float64(1.0), float64(n)

	for i := 0; i < maxTry; i++ {
		mid := (min + max) / 2
		square := mid * mid
		delta := math.Abs((square / float64(n)) - 1)
		if delta <= deltaThreshold {
			return mid
		} else {
			if square > float64(n) {
				max = mid
			} else {
				min = mid
			}
		}
	}

	return -2.0
}
