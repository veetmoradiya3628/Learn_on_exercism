package airportrobot

import "fmt"

// Greeter interface definition
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

// SayHello returns a greeting string formatted with the language name and the specific greeting.
func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}

// Italian struct
type Italian struct{}

// LanguageName returns the name of the language.
func (i Italian) LanguageName() string {
	return "Italian"
}

// Greet returns the greeting message in Italian.
func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

// Portuguese struct
type Portuguese struct{}

// LanguageName returns the name of the language.
func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

// Greet returns the greeting message in Portuguese.
func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}   