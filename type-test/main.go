package main
// DOCS:
// - https://go.dev/blog/declaration-syntax

import (
	"container/list"
	"fmt"
)

func main() {
	x := make(map[string]*list.List)
	x["key"] = list.New()
	x["key"].PushBack("value")
	fmt.Println("x[\"key\"].Front().Value = ", x["key"].Front().Value)

	// ---------------------------------------------------------------------------------------
	mapExamples()

	// --- call external examples ---
	// stringTests()

	arrayTest()

	// sliceTests()
}
