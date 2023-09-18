package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// The RandomInt function generates a random integer between a given minimum and maximum value.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// The RandomString function generates a random string of length n using characters from the alphabet.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// The RandomOwner function returns a randomly generated string of length 6.
func RandomOwner() string {
	return RandomString(6)
}

// The RandomMoney function returns a random integer between 0 and 1000.
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// The RandomCurrency function returns a random currency from a predefined list.
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}

	n := len(currencies)

	return currencies[rand.Intn(n)]
}
