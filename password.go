package utils

import (
	"math/rand"
	"time"
)

// GenerateRandomPassword generates a random string of upper + lower case alphabets and digits
// which is 23 bits long and returns the string
func GetRandomPassword() string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" + digits
	length := 23
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)
}
