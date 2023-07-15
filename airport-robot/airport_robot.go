package airportrobot

import (
	"fmt"
)

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(a string) string
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %v: %v", greeter.LanguageName(), greeter.Greet(name))
}

type Italian struct {
	language string
}

func (i Italian) LanguageName() string {
	i.language = "Italian"
	return i.language
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %v!", name)
}

type Portuguese struct {
	language string
}

func (p Portuguese) LanguageName() string {
	p.language = "Portuguese"
	return p.language
}

func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %v!", name)
}
