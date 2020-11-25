The test [program](interesting.go) tests the math/rand PRNG through the rand.Rand methods Int31, Int63, Uint32 and Uint64 (so there are really four separate tests). Only the first and second positions in each pseudo-random stream are examined (i.e., each method is called twice for each seed). What the program does for each method and position is iterate through all effectively possible seed values, counting the number of times that each value from a predetermined range (-128 to 127) is output.

The most interesting part of the test [results](data/results.txt) are the reports of which numbers and how many numbers don't appear at all, because if a value is never output in the test under a certain configuration (method and position in pseudo-random stream), then it's not possible for those values ever to appear in that configuration, even with the Source initialized with an arbitrary int64 seed value.

Ideally each count should be very close to one, with very few counts of zero.

A clarification: the Int31 and Int63 methods return only nonnegative values, thus the test shows some redundant information regarding those methods (all the negative "We didn't see any" entries are redundant, as are the counts of negative values). Contrast that with Uint32 and Uint64: their unsigned 32-bit and 64-bit return values are reinterpreted as 32-bit or 64-bit two's complement signed values (just so we'd consider more values as interesting in the test), so the full range of their "interesting" values really is interesting in the test results.

#### Summary of the results for Int31 at position 0:

Outputs that never happen: 7 8 9 10 16 17 18 21 22 25 29 31 34 35 38 41 42 47 49 58 59 61 62 63 64 65 66 69 72 73 74 79 84 86 87 90 91 92 93 94 95 97 99 101 104 105 106 108 112 116 117 120 121 127

Thus 54 out of the 128 (0.422) interesting values are never generated.

#### Summary of the results for Int31 at position 1:

Outputs that never happen: 0 1 8 9 10 11 12 17 19 20 21 22 24 31 32 36 37 38 39 42 43 48 49 53 56 57 58 60 62 64 75 81 83 88 90 92 93 95 100 102 104 108 109 113 115 118 119 121 122 125

Thus 50 out of the 128 (0.391) interesting values are never generated.

It's curious how the method Int31 manages to produce the number 111 (decimal) a whopping 7 times as the second value in the pseudo-random sequence. I think this means 111 (but surely not only 111) is disproportionately likely to appear even assuming an arbitrary int64 seed value.

#### Summary of the results for Int63:

No interesting values are ever output for either position 0 or position 1. This basically follows from the fact that 2^63 is so much greater than 2^31 - 1, causing the outputs to get "lost" in the huge int64 empty space.

#### Summary of the results for Uint32 at position 0:

Outputs that never happen: -128 -127 -126 -125 -124 -123 -122 -121 -120 -119 -118 -117 -116 -115 -113 -112 -110 -108 -105 -103 -101 -100 -99 -98 -94 -91 -89 -88 -87 -86 -84 -83 -82 -78 -77 -76 -74 -73 -72 -71 -70 -69 -67 -65 -63 -62 -61 -59 -56 -53 -52 -51 -50 -47 -45 -44 -43 -41 -40 -37 -36 -35 -32 -30 -29 -28 -23 -22 -21 -20 -19 -18 -17 -16 -15 -14 -11 -8 -5 -4 -2 -1 0 2 5 6 11 12 14 15 16 17 18 19 20 21 23 24 27 28 30 32 33 34 35 36 37 41 42 43 44 45 49 50 51 54 58 59 61 62 63 65 67 68 69 70 71 72 75 76 77 78 81 82 83 84 85 87 89 90 93 94 95 96 98 99 101 103 104 107 108 113 115 116 117 118 119 121 122 123 124 125 126 127

Thus 164 out of the 256 (0.641) interesting values are never generated.

#### Summary of the results for Uint32 at position 1:

Outputs that never happen: -128 -127 -126 -125 -124 -123 -122 -121 -120 -119 -118 -117 -116 -115 -114 -113 -112 -111 -110 -109 -107 -105 -104 -103 -100 -99 -98 -96 -94 -93 -87 -86 -85 -84 -83 -82 -78 -77 -76 -75 -74 -73 -72 -68 -67 -66 -65 -64 -63 -59 -55 -53 -52 -50 -49 -46 -45 -44 -43 -42 -41 -40 -39 -38 -36 -34 -32 -30 -28 -26 -25 -24 -23 -22 -21 -19 -18 -17 -15 -12 -11 -10 -9 -7 -5 -3 -2 -1 0 1 2 3 5 7 9 10 13 15 16 17 18 19 20 21 22 23 24 25 26 28 30 32 34 35 36 38 39 40 41 42 43 44 45 48 49 50 53 55 59 62 63 64 65 66 68 72 73 74 75 76 77 78 79 84 85 86 87 92 95 96 97 98 99 101 105 106 107 109 110 112 113 114 115 116 117 118 120 121 123 124 125 127

Thus 172 out of the 256 (0.672) interesting values are never generated.

#### Summary of the results for Uint64:

Same situation as for Int63.

### What does this mean in practice?

For example, if we have

```
	prg31 := rand.New(rand.NewSource(anything))
	a31 := prg31.Int31()
	b31 := prg31.Int31()

	prg63 := rand.New(rand.NewSource(anything))
	a63 := prg63.Int63()
	b63 := prg63.Int63()
```

```a31``` will never be 7, but may be 0; ```b31``` will never be 0, but may be 7; and neither ```a63``` nor ```b63``` are ever in the range of "interesting" values.
