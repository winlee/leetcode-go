package geek_demo

//1734. 解码异或后的排列
//给你一个整数数组 perm ，它是前 n 个正整数的排列，且 n 是个 奇数 。
//
//它被加密成另一个长度为 n - 1 的整数数组 encoded ，满足 encoded[i] = perm[i] XOR perm[i + 1] 。比方说，如果 perm = [1,3,2] ，那么 encoded = [2,1] 。
//
//给你 encoded 数组，请你返回原始数组 perm 。题目保证答案存在且唯一。
//
//
//
//示例 1：
//
//输入：encoded = [3,1]
//输出：[1,2,3]
//解释：如果 perm = [1,2,3] ，那么 encoded = [1 XOR 2,2 XOR 3] = [3,1]
//示例 2：
//
//输入：encoded = [6,5,4,6]
//输出：[2,4,1,5,3]
//
//
//提示：
//
//3 <= n < 10^5
//n 是奇数。
//encoded.length == n - 1
// encoded[0] = perm[0]^perm[1]
// encoded[1] = perm[1]^perm[2]
// encoded[2] = perm[2]^perm[3]
// encoded[3] = perm[3]^perm[4]
// .. odd = encoded[1] ^encoded[3] = perm[1]^perm[2]^perm[3]^perm[4]
// 因为perm是n全排列 所以 all = 1^2^3^n -> perm[0] = all ^ odd
func decode(encoded []int) []int {
	n := len(encoded) + 1
	all := 0
	for i := 1; i <= n; i++ {
		all ^= i
	}

	odd := 0
	for i := 0; i < n-1; i += 2 {
		odd ^= encoded[i]
	}

	perm := make([]int, n)
	perm[0] = all ^ odd
	for i, v := range encoded {
		perm[i+1] = perm[i] ^ v
	}

	return perm
}
