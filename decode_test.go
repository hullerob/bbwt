// Copyright 2013 Robert Hülle.
// Use of this source code is governed by the ISC license
// which can be found in the LICENSE file.

package bbwt

import (
	"bytes"
	"math/rand"
	"testing"
)

func testDecodeSimple(t *testing.T, in, exp []byte) {
	out := Decode(in)
	if out == nil {
		t.Fatal("Decode returned nil.")
	}
	if len(out) != len(exp) {
		t.Fatalf("Decoded string has bad length. expected: %d, got: %d.", len(exp), len(out))
	}
	if bytes.Compare(out, exp) != 0 {
		t.Fatal("Decoded string differs from expectation.")
	}
}

func TestDecodeTrivial(t *testing.T) {
	in := []byte("aaaaaa")
	testDecodeSimple(t, in, in)
}

func TestDecodeSimple(t *testing.T) {
	in := []byte("sms ssiiws")
	exp := []byte("swiss miss")
	testDecodeSimple(t, in, exp)
}

func benchmarkDecode(b *testing.B, size int) {
	var out []byte
	r := rand.New(rand.NewSource(3001))
	in := make([]byte, size)
	for i := range in {
		in[i] = byte(r.Intn(256))
	}
	enc := Encode(in)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = Decode(enc)
	}
	out = out
}

func BenchmarkDecode100B(b *testing.B) { benchmarkDecode(b, 100) }
func BenchmarkDecode1k(b *testing.B)   { benchmarkDecode(b, 1000) }
func BenchmarkDecode10k(b *testing.B)  { benchmarkDecode(b, 10000) }
func BenchmarkDecode100k(b *testing.B) { benchmarkDecode(b, 100000) }
func BenchmarkDecode200k(b *testing.B) { benchmarkDecode(b, 200000) }
func BenchmarkDecode300k(b *testing.B) { benchmarkDecode(b, 300000) }
func BenchmarkDecode400k(b *testing.B) { benchmarkDecode(b, 400000) }
func BenchmarkDecode500k(b *testing.B) { benchmarkDecode(b, 500000) }
func BenchmarkDecode1M(b *testing.B)   { benchmarkDecode(b, 1000000) }
