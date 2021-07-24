package main

import (
	"bufio"
	"fmt"
	"os"
)

const Englishprefix = "Hello"
const Frenchprefix = "Hola"
const Hindiprefix = "Namesta"

type Employee interface {
	Language() string
	Age() int
	Random() (string, error)
}

type Engineer struct {
	Name string
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}
func (e Engineer) Age() int {
	return 25
}
func (e Engineer) Random() (string, error) {
	return e.Name + "Random", nil
}

//A hello function for starters
func hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + " " + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = Frenchprefix

	case "hindi":
		prefix = Hindiprefix

	default:
		prefix = Englishprefix

	}
	return
}

func sumAll(value ...[]int) (result []int) {
	for _, i := range value {
		sum := 0
		for _, j := range i {

			sum += j
		}
		result = append(result, sum)
	}

	return

}

func main() {
	// This will throw an error
	var programmers []Employee
	elliot := Engineer{Name: "Elliot"}
	// Engineer does not implement the Employee interface
	// you'll need to implement Age() and Random()
	programmers = append(programmers, elliot)
	fmt.Println(programmers)
	p, _ := os.ReadFile("test.data")
	fmt.Println(string(p))

	f, err := os.OpenFile("test.data", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	f.WriteString("asdsdsdsffffff")
	p, _ = os.ReadFile("test.data")
	fmt.Println(string(p))

	input := bufio.NewReader(os.Stdin)

	h, _ := input.ReadString('\n')

	fmt.Println(h)

}
