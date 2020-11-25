// Copyright 2020 Neven Sajko. All rights reserved. See LICENSE for a license.

// Counts the int64 values that it is possible to obtain using rand.Rand.Int63 as the first
// output of an arbitrarily seeded source.
//
// This requires around 100 GB of RAM, so I ran it on the c6g.metal EC2 instance, which
// has 128 GB (it also has 60 ARM CPU cores that were left unused in this case). See
// the data directory for the results.
package main

import (
	"fmt"
	"math/rand"
	"runtime/debug"
)

func main() {
	debug.SetGCPercent(20)

	data := make(map[int64]struct{}, 1e9)
	for seed := 0; seed < (1<<31)-1; seed++ {
		data[rand.New(rand.NewSource(int64(seed))).Int63()] = struct{}{}
		if seed&0x1ffffff == 0x1000000 {
			fmt.Printf("%08x: %12d\n", uint64(seed), len(data))
		}
	}

	fmt.Printf("\nCount of all the int64 values obtainable through rand.Rand.Int63: %d\n", len(data))
}
