package vm

import . "luago/api"

func move(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1        // convert to stack index
	vm.Copy(b, a) // copy b to a
}
