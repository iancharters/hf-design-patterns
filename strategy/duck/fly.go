package duck

import "fmt"

type NoFlyBehavior struct{}

func (f NoFlyBehavior) fly() {
	fmt.Println("I can't fly")
}

type FlyWithWings struct{}

func (f FlyWithWings) fly() {
	fmt.Println("I'm flying")
}

type FlyWithRocket struct{}

func (f FlyWithRocket) fly() {
	fmt.Println("I'm flying with a rocket")
}
