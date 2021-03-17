package slice

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3}
	addr := &s
	fmt.Printf(" addr = %p", addr)
	ptrTest(addr) // 拿到地址是一个指针
}

func ptrTest(arr *[]int) {
	arrLen := len(*arr)
	fmt.Println(arrLen)
	*arr = append(*arr, 4)

}
