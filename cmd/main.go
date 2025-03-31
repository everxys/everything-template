package main

import (
	"everything-template/internal/bootstrap"
	"flag"
)

func main() {
	env := flag.String("e", "dev", "environment to run (dev, test, prod)")
	flag.Parse()

	bootstrap.Run(*env)
}
