// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import "fmt"

func main() {
	type S []S
	s := make(S, 0, 1)
	fmt.Printf("%p\n", s)
	s = append(s, s)
	fmt.Println(len(s))
	fmt.Printf("%p %d\n", s, s[0])
	fmt.Println(s)
	/*
		type CycleSlice []CycleSlice
		var cycleSlice CycleSlice
		cycleSlice = append(cycleSlice, cycleSlice)
		cycleSlice = make(CycleSlice, 1)
		cycleSlice[0] = cycleSlice
		for i := range cycleSlice {
			fmt.Println(cycleSlice[i])
		}
		fmt.Printf("cycleSice: %v\n", cycleSlice)

		type cycle struct {
			value int
			tail  *cycle
		}
		var c cycle
		c = cycle{42, &c}
	*/
}
