package xoshiro

import (
	"math/rand"
	"testing"
)

var _ rand.Source64 = (*SplitMix64)(nil)

func TestSplitMix64(t *testing.T) {
	var r SplitMix64
	r.Seed(42)
	// Generated using http://xoshiro.di.unimi.it/splitmix64.c .
	expected := []uint64{
		13679457532755275413,
		2949826092126892291,
		5139283748462763858,
		6349198060258255764,
		701532786141963250,
		16015981125662989062,
		4028864712777624925,
		14769051326987775908,
		6270620877612482005,
		11408980392250668974,
	}
	for _, want := range expected {
		got := r.Uint64()
		if got != want {
			t.Error("incorrect random number sequence: got", got, "want", want)
		}
	}
}

func TestSplitMix64RandSource(t *testing.T) {
	var r SplitMix64
	r.Seed(42)
	for i := 0; i < 10000; i++ {
		got := r.Int63()
		if got < 0 {
			t.Error("expected positive number, got:", got)
		}
	}
}
