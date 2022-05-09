package displayelement

import "github.com/iancharters/hf-design-patterns/observer/observer"

type DisplayElement interface {
	observer.Observer
	display()
}
