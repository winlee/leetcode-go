package leetcode

//
//
/*
✔️外层物品，内层容量：先把石头1试完，再试石头2 → 顺序固定 → 组合
外层容量，内层物品：每个容量都问一遍所有石头 → 顺序自由 → 排列
*/

func lastStoneWeightII(stones []int) int {
    // 你的代码
	sum := 0
	for _,stone := range stones {
		sum += stone
	}

	target := sum/2

	dp:=make([]int,target+1)
	for _, stone := range stones { // 先物品
		for j := target; j >= stone; j-- {   // 后背包，倒序
            dp[j] = max(dp[j], dp[j-stone]+stone)
        }
	}

	return sum-2*dp[target]
}