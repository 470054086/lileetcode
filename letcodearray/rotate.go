package letcodearray

/**
给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。

不占用额外内存空间能否做到？

 

示例 1:

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
示例 2:

给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-matrix-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


/**
0 0 						04   news[j][4-0-1] = matrix[0][0]
 5  1  9 11              x  x  x  5
 x  x  x  x   =旋转后=>   x  x  x  1
 x  x  x  x              x  x  x  9
 x  x  x  x              x  x  x 11

x  x  x  x              x  x  2  x
 2  4  8 10   =旋转后=>   x  x  4  x
 x  x  x  x              x  x  8  x
 x  x  x  x              x  x 10  x

// 翻译为代码为

news[j][len(matrix)-i-1] = matrix[i][j]
 */
func Rotate(matrix [][]int) {
	// 构造一个二位数组
	// 行空间申请
	news := make([][]int, len(matrix))
	// 列空间的申请
	for i := 0; i < len(news); i++ {
		news[i] = make([]int, len(matrix))
	}
	// 循环交换
	for i, v := range matrix {
		for j, _ := range v {
			news[j][len(matrix)-i-1] = matrix[i][j]
		}
	}
	copy(matrix, news)
}

/**
	// 最终旋转的公式
	news[col][n-row-1] = matrix[row][col]
	// 为了改变原来位置的值  我们使用一个临时变量
	temp = news[col][n-row-1]
	news[col][n-row-1] = matrix[row][col]


 */