package answer

import (
	"fmt"

	"github.com/Mkamono/objective-fizz-buzz/answer/fizzbuzz/core"
	"github.com/Mkamono/objective-fizz-buzz/answer/fizzbuzz/rule"
)

func Main() {
	fizzRule := rule.NewCyclicNumberRule(3, "Fizz")
	buzzRule := rule.NewCyclicNumberRule(5, "Buzz")
	passThroughRule := rule.NewPassThroughRule()

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
