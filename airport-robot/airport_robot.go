package airportrobot

import "fmt"

type Greeter interface {
	Greet(visitorName string) string
	LanguageName() string
}

type Italian struct{}
type Portuguese struct{}

func (i Italian) Greet(visitorName string) string {
	return fmt.Sprintf("Ciao %s!", visitorName)
}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Portuguese) Greet(visitorName string) string {
	return fmt.Sprintf("Olá %s!", visitorName)
}

func (i Portuguese) LanguageName() string {
	return "Portuguese"
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}
