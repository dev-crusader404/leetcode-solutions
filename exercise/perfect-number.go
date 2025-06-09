package exercise

import "errors"

// Define the Classification type here.
type Classification string

const (
	ClassificationDeficient Classification = "deficient"
	ClassificationPerfect   Classification = "perfect"
	ClassificationAbundant  Classification = "abundant"
)

var ErrOnlyPositive error = errors.New("not positive number")

func Classify(n int64) (Classification, error) {
	if n < 1 {
		return "", ErrOnlyPositive
	}
	return getClassification(n)
}

func getClassification(n int64) (Classification, error) {
	var total int64
	for i := 1; i < int(n); i++ {
		if int(n)%i == 0 {
			total += int64(i)
		}
	}
	if total == n {
		return ClassificationPerfect, nil
	} else if total < n {
		return ClassificationDeficient, nil
	} else {
		return ClassificationAbundant, nil
	}
}
