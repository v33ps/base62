package base62

import (
	"bytes"
	"math"
)

// alphabet for base62
const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Encode takes in an integer and returns it represented as a base62 string
func Encode(id int) string {
	if id == 0 {
		return string(alphabet[0])
	}
	var digits []int
	chars := make([]byte, 0)
	for id > 0 {
		remainder := id % 62
		digits = append(digits, remainder)
		id = id / 62
	}
	// reverse the list
	// start in the middle of the list and swap the middle element and middle-1
	// then work your way to either ends of the list, doing the same thing.
	// i.e.:  a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	// Swap a[4] & a[5], a[3] & a[6], a[2] & a[7], a[1] & a[8], a[0] & a[9]
	for i := len(digits)/2 - 1; i >= 0; i-- {
		opp := len(digits) - 1 - i
		digits[i], digits[opp] = digits[opp], digits[i]
	}

	// now that we have our base converted values, relate them
	// back to the base62 alphabet
	for _, v := range digits {
		chars = append(chars, alphabet[v])
	}
	return string(chars)
}

// Decode takes a base62 string and returns it as an integer
func Decode(str string) int {
	number := 0
	counter := 0.0
	chars := []byte(alphabet)

	charsLength := float64(len(chars))
	tokenLength := float64(len(str))

	// go through each character in the string
	// find it's position in our alphabet
	// convert it back to base 10
	for _, c := range []byte(str) {
		power := tokenLength - (counter + 1)
		index := bytes.IndexByte(chars, c)
		number = number + index*int(math.Pow(charsLength, power))
		counter = counter + 1
	}
	return number
}
