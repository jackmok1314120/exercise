package main

import "fmt"

type RequestBody struct { //要求data
	Asset         string `json:"asset"`   //传账币种
	OutputAddress string `json:"address"` //提款目标地址
}
type RequestBod []struct { //要求data
	Asset         string `json:"asset"`   //传账币种
	OutputAddress string `json:"address"` //提款目标地址
}
type student struct {
	name  string
	child *student
}

// 打印消息类型, 传入匿名结构体
func printMsgType(msg *struct {
	id   int
	data string
}) {
	// 使用动词%T打印msg的类型
	fmt.Printf("%T\n", msg)
}

type Cat struct {
	Color string
	Name  string
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}
func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

type BlackCat struct { //定义 BlackCat 结构，并嵌入了 Cat 结构体，BlackCat 拥有 Cat 的所有成员，实例化后可以自由访问 Cat 的所有成员。
	Cat // 嵌入Cat, 类似于派生
}

// “构造基类”
func NewCat(name string) *Cat { //NewCat() 函数定义了 Cat 的构造过程，使用名字作为参数，填充 Cat 结构体
	return &Cat{
		Name: name,
	}
}

// “构造子类”
func NewBlackCat(color string) *BlackCat { //NewBlackCat() 使用 color 作为参数，构造返回 BlackCat 指针
	cat := &BlackCat{} //实例化 BlackCat 结构，此时 Cat 也同时被实例化。
	cat.Color = color  //填充 BlackCat 中嵌入的 Cat 颜色属性，BlackCat 没有任何成员，所有的成员都来自于 Cat。
	return cat
}
func main() {
	WallletAsset := "BTC"
	WalletAddress := "31h1yc8MnKj49By5tSjApsaLUnEeJ1g4u4"
	kyt := RequestBody{Asset: WallletAsset, OutputAddress: WalletAddress}
	relation := &student{
		name: "tom",
		child: &student{
			name: "jack",
			child: &student{
				name: "sophia",
			},
		},
	}
	// 实例化一个匿名结构体
	msg := &struct { // 定义部分
		id   int
		data string
	}{ // 值初始化部分
		1024,
		"hello",
	}
	tomCat := NewCatByName("tom")
	pinCat := NewCatByColor("pink")
	printMsgType(msg)
	fmt.Println(tomCat, pinCat)
	fmt.Println(Cat{Color: "black"})
	fmt.Println(kyt)
	fmt.Println(relation)
}
