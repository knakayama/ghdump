package utils

import (
	"fmt"
	"os"
)

func Dieif(err error) {
	if err != nil {
		die(err)
	}
}

func die(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
