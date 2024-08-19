package tasks

import (
	"fmt"
	"os"
)

func ErrorArgsCheck(msg string, lenOfArgs int) {
	if len(os.Args) < lenOfArgs {
		fmt.Println(msg)
		os.Exit(1)
	}
}
