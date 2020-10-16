package main

import "fmt"

func main(){
    var a[5] int
    fmt.Println("values : ", a)
    a[4] = 100
    fmt.Println("print all values : ", a)
    fmt.Println("a[4] : ", a[4])
    fmt.Println("length of array : ", len(a))
    b := [5] int {1,2,3,4,5}
    fmt.Println("all values : ", b)
}