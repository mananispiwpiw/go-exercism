package airportrobot

// Write your code here.
type Greeter interface {
	LanguageName() string
	Greet(a string) string
}

// SayHello() function
func SayHello(name string, g Greeter) string {
	return "I can speak " + g.LanguageName() + ": " + g.Greet(name)
}

// Italian
type Italian struct{}

// Method Italian
func (i Italian) LanguageName() string {
	return "Italian"
}
func (i Italian) Greet(name string) string {
	//"Ciao {name}!"
	return "Ciao " + name + "!"
}

// Portugese
type Portuguese struct{}

// Method Portugese
func (p Portuguese) LanguageName() string {
	return "Portuguese"
}
func (p Portuguese) Greet(name string) string {
	//" {name}!"
	return "Olá " + name + "!"
}

// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
