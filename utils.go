package main

import (
	"os"
	"math/rand"
	"time"
)

var Random *os.File

func init() {
	rand.Seed(time.Now().UnixNano())
}

const characters = "abcdefghijklmnopqrstuvwxyz0123456789"

func shortId(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}
	return string(b)
}