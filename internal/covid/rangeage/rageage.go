package rangeage

import "github.com/pitsanujiw/go-covid/internal/constant"

func FindRangeAge(age int) string {
	if age == 0 {
		return constant.UNKNOWN
	}
	if age >= 0 && age <= 30 {
		return constant.ADULT
	}
	if age >= 31 && age <= 60 {
		return constant.OLD
	}
	return constant.ELDER
}
