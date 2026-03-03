package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
)

const (
	LETTERS       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DIGITS        = "0123456789"
	SPECIAL_CHARS = "!@#$%^&*()-_=+,.?"
)

func main() {
	// Define CLI flags
	length := flag.Int("l", 16, "Length of the password")
	useDigits := flag.Bool("d", false, "Include numbers")
	useSpecial := flag.Bool("s", false, "Include special characters")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of passgen:\n\n")
		fmt.Fprintf(os.Stderr, "passgen is a cryptographically secure password generator.\n")
		fmt.Fprintf(os.Stderr, "Example: ./passgen -l 32 -s=false\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	charset := LETTERS
	if *useDigits {
		charset += DIGITS
	}
	if *useSpecial {
		charset += SPECIAL_CHARS
	}

	password, err := generatePassword(*length, charset)
	if err != nil {
		fmt.Printf("Error generating password: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated Password: %s\n", password)
}

func generatePassword(length int, charset string) (string, error) {
	password := make([]byte, length)
	for i := range password {
		// crypto/rand provides security, unlike math/rand
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[idx.Int64()]
	}
	return string(password), nil
}
