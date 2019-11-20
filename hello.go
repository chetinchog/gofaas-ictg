package main

import "fmt"

// Hello return message
func Hello(user string) string {
	if len(user) == 0 {
		return "Hello Dude!"
	}
	return fmt.Sprintf("Hello %v!", user)
}
