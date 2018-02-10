package main

import "fmt"

func main() {
    
    fmt.Printf("test for loop ...")    
    for i := 0; i < 10; i++ {
        println(i)
    }

    fmt.Printf("test for goto ...")
    j := 0
Here:
    println(j)
    j++
    if j < 10 {
        goto Here
    }
    
    fmt.Printf("test for array loop ...")
    var a = [...]int{'a', 'b', 'c', 'd', 'e'}
    for k,v := range a {
        fmt.Printf("key is %d value is %c", k, v)
    }
}
