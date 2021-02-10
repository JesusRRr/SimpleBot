package main

import "fmt"

type bot interface {
	getGreeting() string
}

type spanishBot struct {
}

type englishBot struct {
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (englishBot) getGreeting() string {
	return "Hi there!"
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
