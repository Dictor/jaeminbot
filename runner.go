package main

import (
	"fmt"

	"github.com/dop251/goja"

	// This initializes gpython for runtime execution and is critical.
	// It defines forward-declared symbols and registers native built-in modules, such as sys and time.
	//_ "github.com/go-python/gpython/stdlib"

	// This is the primary import for gpython.
	// It contains all symbols needed to fully compile and run python.
	"github.com/go-python/gpython/py"
)

func runJavascriptCode(ctx vmMessageContext, args []string) error {
	vm := goja.New()
	vm.Set("send", func(msg string) {
		vmMessageSender(ctx, msg)
	})
	vm.Set("args", args)
	_, err := vm.RunString(ctx.Command.Code)
	return err
}

func runPythonCode(ctx vmMessageContext, args []string) error {
	pyCtx := py.NewContext(py.DefaultContextOpts())
	defer pyCtx.Close()

	code, err := py.Compile(ctx.Command.Code, "<jaeminbot code>", py.ExecMode, 0, true)
	if err != nil {
		return err
	}

	mainModule, err := pyCtx.GetModule("main")
	if err != nil {
		return nil
	}

	/*
		send := func(msg string) {
			vmMessageSender(ctx, msg)
		}
	*/

	result, err := mainModule.Context.RunCode(code, mainModule.Globals, mainModule.GetDict(), nil)
	if err != nil {
		return err
	}
	vmMessageSender(ctx, fmt.Sprintf("%s", result))
	return nil
}
