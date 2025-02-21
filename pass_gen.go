package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func generatePassword(length int) (string, error) {
	// Define the characters that can appear in the password
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+[]{}|;:,.<>?/`~"

	// Create a slice to store the password
	var password strings.Builder

	for i := 0; i < length; i++ {
		// Generate a random index for the chars string
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}

		// Append the character at the random index to the password
		password.WriteByte(chars[index.Int64()])
	}

	// Return the generated password
	return password.String(), nil
}

func main() {
	// Set desired password length
	passwordLength := 12

	// Generate a password
	password, err := generatePassword(passwordLength)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}

	// Print the generated password
	fmt.Println("Generated Password:", password)
}
