package xoshiro

// SplitMix64 is the mix64variant13 pseudorandom number generator using
// GOLDEN_GAMMA (0x9e3779b97f4a7c15) state/seed increments, as described in
// http://gee.cs.oswego.edu/dl/papers/oopsla14.pdf . It is used for seeding the
// Xoshiro256StarStar generator, but can also be used standalone. It implements
// rand.Source and rand.Source64. The implementation is a direct port from
// http://xoshiro.di.unimi.it/splitmix64.c .
type SplitMix64 uint64

// Uint64 implements rand.Source64.
func (x *SplitMix64) Uint64() uint64 {
	*x += 0x9e3779b97f4a7c15
	z := uint64(*x)
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}

// Int63 implements rand.Source by returning the upper bits from x.Uint64.
func (x *SplitMix64) Int63() int64 {
	return int64(x.Uint64() >> 1)
}

// Seed implements rand.Source.
func (x *SplitMix64) Seed(seed int64) {
	*x = SplitMix64(seed)
}
