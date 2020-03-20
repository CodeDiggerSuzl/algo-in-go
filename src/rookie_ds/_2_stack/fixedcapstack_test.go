package _2_stack

import (
	"fmt"
	"testing"
)

func TestFixedCapStack_NewFixedCapStack(t *testing.T) {
	fixedCapStack := NewFixedCapStack(10)
	empty := fixedCapStack.isEmpty()
	fmt.Println(empty)
	pop, err := fixedCapStack.pop()
	if err != nil {
		t.Fatal("empty")
	}
	fmt.Println(pop)

}
