// golang_for_calc
package main

import (
	"fmt"
)

func main() {
	var choose string
	fmt.Println("计算器（1）\n", "数组（2）\n")
	fmt.Scanln(&choose)
	if choose == "1" {
		calc()
	}
	if choose == "2" {
		NumTeam()
	}
}

func NumTeam() {

	var YN string
	var t1 int
	nt := []int{}
	fmt.Printf("第一个数组的第一个数字为:")
	fmt.Scanln(&t1)
	nt = append(nt, t1)
	fmt.Println(nt)

	fmt.Printf("是否继续添加数组(y/n)):")
	fmt.Scanln(&YN)

	if YN == "y" {
	YN:
		var add1 int
		fmt.Printf("input:")
		fmt.Scanf("%v", &add1)
		nt = append(nt, add1)
		fmt.Println(nt)

		fmt.Printf("是否继续添加数组(y/n)):")
		fmt.Scanln(&YN)
		if YN == "y" {
			goto YN //第goto：26行
		}
	} else {
		fmt.Println(nt)
	}

	fmt.Println(nt)
}

func calc() {
	var num1 float64
	var num2 float64
	var sf = "zxc" //符号
	for {
		fmt.Println("在sf里输入'quit'可以退出")
		fmt.Println("")
		fmt.Printf("input num1:")
		fmt.Scanf("%v", &num1)
		fmt.Printf("input sf:")
		fmt.Scanln(&sf)

		//quit
		if sf == "quit" {
			break
		}
		//quit

		fmt.Printf("input num2:")
		fmt.Scanf("%v", &num2)
		if sf == "+" {
			fmt.Printf("%v\n", num1+num2)
		}
		if sf == "-" {
			fmt.Printf("%v\n", num1-num2)
		}
		if sf == "/" {
			fmt.Printf("%v\n", num1/num2)
		}
		if sf == "*" {
			fmt.Printf("%v\n", num1*num2)
		}
	}
}


