package fizzbuzz

import (
	"github.com/Mkamono/objective-fizz-buzz/answer/fizzbuzz/core"
	"github.com/Mkamono/objective-fizz-buzz/answer/fizzbuzz/spec"
)

func FizzBuzz(n int) string {
	fizzRule := spec.NewCyclicNumberRule(3, "Fizz")
	buzzRule := spec.NewCyclicNumberRule(5, "Buzz")
	passThroughRule := spec.NewPassThroughRule()

	fizzbuzzConverter := core.NewNumberConverter([]core.ReplaceRule{
		fizzRule,
		buzzRule,
		passThroughRule,
	})

	return fizzbuzzConverter.Run(n)
}
