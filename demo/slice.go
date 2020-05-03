package main

import "fmt"

func main()  {
	var a []int
	a = append(a, 1)
	fmt.Println("a ", a)
	fmt.Println("a 0 ", a[0])

	b := []int{1, 2, 3}
	fmt.Println("b ", b)
	fmt.Println("b 0 ", b[0])

	c := make([]int, 5, 10)
	fmt.Println("c ", c)
	c = append(c, 1)
	fmt.Println("c append ", c)

	for i := 0; i < len(c); i++ {
		fmt.Println(i, c[i])
	}

	for index, value := range c {
		fmt.Println(index, value)
	}

	//copy
	copy(c, a)
	fmt.Println("c copy ", c)

	//delete
	c = append(c, 3)
	fmt.Println("c before delete ", c)
	c = append(c[:1],c[2:]...)
	fmt.Println("c delete ", c)
}