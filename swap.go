package main
import "fmt"
func main(){
    var a,b int = 10,20
    var c = a
    a = b
    b = c
    fmt.Println(a,b)
}