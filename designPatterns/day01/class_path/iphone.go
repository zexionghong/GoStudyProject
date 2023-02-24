package class_path

import "fmt"

type Iphone struct {
}

func (receiver *Iphone) Call() {

}

type iphone14 struct {
	size int
	*Iphone
}

func (i *iphone14) Play(game string) {
	fmt.Println("玩游戏：", game)
}

type iphone12 struct {
	size int
	*Iphone
}

func (i *iphone12) View(name string) {
	fmt.Println("看视频：", name)
}
