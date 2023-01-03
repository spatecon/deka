package vm

import "unsafe"

type VM struct {
	ip unsafe.Pointer
}

func New() *VM {
	return &VM{}
}

func (vm *VM) Exec(program string) {
	// 1. parse the program
	// ast := parser.Parse(program)

	// 2. compile the program
	// code := compiler.Compile(ast)
	code := []byte{OP_HALT}

	// set instruction pointer to the beginning of the program

	vm.ip = unsafe.Pointer(&code[0])

	vm.eval()
}

func (vm *VM) readByte() byte {
	b := *(*byte)(vm.ip)
	vm.ip = unsafe.Pointer(uintptr(vm.ip) + unsafe.Sizeof(b))
	return b
}

func (vm *VM) eval() {
	for {
		switch vm.readByte() {
		case 0x00:
			return
		}
	}
}
