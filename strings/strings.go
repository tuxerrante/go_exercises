package main
import (
	"fmt"
	"strings"
	"myStrings/runesExamples"
)

/*
*  https://www.calhoun.io/6-tips-for-using-strings-in-go/
*/
func stringBuildExample(){

	// This is more efficient if you are combining lots of strings
	var sb strings.Builder
	sb.WriteString("abc")
	sb.WriteString("def")
	fmt.Printf("String: %s", sb.String())	
}

func main (){
	
	runesExamples.PrintRunesExamples()
}