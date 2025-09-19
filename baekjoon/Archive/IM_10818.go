// bufio로 reader를 만들면, fmt.Scan보다 훨씬 가볍게 접근 가능하다.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	min := 1000001
	max := -1000001

	var temp int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &temp)
		if temp < min {
			min = temp
		}
		if temp > max {
			max = temp
		}
	}
	fmt.Println(min, max)
}
