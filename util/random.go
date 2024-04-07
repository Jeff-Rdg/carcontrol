package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvxyz"
const numbers = "1234567890"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// randomInt generates a random integer between min and max
func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int, value string) string {
	var sb strings.Builder
	k := len(value)

	for i := 0; i < n; i++ {
		c := value[rand.Intn(k)]

		sb.WriteByte(c)
	}

	return sb.String()

}

func RandomName() string {
	return randomString(6, alphabet)
}

func RandomAddress() string {
	return "Rua " + randomString(15, alphabet)
}

func RandomCnpj() string {
	return randomString(14, numbers)
}

func RandomPhone() string {
	return randomString(11, numbers)
}

func RandomVacancy() uint {
	return uint(randomInt(1, 20))
}
