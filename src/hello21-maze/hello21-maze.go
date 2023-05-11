package main

import (
	"fmt"
	"os"
)

// 广度优先算法走迷宫简单实现
/*
原始图，如下图左上角为起点，右下角为终点,只能走带0的格子，带1的格子表示墙
0 1 0 0 0
0 0 0 1 0
0 1 0 1 0
1 1 1 0 0
0 1 0 0 1
0 1 0 0 0
*/

/* 广度优先算法核心思路（注意：每个格子有三种状态（已发现未探索，已发现已探索，未发现））
1，用上左下右走法探索周边数据（比如上图起点0，上面没有数据，左边没有数据，下面是0，右边是1），走到那个格子那个就是中心点
2，如果格子可以走就，就为该格子标识一个数值，这个数值就是当前中心格子数值加1
*/

/*
经过算法标识后的图（最终将每个格子进行倒叙排列就是最优路径，最大的数值就是最少多少步能走完迷宫）
0       4   5   6
1   2   3       7
2       4       8
            10  9
        12  11
        13  12  13
*/

func main() {
	// 定义迷宫数据（只能走带0的格子，带1的格子表示墙）
	//data := [...][...]int{
	//	{0, 1, 0, 0, 0},
	//	{0, 0, 0, 1, 0},
	//	{0, 1, 0, 1, 0},
	//	{1, 1, 1, 0, 0},
	//	{0, 1, 0, 0, 1},
	//	{0, 1, 0, 0, 0},
	//}
	//fmt.Println("", data)
	file, err := os.Open("maze.file")
	if err != nil {
		fmt.Println("读取文件数据出现错误:", err)
	}
	var row, col int
	// 读取第一行数据,并将值填充到 &row和&col变量（就是行和列的数量）
	fmt.Fscanf(file, "%d %d", &row, &col)
	// 创建数据二维数组
	maze := make([][]int, row)
	// 填充二维数组行列数据
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	// 打印读取到的数据
	fmt.Println("定义迷宫数据（只能走带0的格子，带1的格子表示墙）")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
	// 开始广度优先算法走迷宫
	steps := walk(maze, point{0, 0}, point{row - 1, col - 1})
	fmt.Println("经过算法标识后的图（最终将每个格子进行倒叙排列就是最优路径，最大的数值就是最少多少步能走完迷宫）")
	for _, roww := range steps {
		for _, vall := range roww {
			if vall == 0 {
				fmt.Printf("%3s ", "")
			} else {
				fmt.Printf("%3d ", vall)
			}

		}
		fmt.Println()
	}
}

// 迷宫数据一个位置的坐标
type point struct {
	i, j int
}

// 通过当前点的位置计算“上左下右”某一个点的位置
func (p point) add(r point) point {
	return point{
		p.i + r.i,
		p.j + r.j,
	}
}

// 通过坐标获取迷宫图里面的数据
func (p point) at(maze [][]int) (int, bool) {
	// 判断是否越界
	if p.i < 0 || p.i >= len(maze) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(maze[p.i]) {
		return 0, false
	}
	return maze[p.i][p.j], true
}

// 上左下右4个坐标位置的计算
var dirs = [4]point{
	{-1, 0}, // 上边点的坐标相对当前中心点的行坐标减1，列坐标不变
	{0, -1}, // 左边点的坐标相对当前中心点的行坐标不变，列坐标减1
	{1, 0},  // 下边点的坐标相对当前中心点的行坐标加1，列坐标不变
	{0, 1},  // 右边点的坐标相对当前中心点的行坐标不变，列坐标加1
}

/*
走迷宫
maze   迷宫数据
start  开始坐标
enf    结束坐标
*/
func walk(maze [][]int, start point, end point) [][]int {
	// 创建经过算法标识后的数据图（最终将每个格子进行倒叙排列就是最优路径）
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	// 所有可以走的位置的数据
	q := []point{start}
	// 遍历所有可以走的位置数据
	for len(q) > 0 {
		// 当前走到的位置
		cur := q[0]
		// 截取数组从1开始因为第0个已经走了嘛
		q = q[1:]
		// 当前走到终点了就不用再走了
		if cur == end {
			break
		}
		// 从上左下右4个点循环探索数据看能不能走
		for _, dir := range dirs {
			// 下一个点的坐标位置等于当前节点加上"上左下右"其中某一个位置的计算数据
			next := cur.add(dir)
			// 原始迷宫图里面next坐标位置的数据是0并且在steps数组里面的值也是0（因为如果在steps数组里面不是0说明该位置已经走过一回了）并且 next != start（start是起点嘛也不能算）才可以走
			// 通过坐标从原始图里面获取数据
			val, ok := next.at(maze)
			// 如果没有获取到值或者值等于1说明该位置不能走（因为只能走值等于0的位置）
			if !ok || val == 1 {
				// 跳出单次循环看下一个点
				continue
			}
			// 通过坐标从经过算法标识后的数据图里面获取数据
			vall, okk := next.at(steps)
			// 如果没有获取到值或者值不等于0说明该位置已经走过一回了，不用再计算了
			if !okk || vall != 0 {
				// 跳出单次循环看下一个点
				continue
			}
			// 如果等于起点也不用再计算了，起点最开始就算过了嘛
			if next == start {
				continue
			}
			// 获取当前走到的位置的数值
			stepsValue, _ := cur.at(steps)
			// 现在是标识可走位置的值（如果格子可以走就，就为该格子标识一个数值，这个数值就是当前中心格子数值加1）
			steps[next.i][next.j] = stepsValue + 1
			// 将next位置添加到所有可以走的位置数据里面
			q = append(q, next)
		}
	}
	return steps
}
