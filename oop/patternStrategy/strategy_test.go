package patternStrategy

import "testing"

func TestStrategy(t *testing.T) {
	var duck Duck
	duck = NewMallardDuck()

	// fly
	duck.Fly()

	// quack
	duck.Quack()

	// give super powers to mallard duck
	var flyBehavior FlyBehavior
	flyBehavior = SuperFlyBehavior{}

	var quackBehavior QuackBehavior
	quackBehavior = SuperQuackBehavior{}

	duck.SetFlyBehavior(flyBehavior)
	duck.SetQuackBehavior(quackBehavior)

	// fly
	duck.Fly()

	// quack
	duck.Quack()
}
