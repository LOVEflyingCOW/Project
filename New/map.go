package main

import "fmt"

type Student struct {
	Name string
}
type Class struct {
	Student
	Id string
}

func (s Student) Sdudy() {
	fmt.Printf("%s正在学习", s.Name)
}

func main() {
	// var name = map[int]string{
	// 	1: "张三",
	// 	2: "李四",
	// }
	// fmt.Println(name)
	// value := name[1]
	// println(value)

	// var age int
	// fmt.Printf("请输入你的年龄:")
	// fmt.Scan(&age)
	// if age < 18 {
	// 	fmt.Println("未成年")
	// 	return
	// }

	S1 := Student{Name: "牛哥"}
	S1.Sdudy()

}
