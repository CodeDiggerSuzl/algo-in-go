package slice

//import (
//    "C"
//    "fmt"
//    "unsafe"
//)
//
//type Slice struct {
//    Data unsafe.Pointer // 万能指针类型 对应C语言中的void* TODO
//    len  int
//    cap  int
//}
//
//const TAG = 8
//
//func (s *Slice) Create(l int, c int, Data ...int) {
//    if len(Data) == 0 {
//        return
//    }
//    // 长度小于0 容量小于0 长度大于容量 数据大于长度
//    if l < 0 || c > 0 || l > c || len(Data) > l {
//        return
//    }
//    // ulonglong unsigned long long  无符号的长长整型
//    // 通过C语言代码开辟空间 存储数据
//    // 如果堆空间开辟失败 返回值为NULL 相当于nil 内存地址编号为0的空间
//    s.Data = C.mallco(C.ulonglong(c) * 8)
//    s.len = l
//    s.cap = c
//    // 转化为可以计算的指针类型
//    p := uintptr(s.Data)
//    for _, v := range Data {
//        // 数据储存
//        *(*int)(unsafe.Pointer(p)) = v
//        // 指针偏移
//        p += TAG
//    }
//}
//
//// Print the slice
//func (s *Slice) PrintSlice() {
//    if s == nil {
//        return
//    }
//    // 将万能指针转成可以计算的指针
//    p := uintptr(s.Data)
//    for i := 0; i < s.len; i++ {
//        fmt.Print(*(*int)(unsafe.Pointer(p)), "")
//        p += TAG
//    }
//}
//
//// Append the slice
//func (s *Slice) Append(Data ...int) {
//    if s == nil {
//        return
//    }
//    if len(Data) == 0 {
//        return
//    }
//    // 如果添加数据超过了容量
//    if s.len+len(Data) > s.cap {
//        // C.realloc(指针,字节大小)
//        s.Data = C.realloc(s.Data, C.ulonglong(s.cap)*2*8)
//        s.cap = s.cap * 2
//    }
//    p := uintptr(s.Data)
//    for i := 0; i < s.len; i++ {
//        p += TAG
//    }
//    // 添加数据
//    for _, v := range Data {
//        *(*int)(unsafe.Pointer(p)) = v
//        p += TAG
//    }
//    // update the len
//    s.len = s.len + len(Data)
//}
//
//func (s *Slice) SearchData(index int) int {
//    if s == nil || s.Data == nil {
//        return -1
//    }
//
//    p := uintptr(s.Data)
//
//    for i := 0; i < s.len; i++ {
//        if *(*int)(unsafe.Pointer(p)) == index {
//            return i
//        }
//        p += TAG
//    }
//    return -1
//}
//func (s *Slice) DeleteDate(index int) {
//    if s == nil || s.Data == nil {
//        return
//    }
//    if index < 0 || index > s.len-1 {
//        return
//    }
//    p := uintptr(s.Data)
//    for i := 0; i < index; i++ {
//        p += TAG
//    }
//    temp := p
//    for i := index; i < s.len; i++ {
//        temp += TAG
//        *(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(temp))
//
//        p += TAG
//    }
//    s.len--
//
//}
//func (s *Slice) Insert(index int, Data int) {
//    if s == nil || s.Data == nil {
//        return
//    }
//    if index < 0 || index > s.len-1 {
//        return
//    }
//    if index == s.len-1 {
//        s.Append(Data)
//        return
//    }
//
//    // get the insert position
//    p := uintptr(s.Data)
//    for i := 0; i < index; i++ {
//        p += TAG
//    }
//    // get the last ptr pos
//    temp := uintptr(s.Data)
//    temp += TAG * uintptr(s.len)
//
//    // move the data to back
//    for i := s.len; i > index; i-- {
//        *(*int)(unsafe.Pointer(temp)) = *(*int)(unsafe.Pointer(temp - TAG))
//        temp -= TAG
//    }
//    *(*int)(unsafe.Pointer(p)) = Data
//    s.len++
//}
//func (s *Slice) Destroy() {
//    C.free(s.Data)
//    s.Data = nil
//    s.len = 0
//    s.cap = 0
//    s = nil
//}
