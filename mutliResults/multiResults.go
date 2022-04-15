package main

import "fmt"

func swap(y, x string) (string, string){
	return x, y
}


func main(){
	fmt.Println(swap("test","two"))
}