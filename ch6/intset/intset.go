// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
const bit = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/bit, uint(x%bit)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/bit, uint(x%bit)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersect of s and t.
func (s *IntSet) IntersectWith(t *IntSet) *IntSet {
	var z IntSet
	for i, tword := range t.words {
		if i < len(s.words) {
			z.words = append(z.words, s.words[i]&tword)
		}
	}
	return &z
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
	var z IntSet
	for i, tword := range t.words {
		if i < len(s.words) {
			tmp := s.words[i] & tword
			z.words = append(z.words, s.words[i]-tmp)
		}
	}
	return &z
}

// SymmetricDifference sets s to the difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	var z IntSet
	if len(t.words) == 0 {
		return s
	}
	len1, len2 := len(s.words), len(t.words)
	n := len1
	if len1 < len2 {
		n = len2
	}

	for i := 0; i < n; i++ {
		if i < len1 && i < len2 {
			tmp := s.words[i] & t.words[i]
			z.words = append(z.words, (s.words[i]-tmp)|(t.words[i]-tmp))
		} else if i < len1 {
			z.words = append(z.words, s.words[i])
		} else {
			z.words = append(z.words, t.words[i])
		}
	}
	return &z
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bit; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bit*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

// return the number of elements
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bit; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/bit, uint(x%bit)
	for word >= len(s.words) {
		return
	}
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

func (s *IntSet) Copy() *IntSet {
	var b IntSet
	for _, word := range s.words {
		b.words = append(b.words, word)
	}
	return &b
}

func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) Elems() []uint {
	var ans []uint
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bit; j++ {
			if word&(1<<uint(j)) != 0 {
				ans = append(ans, uint(bit*i+j))
			}
		}
	}
	return ans
}
