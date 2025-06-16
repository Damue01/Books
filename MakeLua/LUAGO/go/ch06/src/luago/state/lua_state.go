package state

import "luago/binchunk"

type luaState struct {
	stack *luaStack           // 栈
	proto *binchunk.Prototype // 当前执行的函数原型
	pc    int                 // 程序计数器
}

func New(stackSize int, proto *binchunk.Prototype) *luaState {
	return &luaState{
		stack: newLuaStack(stackSize),
		proto: proto,
		pc:    0,
	}
}
