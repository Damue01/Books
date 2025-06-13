package api

type LuaType int
type ArithOp int    // 算术操作符
type CompareOp int  // 比较操作符

type LuaState interface {
	/* basic stack operations */
	GetTop() int             // 返回栈顶索引
	AbsIndex(idx int) int    // 将索引转换为绝对索引
	CheckStack(n int) bool   // 检查栈空间是否足够
	Pop(n int)               // 弹出n个元素
	Copy(fromIdx, toIdx int) // 将fromIdx位置的值复制到toIdx位置
	PushValue(idx int)       // 将idx位置的值压入栈顶
	Replace(idx int)         // 将栈顶的值替换为idx位置的值
	Insert(idx int)          // 将栈顶的值插入到idx位置
	Remove(idx int)          // 删除idx位置的值
	Rotate(idx, n int)       // 将idx位置的值向下移动n个位置
	SetTop(idx int)          // 设置栈顶索引为idx
	/* access functions (stack -> Go) */
	TypeName(tp LuaType) string        // 返回LuaType对应的名称
	Type(idx int) LuaType              // 返回idx位置的值的类型
	IsNone(idx int) bool               // 判断idx位置的值是否为nil
	IsNil(idx int) bool                // 判断idx位置的值是否为nil
	IsNoneOrNil(idx int) bool          // 判断idx位置的值是否为nil或None
	IsBoolean(idx int) bool            // 判断idx位置的值是否为布尔类型
	IsInteger(idx int) bool            // 判断idx位置的值是否为整数类型
	IsNumber(idx int) bool             // 判断idx位置的值是否为数字类型
	IsString(idx int) bool             // 判断idx位置的值是否为字符串类型
	ToBoolean(idx int) bool            // 将idx位置的值转换为布尔类型
	ToInteger(idx int) int64           // 将idx位置的值转换为整数64类型
	ToIntegerX(idx int) (int64, bool)  // 将idx位置的值转换为整数64类型，返回是否成功
	ToNumber(idx int) float64          // 将idx位置的值转换为浮点数类型
	ToNumberX(idx int) (float64, bool) // 将idx位置的值转换为浮点数类型，返回是否成功
	ToString(idx int) string           // 将idx位置的值转换为字符串类型
	ToStringX(idx int) (string, bool)  // 将idx位置的值转换为字符串类型，返回是否成功
	/* push functions (Go -> stack) */
	PushNil()             // 压入nil值
	PushBoolean(b bool)   // 压入布尔值
	PushInteger(i int64)  // 压入整数值
	PushNumber(n float64) // 压入浮点数值
	PushString(s string)  // 压入字符串值
	/* comparison and arithmetic functions */
	Arith(op ArithOp) // 执行算术操作
	Compare(idx1, idx2 int, op CompareOp) bool // 比较两个值
	Len(idx int) // 返回idx位置的值的长度
	Concat(n int) // 将n个字符串连接起来
}
