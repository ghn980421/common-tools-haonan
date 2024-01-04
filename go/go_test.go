package _go

import (
	"fmt"
	"testing"
)

type Object struct {
	number int
}

type Toy struct {
	Object
}

func Test_Toy(t *testing.T) {
	toy1 := new(Toy)
	toy2 := &Toy{}
	print(toy1)
	print(toy2)
}

func Test_Defer(t *testing.T) {
	func(a int) {
		defer fmt.Printf("\n1. defer a=%d\n", a) // 一样， 将入参作为值传递给对应方法
		defer func(a int) {                      // 编译之后已经将a值传递到defer的函数结构内了
			fmt.Printf("\n2. defer a=%d\n", a)
			print(&a)
		}(a)
		defer func() {
			fmt.Printf("\n3. defer a=%d\n", a)
			print(&a)
		}()
		a++
		print(&a)
	}(1)
}
