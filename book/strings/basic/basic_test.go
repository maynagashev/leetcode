package basic

import (
	"bytes"
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {

}

/*
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkJoin-12                 2400486               478.8 ns/op          2688 B/op          1 allocs/op
BenchmarkJoinConcat-12            385388              2649 ns/op           15840 B/op         12 allocs/op
BenchmarkConcat-12               1000000             60708 ns/op          503994 B/op          1 allocs/op
BenchmarkBuffer-12              126828178                8.248 ns/op           2 B/op          0 allocs/op
BenchmarkCopy-12                257498829                4.516 ns/op           0 B/op          0 allocs/op
BenchmarkStringBuilder-12       499872146                3.216 ns/op           5 B/op          0 allocs/op

*/

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		join(song)
	}
}

func BenchmarkJoinConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		joinConcat(song)
	}
}

func BenchmarkConcat(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str += "x"
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); str != s {
		b.Errorf("unexpected result; got=%s, want=%s", str, s)
	}
}

func BenchmarkBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); buffer.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
	}
}

func BenchmarkCopy(b *testing.B) {
	bs := make([]byte, b.N)
	bl := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bl += copy(bs[bl:], "x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); string(bs) != s {
		b.Errorf("unexpected result; got=%s, want=%s", string(bs), s)
	}
}

// Go 1.10
func BenchmarkStringBuilder(b *testing.B) {
	// ZERO-VALUE:
	//
	// It's ready to use from the get-go.
	// You don't need to initialize it.
	var strBuilder strings.Builder

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		strBuilder.WriteString("x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); strBuilder.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", strBuilder.String(), s)
	}
}
