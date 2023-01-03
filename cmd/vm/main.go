package main

import "github.com/spatecon/deka/vm"

func main() {
	VM := vm.New()

	VM.Exec("1 + 2")
}
