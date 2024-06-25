package spec

import (
	"fmt"

	"github.com/Mkamono/objective-fizz-buzz/answer/fizzbuzz/core"
)

// 倍数のときに指定した文字列を追加するReplaceRule
var _ core.ReplaceRule = (*cyclicNumberRule)(nil)

func NewCyclicNumberRule(divisor int, word string) *cyclicNumberRule {
	return &cyclicNumberRule{divisor: divisor, word: word}
}

type cyclicNumberRule struct {
	divisor int
	word    string
}

func (r *cyclicNumberRule) Match(carry string, n int) bool {
	return n%r.divisor == 0
}

func (r *cyclicNumberRule) Apply(carry string, n int) string {
	return carry + r.word
}

// 通常の数字をそのまま返すReplaceRule
var _ core.ReplaceRule = (*passThroughRule)(nil)

func NewPassThroughRule() *passThroughRule {
	return &passThroughRule{}
}

type passThroughRule struct {
}

func (r *passThroughRule) Match(carry string, n int) bool {
	return carry == ""
}

func (r *passThroughRule) Apply(carry string, n int) string {
	return carry + fmt.Sprint(n)
}
