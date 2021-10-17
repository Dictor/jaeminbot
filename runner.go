package main

import (
	"github.com/dop251/goja"
)

func runCode(ctx vmMessageContext) error {
	vm := goja.New()
	vm.Set("send", func(msg string) {
		vmMessageSender(ctx, msg)
	})
	_, err := vm.RunString(ctx.Command.Code)
	return err
}
