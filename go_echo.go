package main
import "fmt"
import "os"
import "strings"
import "time"

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please provide atleast one argument")
		return
	}
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Elepsed time: ", elapsed)
	start = time.Now()
	var s string
	for i := 1; i < len(os.Args); i++ {
		s = s + os.Args[i] + " "
	}
	fmt.Println(s)
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Println("Elepsed time: ", elapsed)
}
// Run with one or more argument
//go run args.go