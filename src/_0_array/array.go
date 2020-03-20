package _0_array

import (
	"errors"
	"fmt"
)

type Array struct {
	data   []int
	length uint
}

// init memory for array
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

// get Len
func (a *Array) Len() uint {
	return a.length
}

func (a *Array) IsIndexOutOfRange(index uint) bool {
	if index >= uint(cap(a.data)) {
		return true
	}
	return false
}

// find
func (a *Array) Find(index uint) (int, error) {
	if a.IsIndexOutOfRange(index) {
		return -1, errors.New("index out of range")
	}
	return a.data[index], nil
}

// insert
func (a *Array) Insert(index uint, value int) (bool, error) {
	// if index > cap(a.data) {
	if a.Len() == uint(cap(a.data)) {
		return false, errors.New("full array")
	}
	if index < 0 || index > uint(cap(a.data)) {
		return false, errors.New("index out fo range")
	}
	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = value
	a.length++
	return true, nil
}

// insert to tail
func (a *Array) Insert2Tail(v int) (bool, error) {
	return a.Insert(a.Len(), v)
}

// delete
func (a *Array) Delete(index uint) (int, error) {
	if a.IsIndexOutOfRange(index) {
		return -1, errors.New("out of range")
	}
	v := a.data[index]

	for i := index; i < a.Len(); i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return v, nil
}

func (a *Array) Print() {
	var format string
	for i := uint(0); i < a.Len(); i++ {
		format += fmt.Sprintf("|%+v", a.data[i])
	}
	fmt.Println(format)
}
