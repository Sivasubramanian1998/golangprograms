package main
import "fmt"
func main(){
    var a,b int = 10,20
    a = a+b
    b = a-b
    a = b-a
    fmt.Println(a,b)
}