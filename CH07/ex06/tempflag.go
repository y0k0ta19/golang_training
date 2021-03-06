// Copyright © 2016 "Shun Yokota" All rights reserved

// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"
)

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
