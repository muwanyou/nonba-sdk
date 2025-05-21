package enum

// CyclePolarity 定义枚举字段.
type CyclePolarity string

// CyclePolarity 枚举值.
const (
	CyclePolarityPositive CyclePolarity = "positive"
	CyclePolarityNegative CyclePolarity = "negative"
)

// DefaultCyclePolarity 默认值
const DefaultCyclePolarity = CyclePolarityPositive

func (c CyclePolarity) String() string {
	return string(c)
}
