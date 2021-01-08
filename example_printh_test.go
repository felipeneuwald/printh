package printh_test

import (
	"fmt"
	"log"
	"os"

	"github.com/felipeneuwald/printh"
)

func ExampleInfo() {
	printh.Info("Info message")
}

func ExampleInfo_stdout() {
	log.SetOutput(os.Stdout)
	printh.Info("Info message printed to os.Stdout")
}

func ExampleDebug() {
	d := true
	printh.Debug(d, "Debug message")
}

func ExampleErr() {
	err := fmt.Errorf("New error")
	printh.Err(err)
}

func ExampleErrFatal() {
	err := fmt.Errorf("New fatal error")
	printh.ErrFatal(err)
	fmt.Println("This code will not run")
}
