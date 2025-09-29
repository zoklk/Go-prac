// 숫자 문자 섞어서 혹은 단어 단위로 읽으려면 scanner 쓰기
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()

		fmt.Printf("%c%c\n", s[0], s[len(s)-1])
	}
}
