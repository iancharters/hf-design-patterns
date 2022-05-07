package duck

type FlyBehavior interface {
	fly() string
}

type QuackBehavior interface {
	quack() string
}

// type Duck interface {
// 	display() string
// 	swim() string
// 	quack() string
// 	fly() string
// }

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func New() *Duck {
	return &Duck{}
}

func (d *Duck) Display() string {
	return "I'm a duck"
}

func (d *Duck) Swim() string {
	return "I'm swimming"
}

func (d *Duck) Quack() string {
	return d.quackBehavior.quack()
}

func (d *Duck) Fly() string {
	return d.flyBehavior.fly()
}
