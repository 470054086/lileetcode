package letcodearray

/**
给定一个包含了一些 0 和 1 的非空二维数组 grid 。

一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)

 

示例 1:

[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。

示例 2:

[[0,0,0,0,0,0,0,0]]
对于上面这个给定的矩阵, 返回 0。

 

注意: 给定的矩阵grid 的长度和宽度都不超过 50。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-area-of-island
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/*

官方解释

我们想知道网格中每个连通形状的面积，然后取最大值。

如果我们在一个土地上，以 4 个方向探索与之相连的每一个土地（以及与这些土地相连的土地），那么探索过的土地总数将是该连通形状的面积。

为了确保每个土地访问不超过一次，我们每次经过一块土地时，将这块土地的值置为 0。这样我们就不会多次访问同一土地。

 */

/**
  1. 将矩阵看成一个上下左右的数据 当前的点和上下左右进行比较
  2. 走过的点 都置为0 将不会再走
 */

func MaxAreaOfIsland(grid [][]int) int {
	area := 0
	for i, v := range grid {
		for j, _ := range v {
			area = max(dfs(grid, i, j), area)
		}
	}
	return area
}

func dfs(grid [][]int, i, j int) int {
	// 如果是最前面和最后的话 都是0
	if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] != 1 {
		return 0
	}
	// 定义初始为0
	grid[i][j] = 0
	// 定义走过了一次
	ans := 1
	// 判断附近是否满足
	di := [4]int{0, 0, 1, -1}
	dj := [4]int{1, -1, 0, 0}
	for index := 0; index != 4; index++ {
		index_i := i + di[index]
		index_j := j + dj[index]
		ans += dfs(grid, index_i, index_j)
	}
	return ans

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}




/**
	更简单的方式

 */

func maxAreaOfIsland(land [][]int) int {
	ans := 0
	for i, v := range land {
		for j, _ := range v {
			ans = max(bcf(land, i, j), ans)
		}
	}
	return  ans
}

func bcf(land [][]int, i, j int) int {
	// 判断递归的灵界条件
	if i < 0 || i >= len(land) || j < 0 || j >= len(land[0]) || land[i][j] == 0 {
		return 0
	}
	// 走过的路定义为0
	land[i][j] = 0
	// 像四个方向 不同的遍历
	return 1 + bcf(land, i-1, j) + bcf(land, i+1, j) + bcf(land, i, j-1) + bcf(land, i, j+1)
}

