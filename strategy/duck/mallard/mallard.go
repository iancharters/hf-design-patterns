package mallard

import (
	"fmt"
	"github.com/iancharters/hf-design-patterns/strategy/duck"
)

type Mallard struct {
	*duck.Duck
}

func New(fb duck.FlyBehavior, qb duck.QuackBehavior) *Mallard {
	return &Mallard{
		Duck: duck.New(fb, qb),
	}
}

func (m *Mallard) Display() {
	fmt.Println("I'm a mallard!")
}
