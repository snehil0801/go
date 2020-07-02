package main
import (
	"fmt"
	"os"
	"strconv"
)
func main() {
	//First Example with simeple three statment for loop
	for i:=1; i<len(os.Args); i++ {
		fmt.Println("Arg " + strconv.Itoa(i) + ": " + os.Args[i])
	}

	//Second Example with range function
	for _,arg := range os.Args[1:] {
		fmt.Println("Arg " + ": " + arg)
	}

	// Third with index with range func
	for i,arg := range os.Args[1:] {
		fmt.Println("Arg " + strconv.Itoa(i+1) + ": " + arg)
	}

	//While loop Example
	count := 1
	for count < len(os.Args){
		fmt.Println("Arg wth while: " + os.Args[count])
		count++
	}

}
