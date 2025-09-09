package airportrobot

import "fmt"

type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

func SayHello(name string, greet Greeter) string {
    return fmt.Sprintf("I can speak %s: %s", greet.LanguageName(), greet.Greet(name))
}

type Italian struct {}

func (lang Italian) LanguageName() string {
    return "Italian"
}
func (lang Italian) Greet(name string) string {
    return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct {}

func (lang Portuguese) LanguageName() string {
    return "Portuguese"
}
func (lang Portuguese) Greet(name string) string {
    return fmt.Sprintf("Olá %s!", name)
}