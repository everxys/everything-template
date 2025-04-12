package main

import (
	"flag"

	"everything-template/internal/bootstrap"
)

func main() {
	env := flag.String("e", "dev", "environment to run (dev, test, prod)")
	flag.Parse()

	bootstrap.Run(*env)
}
