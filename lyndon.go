// Copyright 2013 Robert Hülle.
// Use of this source code is governed by the ISC license
// which can be found in the LICENSE file.

package bbwt

// Factor returns Lyndon factorization of data.
func Factor(data []byte) (words [][]byte) {
	for len(data) > 0 {
		a := 0
		b := 1
		for b < len(data) && data[a] <= data[b] {
			if data[a] == data[b] {
				a++
			} else {
				a = 0
			}
			b++
		}
		l := b - a
		for a >= 0 {
			words = append(words, data[:l])
			data = data[l:]
			a -= l
		}
	}
	return
}
