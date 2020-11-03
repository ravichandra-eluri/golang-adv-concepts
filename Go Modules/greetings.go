package greetings

import "fmt"

//var message string

func Hello(name string) string  {
	message := fmt.Sprintf("Hi, %v. Welcome!",name)
	return message
	
}