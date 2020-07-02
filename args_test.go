package main
import "fmt"
import "os.Args"

func main() {
	if err != nil {
		return 
	}
	if len(os.Args) <= 1 {
		fmt.Println("Please provide atleast one argument")
		return
	}
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("Arg:" + i + " :" + os.Args[i])
	}
}