package main

import (
	"fmt"
	"math"
	"github.com/shopspring/decimal"
)

//相等
func intEqual(a int, b int) bool {
	return a == b
}

//比较大小
func intCompare(a int, b int) int  {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

//加
func addInt(a int, b int) int {
	return a + b
}

//减
func subInt(a int, b int) int {
	return a - b
}

//乘
func mulInt(a int, b int) int {
	return a * b
}

//除
func dividInt(a int, b int) int {
	return a / b
}

//乘方
func powerFloat(a float64, b float64) float64 {
	return math.Pow(a, b)
}

func main()  {
	var (
		a int = 2
		b int = 1
		c float64 = 2.0
		d float64 = 3.0
	)

	fmt.Println("相等 ", intEqual(a, b))
	fmt.Println("比较 ", intCompare(a, b))
	fmt.Println("加法 ", addInt(a, b))
	fmt.Println("减法 ", subInt(a, b))
	fmt.Println("乘法 ", mulInt(a, b))
	fmt.Println("除法 ", dividInt(a, b))
	fmt.Println("乘方 ", powerFloat(d, c))

	var (
		e decimal.Decimal = decimal.NewFromInt(4)
		f decimal.Decimal = decimal.NewFromInt(2)
		g decimal.Decimal = decimal.NewFromInt(-3)
	)
	fmt.Println("相等 ", e.Equal(f))
	fmt.Println("比较 ", e.Cmp(f))
	fmt.Println("加法 ", e.Add(f))
	fmt.Println("减法 ", e.Sub(f))
	fmt.Println("乘法 ", e.Mul(f))
	fmt.Println("除法 ", e.Div(f))
	fmt.Println("乘方 ", e.Pow(f))
	fmt.Println("abs ", g.Abs())
	fmt.Println("neg ", g.Neg())
	fmt.Println("mod ", g.Mod(f))
	fmt.Println("bigFloat ", e.BigFloat())
}