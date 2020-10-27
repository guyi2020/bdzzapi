package main

import (
	"fmt"
	"time"
)

// defer 延迟调用函数

func main() {
	// 先进后出，和栈一样
	defer fmt.Println("最后执行...")
	defer outFunc()
	fmt.Println("执行...")
	/*
		执行...
		outFunc执行...
		outFunc最后执行...
		最后执行...
	*/

	for i := 0; i < 10; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}

// 外围函数
func outFunc() {
	defer fmt.Println("outFunc最后执行...")
	fmt.Println("outFunc执行...")

	// ?before=2020-11-12T23:59:59&after=2020-10-26T00:00:00
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println(timeStr)
}
