// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package popcountByShift

import (
	"golang_training/CH02/ex03"
	"testing"

	"gopl.io/ch2/popcount"
)

func TestPopcountByShift(t *testing.T) {
	const (
		bit0   = 0
		bit1   = 1
		bit10  = 2
		bit255 = 8
		bitF16 = 64
	)
	count := PopCount(0)
	if count != bit0 {
		t.Errorf("Calculated value: %d Expected value: %d", count, bit0)
	}
	count = PopCount(1)
	if count != bit1 {
		t.Errorf("Calculated value: %d Expected value: %d", count, bit1)
	}
	count = PopCount(10)
	if count != bit10 {
		t.Errorf("Calculated value: %d Expected value: %d", count, bit10)
	}
	count = PopCount(255)
	if count != bit255 {
		t.Errorf("Calculated value: %d Expected value: %d", count, bit255)
	}
	count = PopCount(0xFFFFFFFFFFFFFFFF)
	if count != bitF16 {
		t.Errorf("Calculated value: %d Expected value: %d", count, bitF16)
	}
}

func BenchmarkPopcountByShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0xFFFFFFFFFFFFFFFF)
	}
}

func BenchmarkPopcount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount2.PopCount(0xFFFFFFFFFFFFFFFF)
	}
}

func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0xFFFFFFFFFFFFFFFF)
	}
}
