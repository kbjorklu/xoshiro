package xoshiro

import (
	"math/rand"
	"testing"
)

var _ rand.Source64 = (*Xoshiro256StarStar)(nil)

func TestXoshiro256StarStar(t *testing.T) {
	var r Xoshiro256StarStar
	r.Seed(42)
	// Generated using http://xoshiro.di.unimi.it/xoshiro256starstar.c .
	expected := []uint64{
		1546998764402558742,
		6990951692964543102,
		12544586762248559009,
		17057574109182124193,
		18295552978065317476,
		14199186830065750584,
		13267978908934200754,
		15679888225317814407,
		14044878350692344958,
		10760895422300929085,
	}
	for _, want := range expected {
		got := r.Uint64()
		if got != want {
			t.Error("incorrect random number sequence: got", got, "want", want)
		}
	}

	r.Jump()
	expected = []uint64{
		10214096832459272891,
		3817095900452140333,
		16721143367497718667,
		13772684382182216847,
		1752989775392318229,
		18425343068764143960,
		17731448209264309203,
		3783290459727839037,
		10211543263381176166,
		987574451100413377,
	}
	for _, want := range expected {
		got := r.Uint64()
		if got != want {
			t.Error("incorrect random number sequence, got:", got, "want:", want)
		}
	}

}

func TestXoshiro256StarStarRandSource(t *testing.T) {
	var r Xoshiro256StarStar
	r.Seed(42)
	for i := 0; i < 10000; i++ {
		got := r.Int63()
		if got < 0 {
			t.Error("expected positive number, got:", got)
		}
	}
}
