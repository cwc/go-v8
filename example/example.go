package main

import (
	"fmt"
	"github.com/mattn/go-v8"
)

func main() {
	v8ctx := v8.NewContext()

	// setup console.log()
	v8ctx.Eval(`
	this.console = { "log": function(args) { _console_log.apply(null, arguments) }}`)
	v8ctx.AddFunc("_console_log", func(args ...interface{}) interface{} {
		fmt.Printf("Go console log: ")
		for i := 0; i < len(args); i++ {
			fmt.Printf("%v ", args[i])
		}
		fmt.Println()
		return ""
	})

	ret := v8ctx.MustEval(`
	var a = 1;
	var b = 'B'
	a += 2;
	a;
	`)
	fmt.Println("Eval result:", int(ret.(float64))) // 3

	v8ctx.Eval(`console.log(a + '年' + b + '組 金八先生！', 'something else')`) // 3b
	v8ctx.Eval(`console.log("Hello World, こんにちわ世界")`)                    // john

	v8ctx.AddFunc("func_call", func(args ...interface{}) interface{} {
		f := func(args ...interface{}) interface{} {
			return "V8"
		}
		ret, _ := args[0].(v8.V8Function).Call("Go", 2, 1, f)
		return ret
	})

	fmt.Println(v8ctx.MustEval(`
	func_call(function() {
		return "Hello " + arguments[0] + (arguments[1] - arguments[2])
			+ ", Hello " + arguments[3]();
	})
	`).(string))
}
