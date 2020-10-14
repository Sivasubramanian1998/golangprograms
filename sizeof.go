package main
import "fmt"
import "unsafe"
func main(){
    var a int
    var b float64
    var c byte
    fmt.Println(unsafe.Sizeof(a))
    fmt.Println(unsafe.Sizeof(b))
    fmt.Println(unsafe.Sizeof(c))
}