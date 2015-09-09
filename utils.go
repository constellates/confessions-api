package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

/**
 * Short Id
 *
 * Takes an integer n and returns a
 * random alpha-numeric string of
 * length n.
 */
func shortId(n int) string {
	const characters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}
	return string(b)
}