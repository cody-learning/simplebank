package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	min := int64(10)
	max := int64(100)

	for i := 0; i < 100; i++ {
		result := RandomInt(min, max)
		require.GreaterOrEqual(t, result, min)
		require.LessOrEqual(t, result, max)
	}

	result := RandomInt(50, 50)
	require.Equal(t, int64(50), result)
}

func TestRandomString(t *testing.T) {
	length := 10
	result := RandomString(length)
	require.Equal(t, length, len(result))

	for _, char := range result {
		require.Contains(t, alphabet, string(char))
	}

	for _, n := range []int{0, 1, 5, 20, 50} {
		result := RandomString(n)
		require.Equal(t, n, len(result))
	}

	result1 := RandomString(10)
	result2 := RandomString(10)

	if len(result1) > 0 {
		require.NotEqual(t, result1, result2)
	}
}

func TestRandomOwner(t *testing.T) {
	result := RandomOwner()
	require.Equal(t, 6, len(result))

	for _, char := range result {
		require.Contains(t, alphabet, string(char))
	}

	result1 := RandomOwner()
	result2 := RandomOwner()
	require.NotEqual(t, result1, result2)
}

func TestRandomMoney(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := RandomMoney()
		require.GreaterOrEqual(t, result, int64(0))
		require.LessOrEqual(t, result, int64(1000))
	}
}

func TestRandomCurrency(t *testing.T) {
	validCurrencies := map[string]bool{
		"USD": true,
		"EUR": true,
		"CAD": true,
	}

	for i := 0; i < 100; i++ {
		result := RandomCurrency()
		require.True(t, validCurrencies[result], "currency should be USD, EUR, or CAD")
	}

	foundCurrencies := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		result := RandomCurrency()
		foundCurrencies[result] = true
		if len(foundCurrencies) == 3 {
			break
		}
	}
	require.Equal(t, 3, len(foundCurrencies), "all three currencies should eventually be returned")
}
