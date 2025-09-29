// scanner로 입력받으면 문자열, 숫자열로 전환하려면 int(s - '0')
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	n := scanner.Text()

	sum := 0
	for i := 0; i < k; i++ {
		sum += int(n[i] - '0')
	}

	fmt.Fprintln(writer, sum)
}
