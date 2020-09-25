package letcodearray

func MovingCount(m int, n int, k int) int {
	//构建一个二维数组
	used := make([][]int, m)
	// 列空间的申请
	for i := range used {
		used[i] = make([]int, n)
	}
	return moveCount(m, n, 0, 0, used, k)
}

func moveCount(m, n, i, j int, used [][]int, k int) int {
	// 1.如果已经走过 直接返回
	// 2.如果已经越界 直接返回
	// 3.如果在这个范围 直接返回
	sum := i%10 + i/10 + j%10 + j/10
	if i < 0 || i >= m || j < 0 || j >= n || used[i][j] == 1 || sum > k {
		return 0
	}
	// 走过的标记
	used[i][j] = 1

	// 使用递归
	return 1 + moveCount(m, n, i-1, j, used, k) + moveCount(m, n, i+1, j, used, k) + moveCount(m, n, i, j+1, used, k) + moveCount(m, n, i, j-1, used, k)
}
