package _2_stack

import "errors"

type FixedCapStack struct {
	len int
	N   []int
}

func NewFixedCapStack(n int) *FixedCapStack {
	if n < 0 {
		return nil
	}
	return &FixedCapStack{0, make([]int, n)}
}

func (a *FixedCapStack) Size() int {
	return a.len
}

func (a *FixedCapStack) IsEmpty() bool {
	return a.len == 0
}

func (a *FixedCapStack) Pop() (int, error) {
	if a.IsEmpty() {
		return -1, errors.New("empty")
	}
	i := a.N[a.len]
	a.len--
	return i, nil
}

func (a *FixedCapStack) Push(v int) {
	a.len++
	a.N[a.len] = v
}

func (a *FixedCapStack) resize() {

}
