package _0_array

import "testing"

func TestArray_Insert(t *testing.T) {
	capacity := 10
	array := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		_, err := array.Insert(uint(i), i+1)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	array.Print()
	_, err := array.Delete(33)
	if err != nil {
		t.Fatal(err.Error())
	}
	array.Print()
}

func TestArray_Find(t *testing.T) {
	arr := NewArray(uint(10))
	for i := 0; i < 10; i++ {
		_, err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()
}
