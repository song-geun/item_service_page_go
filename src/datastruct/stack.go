package datastruct

type stack struct {
	value     int
	next      *stack
	prev      *stack
	Push_back func(int)
	isempty   bool
	Isempty   func() bool
	Top       func() int
	Del       func()
	Pop       func()
}

var Stack stack
var nowStack stack

func New() {
	Stack.Push_back = push_back
	Stack.isempty = true
	Stack.Isempty = isempty
	Stack.Top = top
	Stack.Del = del
	Stack.prev = &Stack
	nowStack = Stack
	nowStack.isempty = false
	Stack.next = &nowStack
}

func push_back(val int) {
	nowStack = stack{}
	nowStack.next = nil
	nowStack.value = val
	nowStack.Push_back = push_back
	nowStack.isempty = false
	Stack.prev = Stack.next
	Stack.next = &nowStack
}
func top() int {
	if Stack.prev.isempty {
		return -1
	}
	return Stack.prev.value
}
func isempty() bool {
	return Stack.isempty
}
func del() {
	Stack.prev = Stack.next
}
