package minStack

type MinStack struct {
	items map[int]int
	topIndex int
	minimumItems map[int]int
	minTopIndex int
}

func Constructor() MinStack {
    return MinStack {
		items: make(map[int]int),
		topIndex: -1,
		minimumItems: make(map[int]int),
		minTopIndex: -1,
	}
}

func (this *MinStack) Push(x int)  {
	this.topIndex++
	this.items[this.topIndex] = x
	if this.minTopIndex == -1 || x <= this.minimumItems[this.minTopIndex] {
		this.minTopIndex++
		this.minimumItems[this.minTopIndex] = x
	}
}

func (this *MinStack) Pop()  {
	if this.items[this.topIndex] == this.minimumItems[this.minTopIndex] {
		this.minTopIndex--
	}
	this.topIndex--
}

func (this *MinStack) Top() int {
    return this.items[this.topIndex]
}


func (this *MinStack) GetMin() int {
    return this.minimumItems[this.minTopIndex] 
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */