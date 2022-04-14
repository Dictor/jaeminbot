package main

import (
	"github.com/dop251/goja"

	// This initializes gpython for runtime execution and is critical.
	// It defines forward-declared symbols and registers native built-in modules, such as sys and time.
	_ "github.com/go-python/gpython/modules"

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

	mainModule, err := pyCtx.ModuleInit(&py.ModuleImpl{CodeSrc: "print('jaeminbot running context ready!')"})
	if err != nil {
		return err
	}

	send := func(module py.Object, args py.Tuple) (py.Object, error) {
		msg, ok := args[0].(py.String)
		if !ok {
			return nil, py.ExceptionNewf(py.TypeError, "jaemin_send의 첫번째 매개변수는 반드시 문자열이여야합니다!")
		}
		vmMessageSender(ctx, string(msg))
		return nil, nil
	}
	pysend, err := py.NewMethod("jaemin_send", send, 0, "")
	if err != nil {
		return err
	}
	py.SetAttrString(mainModule.Globals, "jaemin_send", pysend)

	_, err = mainModule.Context.RunCode(code, mainModule.Globals, mainModule.GetDict(), nil)
	if err != nil {
		return err
	}
	return nil
}
