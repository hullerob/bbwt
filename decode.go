// Copyright 2013 Robert Hülle.
// Use of this source code is governed by the ISC license
// which can be found in the LICENSE file.

package bbwt

// Decode implements inverse transformation to Bijective BWT.
func Decode(data []byte) (out []byte) {
	t := constructT(data)
	out = make([]byte, len(data))
	i := len(data) - 1
	for j := range data {
		if t[j] == -1 {
			continue
		}
		k := j
		for t[k] != -1 {
			out[i] = data[k]
			i--
			k, t[k] = t[k], -1
		}
	}
	return
}

func constructT(data []byte) []int {
	t := make([]int, len(data))
	var counts [256]int
	for _, b := range data {
		counts[b]++
	}
	var cumCounts [256]int
	for i := 1; i < 256; i++ {
		cumCounts[i] = cumCounts[i-1] + counts[i-1]
	}
	for i, b := range data {
		t[i] = cumCounts[b]
		cumCounts[b]++
	}
	return t
}
