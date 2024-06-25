package core

type ReplaceRule interface {
	Match(carry string, n int) bool
	Apply(carry string, n int) string
}

type NumberConverter interface {
	Run(n int) string
}

var _ NumberConverter = (*numberConverter)(nil)

type numberConverter struct {
	rules []ReplaceRule
}

func NewNumberConvter(rules []ReplaceRule) *numberConverter {
	return &numberConverter{rules: rules}
}

func (c *numberConverter) Run(n int) string {
	var result string
	for _, rule := range c.rules {
		if rule.Match(result, n) {
			result = rule.Apply(result, n)
		}
	}
	return result
}
