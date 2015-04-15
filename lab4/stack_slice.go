// +build !solution

// Leave an empty line above this comment.
package lab4

const DefaultCap = 10

type SliceStack struct {
	slice []interface{}
	top   int
}

func NewSliceStack() *SliceStack {
	return &SliceStack{
		slice: make([]interface{}, DefaultCap),
		top:   -1,
	}
}

func (ss *SliceStack) Len() int {
	return ss.top + 1
}

func (ss *SliceStack) Push(value interface{}) {
	ss.top++

	if ss.top == len(ss.slice) {
		// Reallocate
		newSlice := make([]interface{}, len(ss.slice)*2)
		copy(newSlice, ss.slice)
		ss.slice = newSlice
	}

	ss.slice[ss.top] = value
}

func (ss *SliceStack) Pop() (value interface{}) {
	if ss.top > -1 {
		defer func() { ss.top-- }()
		return ss.slice[ss.top]
	}

	return nil
}
