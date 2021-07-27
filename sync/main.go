package main

import "sync"

type MyCounter struct {
	mu    sync.Mutex
	count int
}

func (m *MyCounter) add() {

	m.mu.Lock()
	defer m.mu.Unlock()

	m.count++

}

func (m *MyCounter) value() int {
	return m.count
}

func main() {

}
