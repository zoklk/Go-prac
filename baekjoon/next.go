package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	x := make([]int, a)
	for i := 0; i < a; i++ {
		var temp int
		fmt.Scan(&temp)
		x = append(x, temp)
	}
	for i := 0; i < a; i++ {

	}
}
