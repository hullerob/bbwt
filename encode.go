// Copyright 2013 Robert Hülle.
// Use of this source code is governed by the ISC license
// which can be found in the LICENSE file.

package bbwt

type rot struct {
	// index of Lyndon word
	w int
	// rotation of Lyndon word
	r int
}

// Encode implements Bijective Burrows-Wheeler Transform.
func Encode(data []byte) (out []byte) {
	factors := Factor(data)
	rots := make([][]rot, 0, len(factors))
	for i, w := range factors {
		r := make([]rot, len(w))
		for j := range w {
			r[j] = rot{i, j}
		}
		rots = append(rots, r)
	}
	sRot := sortRot(factors, rots)
	out = make([]byte, len(data))
	for i, r := range sRot {
		ri := r.r - 1
		if ri < 0 {
			ri += len(factors[r.w])
		}
		out[i] = factors[r.w][ri]
	}
	return
}

func lessRot(factors [][]byte, i, j rot) bool {
	wi := i.w
	ri := i.r
	li := len(factors[wi])
	wj := j.w
	rj := j.r
	lj := len(factors[wj])
	for k := 0; k < li*lj; k++ {
		if factors[wi][ri] < factors[wj][rj] {
			return true
		} else if factors[wi][ri] > factors[wj][rj] {
			return false
		}
		ri++
		rj++
		if ri == li {
			ri = 0
		}
		if rj == lj {
			rj = 0
		}
	}
	return false
}

func sortFactor(factors [][]byte, rots []rot) {
	sa := BuildSA(factors[rots[0].w])
	for i, s := range sa {
		rots[i].r = s
	}
}

func sortRot(factors [][]byte, rots [][]rot) []rot {
	wordCount := 0
	for _, r := range rots {
		wordCount++
		sortFactor(factors, r)
	}
	return mergeRots(factors, rots)
}

func mergeRots(factors [][]byte, rots [][]rot) []rot {
	var merged []rot
	for len(rots) > 0 {
		merged = mergeRot2(factors, merged, rots[0])
		rots = rots[1:]
	}
	return merged
}

func mergeRot2(factors [][]byte, a, b []rot) []rot {
	la := len(a)
	lb := len(b)
	length := la + lb
	out := make([]rot, length)
	i := 0
	j := 0
	k := 0
	for i < la && j < lb {
		if lessRot(factors, b[j], a[i]) {
			out[k] = b[j]
			j++
		} else {
			out[k] = a[i]
			i++
		}
		k++
	}
	if i < la {
		copy(out[k:], a[i:])
	} else if j < lb {
		copy(out[k:], b[j:])
	}
	return out
}
