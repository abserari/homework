package main

import "fmt"

//针对奇数矩阵，即中位数不需要计算的情况 如：1，3，5，7
func main() {
	var n int
	fmt.Scan(&n)

	between := n / 2
	fmt.Println(between*n + between*(n-between) + 1)
}
