package main

import "fmt"

func main() {
	//uch := make(chan int, 0)
	//uch <- 1
	//<-uch

	ch := make(chan int, 1)
	ch <- 1
	//<-ch

	m := make(map[int]int)
	m[0] = 0

	var i int
	var exists bool
	if i, exists = m[0]; !exists {
		fmt.Println("is not exit")
	}
	fmt.Println(i)
}
