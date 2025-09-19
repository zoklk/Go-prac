package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(reader, &N, &M)
	box := make([]int, N)

	for i := 0; i < M; i++ {
		var a, b, c int
		fmt.Fscan(reader, &a, &b, &c)
		for j := a - 1; j < b; j++ {
			box[j] = c
		}
	}

	for i := 0; i < N; i++ {
		fmt.Printf("%d ", box[i])
	}
}
