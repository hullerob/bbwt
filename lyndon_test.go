// Copyright 2013 Robert Hülle.
// Use of this source code is governed by the ISC license
// which can be found in the LICENSE file.

package bbwt

import (
	"bytes"
	"testing"
)

func TestFactorOneChar(t *testing.T) {
	in := []byte("a")
	out := Factor(in)
	if out == nil {
		t.Fatal("Slice of Lyndon words is nil.")
	}
	if len(out) != 1 {
		t.Fatalf("Bad number of Lyndon words. expected: %d, got: %d.", 1, len(out))
	}
	if out[0] == nil {
		t.Fatal("Lyndon word is nil.")
	}
	if len(out[0]) != 1 {
		t.Fatalf("Bad length of Lyndon word. expected: %d, got: %d.", 1, len(out[0]))
	}
	if out[0][0] != 'a' {
		t.Fatal("Bad character in Lyndon word.")
	}
}

func TestFactorUniformString(t *testing.T) {
	in := []byte("aaaaaaa")
	out := Factor(in)
	if out == nil {
		t.Fatal("Slice of Lyndon words is nil.")
	}
	if len(out) != len(in) {
		t.Fatalf("Bad number of Lyndon words. expected: %d, got: %d.", len(in), len(out))
	}
	for i, w := range out {
		if w == nil {
			t.Fatalf("Lyndon word #%d is nil.", i)
		}
		if len(w) != 1 {
			t.Fatalf("Lyndon word #%d has bad length. expected: %d, got: %d.", i, 1, len(w))
		}
		if w[0] != 'a' {
			t.Fatal("Bad character in Lyndon word.")
		}
	}
}

func testFactorSimple(t *testing.T, in []byte, expected [][]byte) {
	out := Factor(in)
	if out == nil {
		t.Fatal("Slice of Lyndon words is nil")
	}
	if len(out) != len(expected) {
		t.Fatalf("Bad number of Lyndon words. expected: %d, got: %d.", len(expected), len(out))
	}
	for i, we := range expected {
		if len(we) != len(out[i]) {
			t.Fatalf("Bad length of Lyndon word #%d. expected: %d, got: %d.", i, len(we), len(out[i]))
		}
		if bytes.Compare(we, out[i]) != 0 {
			t.Fatalf("Bad Lyndon word #%d.", i)
		}
	}
}

func TestFactor1(t *testing.T) {
	in := []byte("swiss miss")
	e := make([][]byte, 0, 3)
	e = append(e, []byte("sw"), []byte("iss"), []byte(" miss"))
	testFactorSimple(t, in, e)
}
