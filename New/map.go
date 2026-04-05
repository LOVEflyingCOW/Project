package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
}
type Class struct {
	Student
	Id string
}
type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age",omitempty`
	Password string `json:"-"`
}

func (s Student) Study() {
	fmt.Printf("%s正在学习", s.Name)
}
func (c Class) Info() {
	fmt.Printf("%s正在%s班学习", c.Name, c.Id)
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
	S2 := Class{Student: S1, Id: "1"}
	S1.Study()
	fmt.Println()
	S2.Study()
	fmt.Println()
	S2.Info()
	user := User{Name: "牛牛", Age: 18, Password: "666666"}
	Data, _ := json.Marshal(user)
	fmt.Println(string(Data))
}
