package slice

import (
    "fmt"
    "testing"
)

func TestSlice(t *testing.T) {
    var s Slice
    // 创建切片
    s.Create(5, 5, 1, 2, 3, 4, 5)
    s.Append(6, 7, 8)
    s.Append(9, 10, 11)
    // 打印切片
    s.PrintSlice()

    fmt.Println("\n长度：", s.len)
    fmt.Println("容量：", s.cap)

    // 根据下标获取元素
    value := s.SearchData(11)
    fmt.Println("值：", value)

    // 根据元素获取下标
    index := s.SearchData(6)
    fmt.Println("下标：", index)

    // 删除元素
    s.DeleteDate(5)
    s.DeleteDate(5)
    s.PrintSlice()
    fmt.Println()

    // 指针偏移
    s.Insert(8, 666)
    s.PrintSlice()

    // slice销毁
    fmt.Println(s)
    s.Destroy()
    fmt.Println(s)

}
