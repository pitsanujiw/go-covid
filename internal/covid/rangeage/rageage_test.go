package rangeage

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/pitsanujiw/go-covid/internal/constant"
)

func TestFindRangeAge(t *testing.T) {
	t.Run("should return unknown", func(t *testing.T) {
		got := FindRangeAge(0)

		require.Equal(t, got, constant.UNKNOWN)
	})

	t.Run("should return adult", func(t *testing.T) {
		got := FindRangeAge(20)

		require.Equal(t, got, constant.ADULT)
	})

	t.Run("should return old", func(t *testing.T) {
		got := FindRangeAge(35)

		require.Equal(t, got, constant.OLD)
	})

	t.Run("should return elder", func(t *testing.T) {
		got := FindRangeAge(80)

		require.Equal(t, got, constant.ELDER)
	})
}
