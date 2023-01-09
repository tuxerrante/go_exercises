// https://go.dev/tour/concurrency/8
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"sync"
)

func TreeWalkHandler(t *tree.Tree, ch chan int, wg *sync.WaitGroup) {

	wg.Add(1)
	go Walk(tree.New(1), ch, wg)

}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, wg *sync.WaitGroup) {
	fmt.Printf(" -- Current node: %d\n", t.Value)

	defer wg.Done()

	ch <- t.Value

	if (t.Left != nil && *t.Left != tree.Tree{}) {
		fmt.Printf(" <- going left (%d)\n", t.Left.Value)
		wg.Add(1)
		go Walk(t.Left, ch, wg)
	}
	if (t.Right != nil && *t.Right != tree.Tree{}) {
		fmt.Printf(" -> going right (%d)\n", t.Right.Value)
		wg.Add(1)
		go Walk(t.Right, ch, wg)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	// Communication over an unbuffered channel causes the sending and receiving goroutines to
	// synchronize. unbuffered channels are sometimes called synchronous channels.
	// When a value is sent on an unbuffered channel, the receipt of the value happns before the awakening of the sending goroutine.
	// ch := make(chan int)
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	TreeWalkHandler(tree.New(1), ch, &wg)

	wg.Wait()
	close(ch)

	for value := range ch {
		fmt.Printf("_ Received %d.\tLenght(channel)=%d\n", value, len(ch))
	}
}
