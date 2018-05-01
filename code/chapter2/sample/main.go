package main

import (
	"log"
	"os"

	_ "./matchers"
	"./search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	// run
	search.Run("president")
}
