package union_find_test

import (
	. "github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/union-find"
	"math/rand"
	"testing"
	"time"
)

func TestUnionFind_Union(t *testing.T) {
	n := 1 << 10
	u := New(n)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		u.Union(a, b)

		if !u.IsConnected(a, b) {
			t.Errorf("Union %d, %d", a, b)
		}
	}
}

func TestUnionFind_UnionS(t *testing.T) {
	n := 1 << 10
	u := New(n)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		u.UnionS(a, b)

		if !u.IsConnected(a, b) {
			t.Errorf("UnionS %d, %d", a, b)
		}
	}
}

func TestUnionFind_UnionR(t *testing.T) {
	n := 1 << 10
	u := New(n)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		u.UnionR(a, b)

		if !u.IsConnected(a, b) {
			t.Errorf("UnionR %d, %d", a, b)
		}
	}
}

func BenchmarkUnionFind_Union1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		n := 1 << 10
		u := New(n)
		rand.Seed(time.Now().UnixNano())
		b.StartTimer()
		for i := 0; i < n; i++ {
			a := rand.Int() % n
			b := rand.Int() % n
			u.Union(a, b)
		}

		for i := 0; i < n; i++ {
			a := rand.Int() % n
			b := rand.Int() % n
			u.IsConnected(a, b)
		}

		b.StopTimer()
	}
}

func BenchmarkUnionFind_UnionR1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		n := 1 << 10
		u := New(n)
		rand.Seed(time.Now().UnixNano())
		b.StartTimer()
		for i := 0; i < n; i++ {
			a := rand.Int() % n
			b := rand.Int() % n
			u.UnionR(a, b)
		}

		for i := 0; i < n; i++ {
			a := rand.Int() % n
			b := rand.Int() % n
			u.IsConnected(a, b)
		}

		b.StopTimer()
	}
}

func BenchmarkUnionFind_UnionS1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		n := 1 << 10
		u := New(n)
		rand.Seed(time.Now().UnixNano())
		b.StartTimer()
		for i := 0; i < n; i++ {
			a := rand.Int() % n
			b := rand.Int() % n
			u.UnionS(a, b)
		}

		for i := 0; i < n; i++ {
			a := rand.Int() % n
			b := rand.Int() % n
			u.IsConnected(a, b)
		}

		b.StopTimer()
	}
}
