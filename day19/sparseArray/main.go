package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
	"strconv"
)

type ChessMapNode struct {
	Row int
	Col int
	Val int
}

func writeSparceArr() { 
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	var sparceArr []ChessMapNode
	valNode := ChessMapNode{
		Row : 11,
		Col : 11,
		Val : 0,
	}
	sparceArr = append(sparceArr, valNode)	
	for i,v1 := range chessMap {
		for j, v2 := range v1 {
			if v2 != 0 {
				valNode = ChessMapNode{
					Row : i,
					Col : j,
					Val : v2,
				}
				sparceArr = append(sparceArr, valNode)	
			}
		}
	}
	fmt.Println("sparceArr=", sparceArr)
	//将稀疏数组保存在文件中
	//1.创建并打开文件
	file, err := os.OpenFile("C:/Users/duanfuhao/go/src/learngo/day19/sparseArray/file.txt",
	os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666) // For read access.
	if err != nil {
		fmt.Println("打开文件错误，err=",err.Error())
	}
	defer file.Close()

	//2.buf写入文件
	writer := bufio.NewWriter(file)
	var info string
	for _, v := range sparceArr {
		info += fmt.Sprintf("%v %v %v\r\n",v.Row,v.Col, v.Val)
	}
	_, err = writer.WriteString(info)
	if err != nil {
		fmt.Println("写入文件错误，err=",err.Error())
	}
	writer.Flush()
	//fmt.Println("file=", file)
}

//打开稀疏数组文件
func readSparceArr() {
	file, err := os.OpenFile("C:/Users/duanfuhao/go/src/learngo/day19/sparseArray/file.txt",
	os.O_RDWR, 0666) // For read access.
	if err != nil {
		fmt.Println("打开文件错误，err=",err.Error())
	}
	defer file.Close()

	//2.一行行读取文件内容
	//var str []byte
	reader := bufio.NewReader(file)
	flag := 0
	var chessMap [11][11]int
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.Trim(line, "\r\n")
		lineArr := strings.Split(line, " ")
		row, _ := strconv.Atoi(lineArr[0])
		col, _ := strconv.Atoi(lineArr[1])
		val, _ := strconv.Atoi(lineArr[2])
		if flag == 0 {
			for i := 0; i < row; i++ {
				for j := 0; j < col; j++ {
					chessMap[i][j] = val
				}
			}
		} else {
			chessMap[row][col] = val
		}
		flag++
		//fmt.Printf("lineArr %T, %v\n",lineArr[0], lineArr[0])
	}
	fmt.Println(chessMap)
}

func initArr() {
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}
}

func main()  {
	//writeSparceArr()
	readSparceArr()
}
