package main

import (
	"bufio" // to read input
	"fmt"   // Package for printing
	"os"    // argument and file operation
)

func main() {
	counts := make(map[string]int) //Define map which will store the result
	files := os.Args[1:]          // Check if there are any arguments provided
	if len(files) >= 1  {           // if there are arguments then treat them as input file
		for _, arg := range files { // open files one by one
			f, err := os.Open(arg)
			if err != nil { // Check if there is any error while opening a file
				fmt.Println(os.Stderr, "dup2: %v\n", err) //Print error along with error statement from stderr
				continue                                  // Continue to next file
			}
			countLines(f, counts) //Call local method to populate map
			f.Close()             //close open file
		}
	} else {
		countLines(os.Stdin, counts) //if files are not provided as arguments read from stdin
	}
	for line, n := range counts { //put range on map and print out result
		fmt.Printf("%d\t%s\n", n, line)
	}
}
func countLines(f *os.File, counts map[string]int) { //Local method to populate final map
	input := bufio.NewScanner(f) //Scanner to open file pointer
	for input.Scan() {           //Scan input and update map value
		counts[input.Text()]++
	}
}
