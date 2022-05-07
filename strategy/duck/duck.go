package duck

import "fmt"

type FlyBehavior interface {
	fly()
}

type QuackBehavior interface {
	quack()
}

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func New(fb FlyBehavior, qb QuackBehavior) *Duck {
	return &Duck{
		flyBehavior:   fb,
		quackBehavior: qb,
	}
}

func (d *Duck) Display() {
	fmt.Println("I'm a duck")
}

func (d *Duck) Swim() {
	fmt.Println("I'm swimming")
}

func (d *Duck) Quack() {
	d.quackBehavior.quack()
}

func (d *Duck) Fly() {
	d.flyBehavior.fly()
}
