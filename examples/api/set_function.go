package main

import (
	"github.com/gojisvm/gojis"
)

// #nosec
func main() {
	vm := gojis.NewVM()

	vm.SetFunction("greet", func(gojis.Args) gojis.Object {
		_, _ = vm.Eval(`console.log("Hello World!");`)

		// or

		_, _ = vm.Lookup("eval").CallWithArgs(`console.log("Hello World!");`)

		// or

		_, _ = vm.Lookup("console").Lookup("log").CallWithArgs("Hello", "World!")

		// or

		console := vm.Lookup("console")
		consoleLog := console.Lookup("log")
		_, _ = consoleLog.CallWithArgs("Hello", "World!")
		_, _ = consoleLog.CallWithArgs("I am reusable!")

		return gojis.Undefined
	})

	_, _ = vm.Eval(`greet();`)
	/*
		prints:
		Hello World!
		Hello World!
		Hello World!
		Hello World!
		I am reusable!
	*/

	alerts := make(chan string, 1)
	// consume alerts somewhere else, like:
	// go drainAlerts(alerts)

	vm.SetFunction("alert", func(args gojis.Args) gojis.Object {
		if args.Len() == 0 {
			panic("No argument provided.")
		}

		val := args.Get(0)
		if val.Type() != gojis.TypeString {
			panic("Argument has to be a string")
		}
		alerts <- val.Value().(string)

		return nil
	})

	_, _ = vm.Eval(`alert("Hello");`)
}
