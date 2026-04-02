package main

import "fmt"

func Hello() {
	fmt.Println("123456")
}

func NameAndAge() {
	fmt.Print("请输入你的名字：")
	var name string
	var age int
	fmt.Scan(&name)
	fmt.Println(name)
	fmt.Print("请输入你的年龄：")
	fmt.Scan(&age)
	fmt.Println(age)

}
func main() {
	var name string = "Hello"
	var name2 string = "World"
	name3 := "666"
	fmt.Println("Hello, World!")
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
	Hello()
	fmt.Printf("%s,World", "Hello")
	fmt.Printf("\n")
	// NameandAge()
	var a byte = 'a'
	fmt.Printf("%c%d\n", a, a)
	var z rune = '牛'
	fmt.Printf("%c%d\n", z, z)
	fmt.Println("牛牛牛")
	fmt.Println("\"牛牛\"牛牛")
	fmt.Println(`"FU牛\r牛FU"`)
	var name1 []string = []string{"张三", "李四"}
	var name4 = []string{"张三", "李四"}
	fmt.Println(name1)
	fmt.Println(name4)
	fmt.Println(name1[1])
	var age1 [3]string
	fmt.Println(age1)
	var age2 [3]int
	fmt.Println(age2)
	fmt.Println(age2[0:2])
	fmt.Println(age2[2])
}
