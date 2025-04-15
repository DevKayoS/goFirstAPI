package utils

import "math/rand/v2"

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const lengthChar = 8

func GenCode() string {
	byts := make([]byte, lengthChar)

	for i := range lengthChar {
		byts[i] = characters[byte(rand.IntN(len(characters)))]
	}

	return string(byts)
}
