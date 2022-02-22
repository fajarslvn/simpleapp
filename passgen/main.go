package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZJK"
	character = "!@#$%^&*"
	number    = "0123456789"
	allChar   = lowercase + uppercase + character + number
)

func main() {
	rand.Seed(time.Now().Unix())
	passGen := generatePassword()
	fmt.Println(passGen)
}

func generatePassword() string {
	var (
		pass       strings.Builder
		passLength = 30
		minChar    = 2
		minNum     = 3
		minUpp     = 2
	)

	for i := 0; i < minChar; i++ {
		random := rand.Intn(len(character))
		pass.WriteString(string(character[random]))
	}

	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(number))
		pass.WriteString(string(number[random]))
	}

	for i := 0; i < minUpp; i++ {
		random := rand.Intn(len(uppercase))
		pass.WriteString(string(uppercase[random]))
	}

	// Generate Password
	lenChar := passLength - minChar - minNum - minUpp
	for i := 0; i < lenChar; i++ {
		random := rand.Intn(len(allChar))
		pass.WriteString(string(allChar[random]))
	}
	inRune := []rune(pass.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
