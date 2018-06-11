package xoshiro

import "math/bits"

// Xoshiro512StarStar is the Xoshiro512** pseudorandom number generator
// described at http://xoshiro.di.unimi.it/ . It implements rand.Source and
// rand.Source64. The implementation is a direct port from the reference C
// source: http://xoshiro.di.unimi.it/xoshiro512starstar.c .
type Xoshiro512StarStar struct {
	s0 uint64
	s1 uint64
	s2 uint64
	s3 uint64
	s4 uint64
	s5 uint64
	s6 uint64
	s7 uint64
}

// NewXoshiro512StarStar creates a new Xoshiro512StarStar PRNG.
func NewXoshiro512StarStar(seed int64) *Xoshiro512StarStar {
	x := &Xoshiro512StarStar{}
	x.Seed(seed)
	return x
}

// Uint64 implements rand.Source64 by returning the next Xoshiro random number.
func (x *Xoshiro512StarStar) Uint64() uint64 {
	result := bits.RotateLeft64(x.s1*5, 7) * 9

	t := x.s1 << 11

	x.s2 ^= x.s0
	x.s5 ^= x.s1
	x.s1 ^= x.s2
	x.s7 ^= x.s3
	x.s3 ^= x.s4
	x.s4 ^= x.s5
	x.s0 ^= x.s6
	x.s6 ^= x.s7

	x.s6 ^= t

	x.s7 = bits.RotateLeft64(x.s7, 21)

	return result
}

// Int63 implements rand.Source by returning the upper bits from x.Uint64.
func (x *Xoshiro512StarStar) Int63() int64 {
	return int64(x.Uint64() >> 1)
}

// Seed implements rand.Source by seeding the Xoshiro generator state using the
// first 8 outputs of a SplitMix64 generator seeded by seed.
func (x *Xoshiro512StarStar) Seed(seed int64) {
	var s SplitMix64
	s.Seed(seed)
	x.s0 = s.Uint64()
	x.s1 = s.Uint64()
	x.s2 = s.Uint64()
	x.s3 = s.Uint64()
	x.s4 = s.Uint64()
	x.s5 = s.Uint64()
	x.s6 = s.Uint64()
	x.s7 = s.Uint64()
}

// Jump is equivalent to 2^256 calls to Uint64.
func (x *Xoshiro512StarStar) Jump() {
	jumps := [...]uint64{0x33ed89b6e7a353f9, 0x760083d7955323be, 0x2837f2fbb5f22fae, 0x4b8c5674d309511c, 0xb11ac47a7ba28c25, 0xf1be7667092bcc1c, 0x53851efdb6df0aaf, 0x1ebbc8b23eaf25db}

	var s0 uint64
	var s1 uint64
	var s2 uint64
	var s3 uint64
	var s4 uint64
	var s5 uint64
	var s6 uint64
	var s7 uint64
	for _, jump := range jumps {
		var b uint
		for b = 0; b < 64; b++ {
			if jump&(1<<b) != 0 {
				s0 ^= x.s0
				s1 ^= x.s1
				s2 ^= x.s2
				s3 ^= x.s3
				s4 ^= x.s4
				s5 ^= x.s5
				s6 ^= x.s6
				s7 ^= x.s7
			}
			x.Uint64()
		}
	}
	x.s0 = s0
	x.s1 = s1
	x.s2 = s2
	x.s3 = s3
	x.s4 = s4
	x.s5 = s5
	x.s6 = s6
	x.s7 = s7
}
