package problem/flattenNestedArray

type NestedInteger struct { }

func (this NestedInteger) IsInteger() bool {}
func (this NestedInteger) GetInteger() int {}
func (n *NestedInteger) SetInteger(value int) {}
func (this *NestedInteger) Add(elem NestedInteger) {}
func (this NestedInteger) GetList() []*NestedInteger {}
type NestedIterator struct {
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return nil
}

func (this *NestedIterator) Next() int {
    return nil
}

func (this *NestedIterator) HasNext() bool {
    return nil
}