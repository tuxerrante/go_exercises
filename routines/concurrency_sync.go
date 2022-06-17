package main

import (
	"log"
	"math/rand"
	"time"
	"sync"
)

var wg sync.WaitGroup

func CS_SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		time.Sleep(10 * time.Millisecond)
	}
	// Notify a task is finished.
	wg.Done() // <=> wg.Add(-1)
}

func CS_SayGreetings2(s string){
	wg.Add(1)
	log.Println(s)
	wg.Done()
}

func ConcurrencySync() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2) // register two tasks.
	go CS_SayGreetings("hi!", 10)
	go CS_SayGreetings("hello!", 10)

	go CS_SayGreetings2("hello 2!")

	go func(){
		log.Println("Panic?")
		wg.Done()
	}()

	// the main goroutine enters the blocking state.
	// block until all tasks are finished.
	wg.Wait()
}