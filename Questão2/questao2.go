package main

import "fmt"

func main(){
	var a [5] int
	fmt.Println("a: ", a)

	var Matriz [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            Matriz[i][j] = i + j
        }
    }
    fmt.Println("2d: ", Matriz)
}

