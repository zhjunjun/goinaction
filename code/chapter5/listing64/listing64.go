package main

import (
	"fmt"
	"goinaction/code/chapter5/listing64/counters"
)

func main() {

	// counter := counters.alterCounter(10)

	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}
