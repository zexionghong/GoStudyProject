package day01

import (
	"Study/designPatterns/day01/class_path"
	"fmt"
	"testing"
)

// 实例工厂
func CreateFact(name string) class_path.IncFactory {
	var f class_path.IncFactory
	switch name {
	case "ios":
		f = &class_path.IphoneFactory{}
	case "android":
		f = &class_path.AndroidFactory{}

	}

	return f

}

func TestAbs(t *testing.T) {
	fact := CreateFact("ios")
	afact := CreateFact("android")
	fmt.Println(fact.Create(14))
	hw_phone := afact.Create(14)
	fmt.Println(hw_phone.(class_path.HuaWei))
}
