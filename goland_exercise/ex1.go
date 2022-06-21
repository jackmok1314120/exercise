package main

import (
	"fmt"
	"runtime"
	//"math"
)

//const Pi = 3.14 //常量不能用：=
/*func main() {
fmt.Println("My favorite number is", rand.Intn(100))*/

//函数可以没有参数或接受多个参数。
/*func add(x int, y int) int {
	return x + y
}*/
/*func swap(x, y  string) (string, string) {
	return y, x
}*/
//命令返回值
/*func split(sum float32) (x, y float32) {
	x = sum * 4 / 9
	y = sum - x
	return
} */
//var c, python, java bool//变量1
//var i, j int = 1, 2//变量初始化2
/*var (    //基本类型
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
) */
/*const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
) */
/*func needInt(x int) int { return x*10 + 1 } //数值常量 一个未指定类型的常量由上下文来决定其类型。
func needFloat(x float64) float64 {
	return x * 0.1
} */
//Go 的 if 语句与 for 循环类似
/*func sqrt(x float64) string {
	if x < 0 {   //x<0 进内循环
		return sqrt(-x) + "i"
	} 			//反之外循环直接回传结果
	return fmt.Sprint(math.Sqrt(x))
}*/

/*func pow(x, n, lim float64) float64 {//if 的简短语句
	if v := math.Pow(x, n); v < lim {	//同 for 一样， if 语句可以在条件表达式前执行一个简单的语句
		return v
	}							//该语句声明的变量作用域仅在 if 之内
	return lim
}*/

//1216练习：循环与函数
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 15; i++ {
		z -= (z*z - x) / (2 * z)               //牛顿法
		fmt.Println("牛顿法计算根号第", i+1, "次=：", z) //回传每次z值
	}
	return z
}
func main() {
	/*fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))*/
	//fmt.Println(math.Pi)
	//fmt.Println(add(42, 13))
	//a, b := swap("hello", "world") //多值返回
	//fmt.Println(a, b)
	//fmt.Println(split(17))//长代码影响可读性
	/*var i int// 变量1
	fmt.Println(i, c, python, java)*/
	/*var d,c, python, java = true, false, "no!",true//变量初始化2
	fmt.Println(d,i, j, c, python, java)*/
	/*var i, j int = 1, 2//短变量声明，:= 可在类型明确的地方代替 var 声明； := 结构不能在函数外使用
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java) */

	/*Go 的基本类型有
	bool
	string
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // uint8 的别名
	rune // int32 的别名
	    // 表示一个 Unicode 码点
	float32 float64
	complex64 complex128*/
	/*fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z) */
	/*var i int   //变量声明 初始值 为零值
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s) */

	/*var x, y int = 3, 4    //表达式 T(v) 将值 v 转换为类型 T。
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)*/
	/* v := 3 + 4.999 // 修改这里！类型推导
	fmt.Printf("v is of type %T\n", v ) */
	/*const World = "世界"   //常量不能用：=
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)*/

	/*fmt.Println(needInt(Small))  //数值常量
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big) */
	/*sum := 0	//for回圈
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)*/

	/*sum1 := 1		//初始化语句和后置语句是可选的
	for ; sum1 < 1000; {
		sum1 += sum1
	}
	fmt.Println(sum1)

	sum2 := 1		//去掉分号，因为 C 的 while 在 Go 中叫做 for
	for sum2 < 101 {
		sum2 += sum2
	}*/
	/*fmt.Println(sum2)
	//for{}   无限循环
	fmt.Println(sqrt(2), sqrt(-4))*/
	//1216ex
	/*fmt.Println(		//if 简短句
		pow(3, 2, 10),
		pow(3, 3, 20),
	)*/
	Sqrt(2) //练习：循环与函数
	//fmt.Println("牛顿法求平方根=", Sqrt(2) )//牛顿法求跟号
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
