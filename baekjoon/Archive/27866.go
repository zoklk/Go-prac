// 1줄 단위의 문자열을 읽어오고 싶다면 reader.ReadString('\n')
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	var n int
	fmt.Fscan(reader, &n)

	fmt.Printf("%c", line[n-1])
}
