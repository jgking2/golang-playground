package functions

const (
	randomNumberToCount = 45
	PublicConst         = 13
)

func NakedReturns(count int) (total int) {
	total = randomNumberToCount * count
	return
}
