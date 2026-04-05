package main

import "fmt"

// ---------- 1. 定义接口 ----------
// 规定：凡是“会唱歌”的，都必须实现 Sing()
type Singer interface {
	Sing()
}

// ---------- 2. 定义结构体 ----------
type Chicken struct {
	Name string
}

type Dog struct {
	Name string
}

// ---------- 3. 实现接口方法 ----------
// Chicken 实现了 Singer 接口
func (c Chicken) Sing() {
	fmt.Printf("%s 鸡：咯咯咯\n", c.Name)
}

// Dog 也实现了 Singer 接口
func (d Dog) Sing() {
	fmt.Printf("%s 狗：汪汪汪\n", d.Name)
}

// ---------- 4. 业务逻辑（核心） ----------
// 这里接收的是 Singer 接口，而不是具体的 Chicken 或 Dog
func StartShow(singer Singer) {
	// 不管传入的是鸡还是狗，只要会唱就行
	singer.Sing()
}

func main() {
	// 创建实例
	ck := Chicken{Name: "小黄"}
	dg := Dog{Name: "旺财"}

	// 直接调用（原本写法）
	ck.Sing()
	dg.Sing()

	fmt.Println("----- 换个方式调用 -----")

	// 5. 通过接口调用（多态）
	StartShow(ck) // 传入 Chicken，自动满足接口 Singer
	StartShow(dg) // 传入 Dog，自动满足接口 Singer
}
