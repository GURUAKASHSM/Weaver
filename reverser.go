package main

import (
	"context"
	"log"
	"strings"

	"github.com/ServiceWeaver/weaver"
)
//
// Reverser component.
type Reverser interface {
	Reverse(context.Context, string) (string, error)
}

// Reverser component.
type AddSpacer interface {
	AddSpace(context.Context, string) (string, error)
}

// Implementation of the Reverser component.
type reverser struct {
	weaver.Implements[Reverser]
}

// Implementation of the Reverser component.
type addspacer struct {
	weaver.Implements[AddSpacer]
}

func (r *reverser) Reverse(_ context.Context, s string) (string, error) {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes), nil
}

func (r *addspacer) AddSpace(_ context.Context, str string) (string, error) {

	log.Println("Addspace reverser")
	var result strings.Builder

	// Iterate over each character in the string
	for i, char := range str {
		// Append the current character to the result string
		result.WriteRune(char)

		// If it's not the last character, append a space
		if i < len(str)-1 {
			result.WriteString(" ")
		}
	}

	return result.String(), nil
}
