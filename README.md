[![GoDoc](https://godoc.org/github.com/kbjorklu/xoshiro?status.svg)](https://godoc.org/github.com/kbjorklu/xoshiro)
[![Go Report Card](https://goreportcard.com/badge/github.com/kbjorklu/xoshiro)](https://goreportcard.com/report/github.com/kbjorklu/xoshiro)

## Overview

Package xoshiro implements the Xoshiro256** and SplitMix64 pseudorandom number
generators described at http://xoshiro.di.unimi.it/ . The generators implement
[rand.Source](https://golang.org/pkg/math/rand/#Source) and
[rand.Source64](https://golang.org/pkg/math/rand/#Source64).
