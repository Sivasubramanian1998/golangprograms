package main

import "fmt"

func main(){
    var a[2][3] int 
    for i:=0;i<2;i++{
        for j:=0;j<3;j++{
            a[i][j] = i+j
        }
    }
    fmt.Println("values are : ", a)
}