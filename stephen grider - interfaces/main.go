package main

import "fmt"

type englishBot struct {}
type spanishBot struct {}
type bot interface {
	getGreeting() string
}

func (eB englishBot) getGreeting() string {
	return "Hello!"
}

func (sB spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}