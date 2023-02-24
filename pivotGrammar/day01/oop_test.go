package day01

import (
	"fmt"
	"testing"
)

type IAnimal interface {
	eat()
	play()
}
type Animal struct {
	name string
}

func (receiver *Animal) eat() {
	fmt.Println(fmt.Sprintf("%s吃饭了", receiver.name))
}

func (receiver *Animal) play() {

}

func newAnimal(name string) *Animal {
	return &Animal{name: name}
}

type Cat struct {
	*Animal
}

func newCat(name string) *Cat {
	return &Cat{Animal: newAnimal(name)}
}

func TestAbs(t *testing.T) {
	newCat("ss").eat()
}
