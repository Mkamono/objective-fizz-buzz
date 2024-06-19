package main

import (
	"fmt"

	"github.com/Mkamono/objective-fizz-buzz/fizzbuzz"
)

func main() {
	for i := 1; i <= 30; i++ {
		fmt.Printf("%d: %s\n", i, fizzbuzz.FizzBuzz(i))
	}
}
