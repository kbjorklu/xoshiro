// Package xoshiro implements the Xoshiro256** and SplitMix64 pseudorandom
// number generators described at http://xoshiro.di.unimi.it/ .
package xoshiro

import "math/bits"

// Xoshiro256StarStar is the Xoshiro256** pseudorandom number generator
// described at http://xoshiro.di.unimi.it/ . It implements rand.Source and
// rand.Source64. The implementation is a direct port from the reference C
// source: http://xoshiro.di.unimi.it/xoshiro256starstar.c .
type Xoshiro256StarStar struct {
	s0 uint64
	s1 uint64
	s2 uint64
	s3 uint64
}

// NewXoshiro256StarStar creates a new NewXoshiro256StarStar PRNG.
func NewXoshiro256StarStar(seed int64) *Xoshiro256StarStar {
	x := &Xoshiro256StarStar{}
	x.Seed(seed)
	return x
}

// Uint64 implements rand.Source64 by returning the next Xoshiro random number.
func (x *Xoshiro256StarStar) Uint64() uint64 {
	result := bits.RotateLeft64(x.s1*5, 7) * 9

	t := x.s1 << 17

	x.s2 ^= x.s0
	x.s3 ^= x.s1
	x.s1 ^= x.s2
	x.s0 ^= x.s3

	x.s2 ^= t

	x.s3 = bits.RotateLeft64(x.s3, 45)

	return result
}

// Int63 implements rand.Source by returning the upper bits from x.Uint64.
func (x *Xoshiro256StarStar) Int63() int64 {
	return int64(x.Uint64() >> 1)
}

// Seed implements rand.Source by seeding the Xoshiro generator state using the
// first 4 outputs of a SplitMix64 generator seeded by seed.
func (x *Xoshiro256StarStar) Seed(seed int64) {
	var s SplitMix64
	s.Seed(seed)
	x.s0 = s.Uint64()
	x.s1 = s.Uint64()
	x.s2 = s.Uint64()
	x.s3 = s.Uint64()
}

// Jump is equivalent to 2^128 calls to Uint64.
func (x *Xoshiro256StarStar) Jump() {
	jumps := [...]uint64{0x180ec6d33cfd0aba, 0xd5a61266f0c9392c, 0xa9582618e03fc9aa, 0x39abdc4529b1661c}

	var s0 uint64
	var s1 uint64
	var s2 uint64
	var s3 uint64
	for _, jump := range jumps {
		var b uint
		for b = 0; b < 64; b++ {
			if jump&(1<<b) != 0 {
				s0 ^= x.s0
				s1 ^= x.s1
				s2 ^= x.s2
				s3 ^= x.s3
			}
			x.Uint64()
		}
	}
	x.s0 = s0
	x.s1 = s1
	x.s2 = s2
	x.s3 = s3
}
