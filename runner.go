package main

import (
	"github.com/dop251/goja"
)

func runCode(ctx vmMessageContext, args []string) error {
	vm := goja.New()
	vm.Set("send", func(msg string) {
		vmMessageSender(ctx, msg)
	})
	vm.Set("args", args)
	_, err := vm.RunString(ctx.Command.Code)
	return err
}
