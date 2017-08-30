package main

import (
	"crypto/rand"
	"fmt"
)

// Key returns random string
func Key() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err) // out of randomness, should never happen
	}
	return fmt.Sprintf("%x", buf)
}

func main() {
	fmt.Println(Key())
}
