package main

import "fmt"

func main(){
    var a[3][3] int
    var b[3][3] int
    var c[3][3] int
    for i:=0;i<3;i++{
        for j:=0;j<3;j++{
            fmt.Scan(&a[i][j])
        }
    }
    for i:=0;i<3;i++{
        for j:=0;j<3;j++{
            fmt.Scan(&b[i][j])
        }
    }
    for i:=0;i<3;i++{
        for j:=0;j<3;j++{
            c[i][j] = a[i][j] + b[i][j]
        }
    }
    for i:=0;i<3;i++{
        for j:=0;j<3;j++{
            fmt.Println("the sum of 2d matrix is : ", c[i][j])
        }
    }
}