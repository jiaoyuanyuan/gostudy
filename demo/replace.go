package demo
import "fmt"
func main(){
    s := "qwertyuiop"
    r := []rune(s)
    copy(r[4:4+3], []rune("abc"))
    fmt.Printf("Before: %s\n", s)    
    fmt.Printf("After: %s\n", string(r))    
}
