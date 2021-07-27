package main

type MyCounter struct {
	count int
}

func (m *MyCounter) add() {

	m.count++

}

func (m *MyCounter) value() int {
	return m.count
}

func main() {

}
