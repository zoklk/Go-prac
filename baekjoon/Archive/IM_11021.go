// 입출력 순서만 정확하면 되서, 한번에 입력받고 출력할 필요가 없다.

package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		fmt.Printf("Case #%d: %d\n", i, a+b)
	}
}
