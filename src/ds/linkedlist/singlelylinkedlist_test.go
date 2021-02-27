package linkedlist

import (
    "fmt"
    "testing"
)

var temp = 0

func TestListNew(t *testing.T) {
    fmt.Println(temp)
    incre(temp)
    fmt.Println(temp)
}
func incre(toIncre int) int {

    i := toIncre + 1
    fmt.Printf("i %d", i)
    fmt.Println()
    return i
}
