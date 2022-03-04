package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

// constants
const invalid = "nome inválido"
const empty = " (não pode ser vazio)"
const notAlphabet = " (deve conter apenas letras)"

// Hello returns a greetings or an error
func Hello(name string) (string, error) {
	nameTrimmed := strings.TrimSpace(name)

	if nameTrimmed == "" {
		return "", errors.New(invalid + empty)
	}

	if !alphabetical(nameTrimmed) {
		return "", errors.New(invalid + notAlphabet)
	}

	message := fmt.Sprintf(randomMessage(), nameTrimmed)
	return message, nil
}

// Hellos returns a list of greetings or an error
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// init is called before the main function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// alphabetical returns true if the string contains only letters
func alphabetical(str string) bool {
	for _, chr := range str {
		if !unicode.IsLetter(chr) {
			return false
		}
	}
	return true
}

// randomMessage returns a random message
func randomMessage() string {
	messages := []string{
		"Olá, %v. Seja Bem-vindo(a)!",
		"Que bom te ver por aqui, %v!",
		"Salve, %v! #TamoJunto",
	}
	return messages[rand.Intn(len(messages))]
}

/*
removed after identifying implementations in the go lang with the same purpose

func alphabetical(name string) bool {
	for _, char := range strings.Replace(strings.ToLower(name), " ", "", -1) {
		if char < 'a' || char > 'z' {
			return false
		}
	}
	return true
}
*/
