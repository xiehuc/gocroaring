package gocroaring_test

import (
	"math/rand"
	"testing"

	"github.com/xiehuc/gocroaring"
)

var ordered []uint32
var random []uint32

func init() {
	var i uint32
	for i = 0; i < 50000; i++ {
		ordered = append(ordered, i)
		random = append(random, uint32(rand.Int31n(1e6)/200))
	}
}

func benchmarkAdd(b *testing.B, sl []uint32) {
	for n := 0; n < b.N; n++ {
		rb1 := gocroaring.New()
		for _, i := range sl {
			rb1.Add(i)
		}
	}
}

func benchmarkAddMany(b *testing.B, sl []uint32) {
	for n := 0; n < b.N; n++ {
		rb1 := gocroaring.New()
		rb1.Add(sl...)
	}
}

func benchmarkNewFromPtr(b *testing.B, sl []uint32) {
	for n := 0; n < b.N; n++ {
		rb := gocroaring.New(sl...)
		_ = rb
	}

}

func BenchmarkAddRandom(b *testing.B)  { benchmarkAdd(b, random) }
func BenchmarkAddOrdered(b *testing.B) { benchmarkAdd(b, ordered) }

func BenchmarkAddRandomArity(b *testing.B)  { benchmarkAddMany(b, random) }
func BenchmarkAddOrderedArity(b *testing.B) { benchmarkAddMany(b, ordered) }

func BenchmarkRandomNewFromPtr(b *testing.B)  { benchmarkNewFromPtr(b, random) }
func BenchmarkOrderedNewFromPtr(b *testing.B) { benchmarkNewFromPtr(b, ordered) }
