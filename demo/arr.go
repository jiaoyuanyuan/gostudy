package demo

import "fmt"

func main()  {
	var a[3] int
	a[0] = 11
	fmt.Println("a ", a)
	fmt.Println("a 0 ", a[0])

	b := [3]int{11,22,33}
	fmt.Println("b ", b)
	fmt.Println("b 0 ", b[0])
	for k, v := range b {
		fmt.Println("\nk ", k, " v ", v)
	}

	c := [...]int{1,2,3}
	fmt.Println("c ", c)
	fmt.Println("c 0 ", c[0])

	d := [3][2]int{{11,22},{33,44},{55,66}}
	fmt.Println("d ", d)
	fmt.Println("d 0 ", d[0])

	e := [][]int{{1,2},{3,4},{5,6}}
	fmt.Println("e ", e)
	fmt.Println("e 0 ", e[0])
	for key, val := range e {
		fmt.Println("\nkey ", key, " val ", val)
	}
}