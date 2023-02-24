package day01

/**
工厂模式，系统的产品有多于一个的产品族，而系统只消费其中某一族的产品。
*/
import (
	"fmt"
	"testing"
)

type Phone interface {
	Call(number int)
}

type Iphone struct {
}

func (i *Iphone) Call(number int) {
	fmt.Println("iphone打电话")
}

type Android struct {
}

func (a *Android) Call(number int) {
	fmt.Println("android打电话")
}

func factory(phoneType string) Phone {
	var phone Phone
	switch phoneType {
	case "iphone":
		phone = &Iphone{}
	case "android":
		phone = &Android{}
	default:
		phone = &Android{}

	}
	return phone

}

func Test(t *testing.T) {
	phone := factory("iphone")
	phone.Call(1)
}
