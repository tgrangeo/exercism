package airportrobot
import "fmt"
// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
    LanguageName() string
    Greet(a string) string
}

type Italian struct{
}

type Portuguese struct{
}

func (i Italian)LanguageName()string{
    return "Italian"
}
func (i Italian)Greet(a string)string{
    return "Ciao " + a + "!"
}

func (i Portuguese)LanguageName()string{
    return "Portuguese"
}
func (i Portuguese)Greet(a string)string{
    return "Olá " + a+ "!"
}

func SayHello(str string, g Greeter) string{
    return fmt.Sprintf("I can speak %s: %s", g.LanguageName(),g.Greet(str))
}