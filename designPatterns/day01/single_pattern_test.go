package day01

import (
	"fmt"
	"sync"
	"testing"
)

/**
单例模式，单进程中有且只有一个全局的实例对象，节约开销，创建实例的时候，判断是否存在，存在则返回对象
*/

var boss *Boss
var once sync.Once

type Boss struct {
	name string
}

func CreateBoss(name string) *Boss {

	once.Do(func() {
		boss = &Boss{name: name}
	})
	fmt.Println(boss)
	return boss
	//mu.Lock()
	//defer mu.Unlock()
	//fmt.Println(boss)
	//if boss == nil {
	//	boss = Boss{name: name}
	//}
	//return boss
}

func TestSingle(t *testing.T) {
	b := CreateBoss("111")
	fmt.Println(&b)

	c := CreateBoss("11441")
	c.name = "333"
	fmt.Println(&c)
	//fmt.Println(b.name)
}
