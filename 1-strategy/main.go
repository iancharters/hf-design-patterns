package main

import (
	"github.com/iancharters/hf-design-patterns/1-strategy/duck"
)

func main() {
	d := duck.New()

	d.display()
}
