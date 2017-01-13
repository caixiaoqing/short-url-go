package utils

import (
	"bytes"
)

// characters used for conversion
const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// converts number to base62
func Encode(number int) string {
	if number == 0 {
		return string(alphabet[0])
	}

	strBase62 := ""
	length := len(alphabet)
	for number > 0 {
		strBase62 = string(alphabet[number%length]) + strBase62
		number = number / length
	}

	return strBase62
}

// converts base62 token to int
func Decode(token string) int {
	chars := []byte(alphabet)
	pow := len(chars)

	number := 0
	for _, c := range []byte(token) {
		number = number*pow + bytes.IndexByte(chars, c)
	}

	return number
}
