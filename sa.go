// Copyright 2013 Robert Hülle.
// Use of this source code is governed by the ISC license
// which can be found in the LICENSE file.

package bbwt

func BuildSA(str []byte) []int {
	a := NewSuffixArray(str)
	a.BuildSA()
	return a.SA
}

type SuffixArray struct {
	str  []byte
	SA   []int
	rank []int
	step int
}

func NewSuffixArray(str []byte) *SuffixArray {
	s := &SuffixArray{str: str}
	if len(str) > 255 {
	}
	return s
}

func (s *SuffixArray) BuildSA() {
	if s.str == nil {
		return
	}
	s.SA = make([]int, len(s.str))
	for i := range s.SA {
		s.SA[i] = i
	}
	s.rankInit()
	for s.step = 1; ; s.step *= 2 {
		s.bucketSort()
		lex, end := s.lexRank()
		if end {
			break
		}
		for j, l := range lex {
			s.rank[s.SA[j]] = l
		}
	}
	s.rank = nil
}

func (s *SuffixArray) rankInit() {
	s.rank = make([]int, len(s.str))
	if len(s.str) > 255 {
		for i, b := range s.str {
			s.rank[i] = int(b)
		}
		return
	}
	cnt := make([]int, 256)
	for _, b := range s.str {
		cnt[b]++
	}
	lexName := 0
	for i, c := range cnt {
		if c == 0 {
			continue
		}
		cnt[i] = lexName
		lexName++
	}
	for i, b := range s.str {
		s.rank[i] = cnt[b]
	}
}

func (s *SuffixArray) bucketKey(i, pass int) int {
	i += s.step * pass
	if i < len(s.rank) {
		return 1 + s.rank[i]
	} else {
		return 0
	}
}

func (s *SuffixArray) bucketPass(dst, src []int, pass int) {
	cnt := make([]int, len(src)+1)
	for _, n := range src {
		key := s.bucketKey(n, pass)
		// FIXME: rank is too big for small strings
		cnt[key]++
	}
	cumCnt := make([]int, len(cnt))
	for i := 1; i < len(cumCnt); i++ {
		cumCnt[i] = cumCnt[i-1] + cnt[i-1]
	}
	for _, n := range src {
		key := s.bucketKey(n, pass)
		dst[cumCnt[key]] = n
		cumCnt[key]++
	}
}

func (s *SuffixArray) bucketSort() {
	tmp := make([]int, len(s.SA))
	s.bucketPass(tmp, s.SA, 1)
	s.bucketPass(s.SA, tmp, 0)
}

func (s *SuffixArray) less(a, b int) bool {
	if s.rank[a] != s.rank[b] {
		return s.rank[a] < s.rank[b]
	}
	a += s.step
	b += s.step
	if n := len(s.rank); a < n && b < n {
		return s.rank[a] < s.rank[b]
	} else {
		return a > b
	}
}

func (s *SuffixArray) lexRank() (lex []int, end bool) {
	lex = make([]int, len(s.rank))
	crank := 0
	for i := 1; i < len(s.rank); i++ {
		if s.less(s.SA[i-1], s.SA[i]) {
			crank++
		}
		lex[i] = crank
	}
	n1 := len(lex) - 1
	end = lex[n1] == n1
	return
}
