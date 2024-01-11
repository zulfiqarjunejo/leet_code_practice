package main

import (
	"strconv"
)

// v1
// Runtime: 22 ms Memory: 4.48 MB
// func reverseString(str string) string {
//     if (str == "") {
//         return ""
//     }

//     chars := []rune(str)
//     for i, j := 0, len(chars) - 1; i < j; i, j = i+1, j-1 {
//         chars[i], chars[j] = chars[j], chars[i]
//     }

//     return string(chars)
// }

// func isPalindrome(x int) bool {

//     // 1. Naive approach

//     // convert the integer to string.
//     t := strconv.Itoa(x);

//     // reverse the string and store in a different variable.
//     tRev := reverseString(t)

//     // check if original string is equal to the reversed string.
//     return t == tRev;
// }

// v2
// Runtime: 8 ms Memory: 4.47 MB
// func IsPalindrome(x int) bool {
// 	result := true

// 	t := strconv.Itoa(x)
// 	n := len(t)

// 	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
// 		if t[i] != t[j] {
// 			result = false
// 		}
// 	}

// 	return result
// }

// v3
// Runtime: 4 ms Memory: 4.50 MB
func IsPalindrome(x int) bool {
	t := strconv.Itoa(x)
	n := len(t)

	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
		if t[i] != t[j] {
			return false
		}
	}

	return true
}
