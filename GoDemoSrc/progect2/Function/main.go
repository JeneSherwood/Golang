package main

import (
	"fmt"
	/*该出放置别名，使用后后面的地址失效*/ /*"go_code/progect1/mode03"*/
	_ "go_code/progect1/Fibonacci"
	"go_code/progect1/monkey"
)

func main() {
	//函数的基本介绍
	//基本语法
	//	func 函数名 （形参列表） （返回值列表）{
	//			执行语句
	//			return 返回值列表
	//	}

	// var num1 float64
	// var num2 float64
	// var operator byte = '+'
	// fmt.Println("请输入两个数字以及一个运算符")
	// fmt.Scanln(&num1)
	// fmt.Scanln(&num2)
	// fmt.Scanln(&operator)
	// result := mode03.Cal(num1, num2, operator)
	// fmt.Println("result=", result)

	//包的基本介绍：go的每一个文件都属于一个包，go是以包的形式管理文件和目录结构的。
	//1、区分相同名字的函数、变量、标识符等
	//2、可以很好的管理大量的文件
	//3、控制函数、变量访问范围，即作用域
	//4、为了使其他包的文件可以访问，必须使改文件名的首字母需要大写
	//5、如果包名过长，可以取一个别名，原来的名字不能使用
	//6、同一个包下，不允许有相同的函数名，也不可以有相同的全局变量名

	//go语言函数的调用机制

	//return语句
	//基本语法
	// 	func 函数名（形参列表）（返回值列表）{
	// 		语句。。。。
	// 		return （返回值列表）
	// 	}
	//1、返回多个值时用“_”占位符可以忽略该值
	//2、如果只返回一个值，可以不写括号

	//递归调用
	//1、执行函数时，就会创建一个独立的空间
	//2、函数的局部变量是独立的，不会相互影响
	//3、递归必须向退出递归的条件逼近，否则就会无限递归
	//4、当函数执行完毕，就会return，遵循谁调用，返回给谁。

	//递归课堂练习1
	//斐波那契数列
	//此处调用progect1文件夹下Fibonacci包下Fibonacci函数
	// result := Fibonacci.Fibonacci(10)
	// fmt.Println("result=", result)

	//递归课堂练习2
	//猴子吃桃问题
	//此处调用progect1文件夹下monkey包下monkey函数
	result := monkey.Monkey(1)
	fmt.Println("reslut=", result)

}
