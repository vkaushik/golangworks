package patternStrategy

import "fmt"

type Duck interface {
	Fly()
	Quack()
	SetFlyBehavior(flyBehavior FlyBehavior)
	SetQuackBehavior(quackBehavior QuackBehavior)
}

func (this *MallardDuck) SetFlyBehavior(flyBehavior FlyBehavior) {
	this.flyBehavior = flyBehavior
}
func (this *MallardDuck) SetQuackBehavior(quackBehavior QuackBehavior) {
	this.quackBehavior = quackBehavior
}

type FlyBehavior interface {
	performFly()
}

type QuackBehavior interface {
	performQuack()
}

type MallardDuck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{
		flyBehavior:   NormalFlyBehavior{},
		quackBehavior: NormalQuackBehavior{},
	}
}

type NormalFlyBehavior struct {
}

type NormalQuackBehavior struct {
}

type SuperFlyBehavior struct {
}

type SuperQuackBehavior struct {
}

func (this NormalQuackBehavior) performQuack() {
	fmt.Println("Normal quack...")
}

func (this NormalFlyBehavior) performFly() {
	fmt.Println("Normal fly...")
}

func (this SuperFlyBehavior) performFly() {
	fmt.Println("Super fly...")
}

func (this SuperQuackBehavior) performQuack() {
	fmt.Println("Super quack...")
}

func (this MallardDuck) Fly() {
	this.flyBehavior.performFly()
}

func (this MallardDuck) Quack() {
	this.quackBehavior.performQuack()
}
