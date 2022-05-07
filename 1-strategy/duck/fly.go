package duck

type NoFlyBehavior struct {
}

func (f *NoFlyBehavior) fly() string {
	return "I can't fly"
}

type FlyWithWings struct{}

func (f *FlyWithWings) fly() string {
	return "I'm flying"
}

type FlyWithRocket struct{}

func (f *FlyWithRocket) fly() string {
	return "I'm flying with a rocket"
}
