// This package contains the tools that frequently used by other packages
package tools

import (
	"math/rand"
	"time"
)

/**
* This function receives a pair of min and max float64 numbers
* and generates a random number in that range (inclusive)
 */
func RandomNumberF(rangeLower float64, rangeUpper float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rangeLower + rand.Float64()*(rangeUpper-rangeLower)
}

/**
* This function receives a pair of min and max int64 numbers
* and generates a random number in that range (inclusive)
 */
func RandomNumberI(rangeLower int64, rangeUpper int64) int64 {

	rand.Seed(time.Now().UnixNano())
	return rangeLower + rand.Int63n(rangeUpper-rangeLower+1)
}

/**
* This function receives an slice of string and a needle
* and searches for the needle in the slice and if it finds it,
* it removes it from the slice and return the updated slice
 */
func SliceItemRemove(haystack []string, needle string) []string {
	for i, val := range haystack {
		if val == needle {
			return append(haystack[:i], haystack[i+1:]...)
		}
	}
	return haystack
}

/**
* This function receives an slice of string and an index then
* it removes it from the slice and return the updated slice
 */
func SliceItemRemoveByIndex(slice []string, index int64) []string {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
