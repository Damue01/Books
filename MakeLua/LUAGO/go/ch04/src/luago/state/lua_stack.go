package state

type luaStack struct {
	slots []luaValue // slots存放值
	top   int        // top记录栈顶索引
}

func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
}

func (self *luaStack) check(n int) {
	free := len(self.slots) - self.top
	for i := free; i < n; i++ {
		self.slots = append(self.slots, nil) // 扩容
	}
}

func (self *luaStack) push(val luaValue) {
	if self.top == len(self.slots) {
		panic("lua stack overflow")
	}
	self.slots[self.top] = val
	self.top++
}

func (self *luaStack) pop() luaValue {
	if self.top < 1 {
		panic("lua stack underflow")
	}
	self.top--
	val := self.slots[self.top]
	self.slots[self.top] = nil // 清除引用，避免内存泄漏
	return val
}

// 将索引转换成绝对索引
func (self *luaStack) absIndex(idx int) int {
	if idx > 0 {
		return idx
	}
	return idx + self.top + 1
}

func (self *luaStack) isValid(idx int) bool {
	absIdx := self.absIndex(idx)
	return absIdx > 0 && absIdx <= self.top
}

func (self *luaStack) get(idx int) luaValue {
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		return self.slots[absIdx-1] // Lua索引从1开始，所以需要减1
	}
	return nil
}

func (self *luaStack) set(idx int, val luaValue) {
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		self.slots[absIdx-1] = val
		return
	}
	panic("invalid index: " + string(idx))
}

func (self *luaStack) reverse(from, to int) {
	slots := self.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}
