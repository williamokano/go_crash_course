package fib

import (
	"fmt"
	"math/big"
	"testing"
)

func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(10)
	}
}

func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(20)
	}
}

func BenchmarkFib100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(100)
	}
}

func bigInt(value string) *big.Int {
	n := new(big.Int)
	n, ok := n.SetString(value, 10)
	if !ok {
		panic("The hell with this")
	}

	return n
}

func TestFib(t *testing.T) {
	data := []struct {
		n    uint
		want *big.Int
	}{
		{0, bigInt("0")},
		{1, bigInt("1")},
		{2, bigInt("1")},
		{3, bigInt("2")},
		{4, bigInt("3")},
		{5, bigInt("5")},
		{6, bigInt("8")},
		{10, bigInt("55")},
		{42, bigInt("267914296")},
		{100, bigInt("354224848179261915075")},
		{200, bigInt("280571172992510140037611932413038677189525")},
		{300, bigInt("222232244629420445529739893461909967206666939096499764990979600")},
		{400, bigInt("176023680645013966468226945392411250770384383304492191886725992896575345044216019675")},
	}

	for _, d := range data {
		t.Run(fmt.Sprintf("Testing fib(%d)", d.n), func(t *testing.T) {
			if got := Fib(d.n); got.Cmp(d.want) != 0 {
				t.Errorf("Invalid Fibonacci value for N: %d, got: %s, want: %s", d.n, got.String(), d.want.String())
			}
		})
	}
}
