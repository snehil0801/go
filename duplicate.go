package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line,n := range counts {
		if n > 1 {
			fmt.Println(line + ": " + strconv.Itoa(n))
		}
	}
}