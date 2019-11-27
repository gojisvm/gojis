package gojis_test

import (
	"bytes"
	"testing"

	"github.com/gojisvm/gojis"
	"github.com/gojisvm/gojis/config"
	"github.com/gojisvm/gojis/tools/golden"
)

func TestHelloWorld(t *testing.T) {
	t.SkipNow() // skipped until API is implemented

	vm := gojis.NewVM(&config.Cfg{
		Debug: false,
	})

	var buf bytes.Buffer
	vm.SetConsole(gojis.NewConsole(&buf))

	_, _ = vm.Eval(`console.log("Hello World!");`)

	golden.Equal(t, "TestHelloWorld", buf.Bytes())
}

func TestInlineObject(t *testing.T) {
	t.SkipNow()

	vm := gojis.NewVM(&config.Cfg{
		Debug: false,
	})

	var buf bytes.Buffer
	vm.SetConsole(gojis.NewConsole(&buf))

	_, _ = vm.Eval(`
function foo(arg) {
	console.log(arg.func());
}

function gen(arg) {
	return () => arg;
}

foo(gen({func: () => "Output"})());
	`)

	golden.Equal(t, "TestInlineObject", buf.Bytes())
}
