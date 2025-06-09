package state

type luaState struct {
	stack *luaStack // 栈
}

func New() *luaState {
	return &luaState{
		stack: newLuaStack(20), // 初始化栈，初始大小为20
	}
}
