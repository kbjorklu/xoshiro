package xoshiro

import (
	"math/rand"
	"testing"
)

var _ rand.Source64 = (*Xoshiro512StarStar)(nil)

func TestXoshiro512StarStar(t *testing.T) {
	r := NewXoshiro512StarStar(42)
	// Generated using http://xoshiro.di.unimi.it/xoshiro512starstar.c .
	expected := []uint64{
		1546998764402558742,
		6990951692964543102,
		7962326261430671439,
		17084606997160102170,
		4101882904690527069,
		15166685478548380353,
		16153701568189242320,
		18407474307505565151,
		3534406837212445877,
		8444861339109729234,
	}
	for _, want := range expected {
		got := r.Uint64()
		if got != want {
			t.Error("incorrect random number sequence: got", got, "want", want)
		}
	}

	r.Jump()
	expected = []uint64{
		7693774656216124439,
		8548396176658720266,
		5502576858580612057,
		13348268550237079838,
		16266190634909969197,
		10977588625907609758,
		10853877200930025309,
		2381752334172876492,
		3447354402844146907,
		14849906996664834112,
	}
	for _, want := range expected {
		got := r.Uint64()
		if got != want {
			t.Error("incorrect random number sequence, got:", got, "want:", want)
		}
	}

}

func TestXoshiro512StarStarRandSource(t *testing.T) {
	r := NewXoshiro512StarStar(42)
	for i := 0; i < 10000; i++ {
		got := r.Int63()
		if got < 0 {
			t.Error("expected positive number, got:", got)
		}
	}
}
