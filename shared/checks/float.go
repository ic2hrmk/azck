package checks

import "github.com/pkg/errors"

func IsZeroOrPositiveFloat64() ZeroOrPositiveFloat64Rule {
	return ZeroOrPositiveFloat64Rule{}
}

type ZeroOrPositiveFloat64Rule struct{}

func (ZeroOrPositiveFloat64Rule) Validate(value interface{}) error {
	floatValue, ok := value.(float64)
	if !ok {
		return errors.New("passed value is no of float64 type")
	}

	if floatValue < 0 {
		return errors.New("value is negative")
	}

	return nil
}
