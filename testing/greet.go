package greet

import (
	"errors"
	"fmt"
	"math/rand"
)

var greetings = []string{
	"Hi, %s!",
	"Hello, %s!",
	"Hola, %s!",
	"Greetings, %s!",
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf(greetings[rand.Intn(len(greetings))], name), nil
}

func Greet(names []string) (map[string]string, error) {
	if len(names) == 0 {
		return nil, errors.New("no names provided")
	}
	result := make(map[string]string, len(names))
	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, fmt.Errorf("greet %q: %w", name, err)
		}
		result[name] = msg
	}
	return result, nil
}
