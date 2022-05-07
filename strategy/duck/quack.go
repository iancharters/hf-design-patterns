package duck

import "fmt"

type QuackSound struct{}

func (q QuackSound) quack() {
	fmt.Println("QUACK!")
}

type Squeak struct{}

func (q Squeak) quack() {
	fmt.Println("SQUEAK!")
}

type Silence struct{}

func (q Silence) quack() {
	fmt.Println("...")
}
