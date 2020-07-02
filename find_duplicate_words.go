package main
import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)
// read files content and find duplicate words
func main(){
	//there should be one or more arguments
	if len(os.Args[1:]) < 1 {
		fmt.Println("Please provide at least one file")
		return
	}
	//map to stroe results
	counts := make(map[string]map[string]int)
	//Read files one by one
	for _,file := range os.Args[1:] {
		f,err := ioutil.ReadFile(file)
		// if there is an error to read file, throw error and move on
		if err != nil {
			fmt.Printf("%v\n%s",os.Stderr,err)
			continue
		}
		tempmap := make(map[string]int)
		for _,line := range strings.Split(string(f), " ") {
			tempmap[line]++
		}
		counts[file]=tempmap
	}
	for file,countmap := range counts{
		for word,count := range countmap{
			fmt.Printf("%s\t%s\t%d\n",file,word,count)
		} 
	}

}