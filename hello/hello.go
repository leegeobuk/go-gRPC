package hello

import (
	"rsc.io/quote/v3"
)

// Hello to world
func Hello() string {
	return quote.HelloV3()
}

// Proverb for concurrency
func Proverb() string {
	return quote.Concurrency()
}