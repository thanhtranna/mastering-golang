package main

import (
	"container/ring"
	"fmt"
)

var size int = 10

func main() {
	myRing := ring.New(size + 1)
	fmt.Println("empty ring:", *myRing)

	for i := 0; i < myRing.Len()-1; i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}
	myRing.Value = 2

	sum := 0

	myRing.Do(func(i interface{}) {
		t := i.(int)
		sum += t
	})

	fmt.Println("Sum: ", sum)

	for i := 0; i < myRing.Len(); i++ {
		myRing = myRing.Next()
		fmt.Print(myRing.Value, " ")
	}

	fmt.Println()
}
