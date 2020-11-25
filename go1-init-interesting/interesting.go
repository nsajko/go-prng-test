// Copyright 2020 Neven Sajko. All rights reserved. See LICENSE for a license.

// Testing Go's math/rand PRNG with similar aims as in https://www.pcg-random.org/posts/cpp-seeding-surprises.html
package main

import (
	"fmt"
	"math/rand"
)

const (
	// Determines the range of "interesting" generated values.
	testRange = 1 << 7

	// We test the first N positions of a PRNG's stream.
	N = 2

	rngMethods = 4
)

func experiment(m int, a *[N][2 * testRange]int, ch chan struct{}) {
	for seed := 0; seed < (1<<31)-1; seed++ {
		prg := rand.New(rand.NewSource(int64(seed)))

		for i := 0; i < N; i++ {
			var v int64
			switch m {
			case 0:
				v = int64(prg.Int31())
			case 1:
				v = prg.Int63()
			case 2:
				v = int64(int32(prg.Uint32()))
			case 3:
				v = int64(prg.Uint64())
			}
			if -testRange <= v && v < testRange {
				// Record the interesting result.
				a[i][v+testRange]++
			}
		}
	}
	ch <- struct{}{}
}

func main() {
	rngMethodStrings := [rngMethods]string{
		"Method rand.Rand.Int31\n",
		"Method rand.Rand.Int63\n",
		"Method rand.Rand.Uint32\n",
		"Method rand.Rand.Uint64\n"}

	a := [rngMethods][N][2 * testRange]int{}

	// Perform experiments.
	ch := make(chan struct{}, rngMethods)
	for m := 0; m < rngMethods; m++ {
		go experiment(m, &a[m], ch)
	}
	for m := 0; m < rngMethods; m++ {
		<-ch
	}

	// Print results.
	for m := 0; m < rngMethods; m++ {
		fmt.Println(rngMethodStrings[m])

		for i := 0; i < N; i++ {
			fmt.Printf("Position %d\n", i)

			// Report values that didn't occur in the experiment.
			fmt.Print("We didn't see any of these: ")
			for v := -testRange; v < testRange; v++ {
				if a[m][i][v+testRange] == 0 {
					fmt.Print(" ", v)
				}
			}
			fmt.Println("\n")

			// Report how many times each value we kept track of occured. Give
			// frequency distribution (or a part of it, at least).
			fmt.Println("Counts:")
			dist := [1 << 12]int{}
			for v := -testRange; v < testRange; v++ {
				cnt := a[m][i][v+testRange]

				// Print count.
				fmt.Printf("%4d: %10d\n", v, cnt)

				// Update frequency distribution.
				dist[cnt]++
			}
			fmt.Println("\nFrequency distribution:")
			last := 0
			for i, j := range dist {
				if j != 0 {
					last = i
				}
			}
			for i := 0; i <= last; i++ {
				fmt.Printf("%5d: %5d\n", i, dist[i])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
