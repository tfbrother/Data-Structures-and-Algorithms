package sorting_test

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	. "github.com/tfbrother/Data-Structures-and-Algorithms/algorithms/sorting"
	"testing"
)

//func BenchmarkSortInt1K(b *testing.B) {
//	b.StopTimer()
//	for i := 0; i < b.N; i++ {
//		data := make([]int, 1<<10)
//		for i := 0; i < len(data); i++ {
//			data[i] = i ^ 0x2cc
//		}
//		b.StartTimer()
//		QuickSort4(data)
//		b.StopTimer()
//	}
//}
//
//func BenchmarkSortInt64K(b *testing.B) {
//	b.StopTimer()
//	for i := 0; i < b.N; i++ {
//		data := make([]int, 1<<16)
//		for i := 0; i < len(data); i++ {
//			data[i] = i ^ 0xcccc
//		}
//		b.StartTimer()
//		QuickSort4(data)
//		b.StopTimer()
//	}
//}

func bench(b *testing.B, size int, algo func([]int), name string) {
	b.StopTimer()
	data := make([]int, size)
	x := ^uint32(0)
	for i := 0; i < b.N; i++ {
		for n := size - 3; n <= size+3; n++ {
			for i := 0; i < len(data); i++ {
				x += x
				x ^= 1
				if int32(x) < 0 {
					x ^= 0x88888eef
				}
				data[i] = int(x % uint32(n/5))
			}
			b.StartTimer()
			algo(data)
			b.StopTimer()
			if !IsSorted(data) {
				b.Errorf("%s did not sort %d ints", name, n)
			}
		}
	}
}

func BenchmarkSort51e2(b *testing.B) { bench(b, 1e2, QuickSort5, "QuickSort5") }

func BenchmarkSort41e2(b *testing.B) { bench(b, 1e2, QuickSort4, "QuickSort4") }

//func BenchmarkSort41e4(b *testing.B) { bench(b, 1e4, QuickSort4, "QuickSort4") }
//func BenchmarkSort41e6(b *testing.B) { bench(b, 1e6, QuickSort4, "QuickSort4") }
func BenchmarkSort31e2(b *testing.B) { bench(b, 1e2, QuickSort3, "QuickSort3") }
func BenchmarkSort21e2(b *testing.B) { bench(b, 1e2, QuickSort2, "QuickSort2") }
func BenchmarkSort11e2(b *testing.B) { bench(b, 1e2, QuickSort1, "QuickSort1") }
func BenchmarkSort1e2(b *testing.B)  { bench(b, 1e2, QuickSort, "QuickSort") }

//func BenchmarkSort31e4(b *testing.B) { bench(b, 1e4, QuickSort3, "QuickSort3") }
//func BenchmarkSort31e6(b *testing.B) { bench(b, 1e6, QuickSort3, "QuickSort3") }
