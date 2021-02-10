package main

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

func main() {
}
