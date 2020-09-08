package letcodearray

import "math"


/**
https://leetcode-cn.com/problems/container-with-most-water/solution/

给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器，且 n 的值至少为 2。

输入：[1,8,6,2,5,4,8,3,7]
输出：49

两个柱子，找柱子高一点的，宽一点的，能存水最多的，
*/

/**
暴力求解
*/
func MaxArea(height []int) int {
	max := 0
	length := len(height)
	for i := 0; i < length; i++ {
		// 面积为宽度*高度 选取两个最小的高度 * 宽度
		for j := i + 1; j < length; j++ {
			// 高的话 只能选择最低的 最高的话 水就会溢出
			area := int(math.Min(float64(height[i]), float64(height[j]))) * (j - i)
			if area > max {
				max = area
			}
		}
	}
	return max
}
/**
使用双边指针
*/
func MaxAreaDoublePoint(height []int) int {
	max := 0
	i, j := 0, len(height)-1
	// 只要柱子没有走到中间 就一直循环
	for i < j {
		// 高的话 只能选择最低的 最高的话 水就会溢出
		area := min(height[i], height[j]) * (j - i)
		// 获取最大值
		max = maxInt(area, max)
		// 如果左边的柱子低于右边的柱子  则左边的柱子移动
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
