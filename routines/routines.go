package main

import (
	"log"
	"time"
)


func SayGreetings(s string, t int){
	for i:=0; i<t; i++ {
		
		log.Println(s)
		//  the print functions in the log standard package are synchronized so the texts printed by the two goroutines will not be messed up in one line
		// fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}	
}

func Routines(){
	
	go SayGreetings("Hi!", 10)
	go SayGreetings("Hello!", 10)
	go SayGreetings("Guten Tag!", 10)
	go SayGreetings("Ciao!", 10)
	
	time.Sleep(2 * time.Second)
}