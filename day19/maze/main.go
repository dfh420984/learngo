package main

import (
	"fmt"
)

func setWay(myMap *[8][7]int, i int, j int) bool {
	//先假定一个出路
	if myMap[6][5] == 2 {
		return true
	} else {
		if myMap[i][j] == 0 { //如果这个点是可以探测的
			//假设需要探测的方向是下右上左
			myMap[i][j] = 2            //假设这个点是出路
			if setWay(myMap, i+1, j) { //下
				return true
			} else if setWay(myMap, i, j+1) { //右
				return true
			} else if setWay(myMap, i-1, j) { //上
				return true
			} else if setWay(myMap, i, j-1) { //左
				return true
			} else { //死路
				myMap[i][j] = 3
				return false
			}
		} else { //说明这个点不能探测
			return false
		}
	}
}

//迷宫练习
func main() {
	//创建一个二维数组，模拟迷宫
	var myMap [8][7]int //构建迷宫地图 0.是没有走过的点 1.是墙；2.是出路；3.是走过的路，死路
	//1.先把地图的最上和最下设置 为1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	//2.在把地图最左边和右边设置为1
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}

	//3.设置几个障碍
	// myMap[1][2] = 1
	// myMap[2][2] = 1
	myMap[3][1] = 1
	myMap[3][2] = 1

	//4.输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("开始探测地图")
	setWay(&myMap, 1, 1)
	fmt.Println("开始探测地图完毕")
	//5.输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}
