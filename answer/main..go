package answer

import (
	"fmt"

	"github.com/Mkamono/objective-fizz-buzz/answer/fizzbuzz/core"
	"github.com/Mkamono/objective-fizz-buzz/answer/fizzbuzz/spec"
)

func Main() {
	fizzRule := spec.NewCyclicNumberRule(3, "Fizz")
	buzzRule := spec.NewCyclicNumberRule(5, "Buzz")
	passThroughRule := spec.NewPassThroughRule()

	fizzbuzzConverter := core.NewNumberConvter(
		[]core.ReplaceRule{
			fizzRule,
			buzzRule,
			passThroughRule,
		})
	for i := 1; i <= 30; i++ {
		fmt.Printf("%d: %s\n", i, fizzbuzzConverter.Run(i))
	}
}
