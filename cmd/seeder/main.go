package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ternakkode/packform-backend/internal/seed"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			seed.Exec(args[1:])
		default:
			pusage()
		}
	} else {
		pusage()
	}
}

func pusage() {
	fmt.Println(`Usage: seed [seedNames...]`)
	os.Exit(2)
}
