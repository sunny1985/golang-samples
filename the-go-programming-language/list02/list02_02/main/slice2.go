package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 10)
	fmt.Println("len(mySlice):", len(mySlice)) // 元素个数
	fmt.Println("cap(mySlice):", cap(mySlice)) // 分配空间大小
}
