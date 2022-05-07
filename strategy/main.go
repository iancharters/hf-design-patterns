package main

import (
	"fmt"

	"github.com/iancharters/hf-design-patterns/strategy/duck"
	"github.com/iancharters/hf-design-patterns/strategy/duck/mallard"
)

func main() {
	fmt.Println("Default Impl")
	d := duck.New(
		duck.NoFlyBehavior{},
		duck.Squeak{},
	)

	d.Display()
	d.Swim()
	d.Fly()
	d.Quack()

	fmt.Println("\nMallard Impl")

	m := mallard.New(
		duck.FlyWithWings{},
		duck.QuackSound{},
	)

	m.Display()
	m.Swim()
	m.Fly()
	m.Quack()
}
