package main

import "fmt"

// 基础接口：定义鸟类行为
type Bird interface {
	Fly() string // 飞行行为
}

// 麻雀实现：会飞
type Sparrow struct{}

func (s Sparrow) Fly() string {
	return "麻雀在飞翔！"
}

// 企鹅实现：不会飞，但通过接口适配
type Penguin struct{}

func (p Penguin) Fly() string {
	return "企鹅不会飞，但能在水里游泳！"
}

// 通用函数：接受任何 Bird 类型
func LetBirdFly(b Bird) {
	fmt.Println(b.Fly())
}

func main() {
	sparrow := Sparrow{}
	penguin := Penguin{}

	// 子类替换父类（LSP 验证）
	LetBirdFly(sparrow) // 输出：麻雀在飞翔！
	LetBirdFly(penguin) // 输出：企鹅不会飞，但能在水里游泳！
}
