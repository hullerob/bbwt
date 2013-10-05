// Copyright 2013 Robert Hülle.
// Use of this source code is governed by the ISC license
// which can be found in the LICENSE file.

package bbwt

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestEncodeTrivial(t *testing.T) {
	in := []byte("aaaaa")
	out := Encode(in)
	if out == nil {
		t.Fatal("Encode returned nil.")
	}
	if len(out) != len(in) {
		t.Fatalf("Encoded string has different length. expected: %d, got: %d.", len(in), len(out))
	}
	if bytes.Compare(in, out) != 0 {
		t.Fatal("Encoded string differs from expectation.")
	}
}

func testEncodeSimple(t *testing.T, in, exp []byte) {
	out := Encode(in)
	if out == nil {
		t.Error("Encode returned nil.")
		return
	}
	if len(out) != len(exp) {
		t.Errorf("Encoded string has different length. expected: %d, got:%d.", len(exp), len(out))
		return
	}
	if bytes.Compare(out, exp) != 0 {
		t.Error("Encoded string differs from expectation.")
	}
}

func TestEncode1(t *testing.T) {
	testEncodeSimple(t, []byte("swiss miss"), []byte("sms ssiiws"))
}

func benchmarkEncode(b *testing.B, size int) {
	var out []byte
	r := rand.New(rand.NewSource(3001))
	in := make([]byte, size)
	for i := range in {
		in[i] = byte(r.Intn(256))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = Encode(in)
	}
	out = out
}

func BenchmarkEncode100B(b *testing.B) { benchmarkEncode(b, 100) }
func BenchmarkEncode1K(b *testing.B)   { benchmarkEncode(b, 1000) }
func BenchmarkEncode10K(b *testing.B)  { benchmarkEncode(b, 10000) }
func BenchmarkEncode100K(b *testing.B) { benchmarkEncode(b, 100000) }
func BenchmarkEncode200K(b *testing.B) { benchmarkEncode(b, 200000) }
func BenchmarkEncode300K(b *testing.B) { benchmarkEncode(b, 300000) }
func BenchmarkEncode400K(b *testing.B) { benchmarkEncode(b, 400000) }
func BenchmarkEncode500K(b *testing.B) { benchmarkEncode(b, 500000) }
func BenchmarkEncode1M(b *testing.B)   { benchmarkEncode(b, 1000000) }
