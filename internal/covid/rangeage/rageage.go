package rangeage

import "github.com/pitsanujiw/go-covid/internal/constant"

// Calculate simple for get validate range of age
func FindRangeAge(age int) string {
	// in case unknown age
	if age == 0 {
		return constant.UNKNOWN
	}

	// in case adult
	if age >= 0 && age <= 30 {
		return constant.ADULT
	}

	// in case older
	if age >= 31 && age <= 60 {
		return constant.OLD
	}

	// and the last in case elder, an age more than 60+
	return constant.ELDER
}
