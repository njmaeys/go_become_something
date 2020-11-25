package main

import (
	"fmt"
)

// dummyString returns a dummy value.
func dummyString(valToReturn string) string {
	return valToReturn
}

func main() {
	fmt.Println("Go become something!")

	dummy := dummyString("dummy string")
	fmt.Println(dummy)
}
