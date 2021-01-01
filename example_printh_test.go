package printh_test

import (
	"fmt"

	"github.com/felipeneuwald/printh"
)

func ExampleInfo() {
	printh.Info("info message")
	// Output: 2009/11/10 23:00:00 INFO  [main.main] info message
}

func ExampleDebug() {
	d := true
	printh.Debug(d, "debug message")
	// Output: 2009/11/10 23:00:00 DEBUG [main.main] debug message
}

func ExampleErr() {
	err := fmt.Errorf("new error")
	printh.Err(err)
	// Output: 2009/11/10 23:00:00 ERR   [main.main] (new error)
}
