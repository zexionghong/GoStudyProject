package class_path

import "fmt"

type Android struct {
}

func (receiver *Android) Call() {

}

type HuaWei struct {
	*Android
	Name string
}

func (w HuaWei) Patriotism() {
	fmt.Println("老子爱国")

}
