package minStack

import "testing"

func TestMinStack(t *testing.T) {

	obj := Constructor()
	obj.Push(5)
	obj.Push(9)
	obj.Push(4)
	param_3 := obj.Top()
	param_4 := obj.GetMin()

	obj.Pop()

	t.Log(param_3, param_4)

	param_3 = obj.Top()
	param_4 = obj.GetMin()

	t.Log(param_3, param_4)
}
