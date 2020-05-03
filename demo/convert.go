package demo

import (
	"fmt"
	"strconv"
)

func main()  {
	var s1 string = "1"
	var s2 string = "1.0"
	var int1 int = 1
	var i64 int64 = 1
	//string到int
	i1, _:=strconv.Atoi(s1)
	fmt.Println("int s1 ", i1)
	i2, _:=strconv.Atoi(s2)
	fmt.Println("int s2 ", i2) //0
	//string到int64
	i164, _ := strconv.ParseInt(s1, 10, 64)
	fmt.Println("int64 s1 ", i164)
	//int到string
	s3 := strconv.Itoa(int1)
	fmt.Println("int to string ", s3)
	//int64到string
	s4 := strconv.FormatInt(i64,10)
	fmt.Println("int64 to string ", s4)
	//字符串到float32/float64
	f2, _ := strconv.ParseFloat(s1, 32)
	fmt.Println("float32 ", f2)
	f3, _ := strconv.ParseFloat(s1,64)
	fmt.Println("float64 ", f3)

	f4, _ := strconv.ParseFloat(s2, 32)
	fmt.Println("float32 1.0 ", f4)
	f5, _ := strconv.ParseFloat(s2,64)
	fmt.Println("float64 1.0 ", f5)
	//int64转int
	fmt.Println("int64 to int ", int(i64))
	//int转int64
	toInt64 := int64(i1)
	fmt.Println("int to int64 ", toInt64)
}