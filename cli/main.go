package main

import (
	"cli/reqests"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	ignoreNext := 0

	for _, arg := range args {
		if ignoreNext > 0 {
			ignoreNext -= 1
			continue
		}

		switch arg {
		case "test":
			reqests.CallTestGet()
			break

		case "xxx":
			fmt.Println("jest xxx")
			break

		default:
			fmt.Println("jest default")
		}

	}
}
