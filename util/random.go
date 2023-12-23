package util

import (
	"math/rand"
)

// RandomString generates a random string
func RandomString(size int) (str string) {
	abcd := "abcdefghijklmnopqrstuvwxyzñ1234567890"
	for i := 0; i < size; i++ {
		str += string(abcd[rand.Intn(len(abcd))])
	}

	return str
}

// RandomBool generates a random boolean, true or false
func RandomBool() bool {
	n := rand.Intn(2)
	if n == 0 {
		return false
	} else {
		return true
	}
}
