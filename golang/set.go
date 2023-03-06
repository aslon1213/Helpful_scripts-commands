// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package main

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	size  int // bitsize
	words []uint64
}

func main() {
	a := IntSet{}
	a.Add(1)
	a.Add(144)
	a.Add(9)
	fmt.Println(a.String()) // "{1 9 144}"
	a.Remove(1)
	fmt.Println(a.String()) // "{9 144}"
	b := a.Copy()
	fmt.Println(b.String()) // "{9 144}"
	a.Clear()
	fmt.Println(a.String()) // "{}"
	b.Clear()

	a.addAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	b.addAll(5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	fmt.Println(a.String()) // "{1 2 3 4 5 6 7 8 9 10}"
	fmt.Println(b.String()) // "{5 6 7 8 9 10 11 12 13 14 15}"
	a.DifferenceWith(b)
	fmt.Println(a.String()) // "{1 2 3 4}"
	a.Clear()
	a.addAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	a.intersectWith(b)
	fmt.Println(a.String()) // "{5 6 7 8 9 10}"
	a.Clear()
	a.addAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	a.SymmetricDifference(b)
	fmt.Println(a.String()) // "{1 2 3 4 11 12 13 14 15}"
	fmt.Println(a.Elems())
}
func (s *IntSet) Elems() []int64 {
	var elems []int64

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, int64(64*i+j))
			}
		}
	}

	return elems
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
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

func (s *IntSet) Len() int {

	return len(s.words)
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)

	if word < len(s.words) && s.words[word]&(1<<bit) != 0 {
		s.words[word] &= ^(1 << bit)
	}

}

func (s *IntSet) addAll(x ...int) {
	for _, v := range x {
		s.Add(v)
	}
}

func (s *IntSet) intersectWith(t *IntSet) {
	if len(s.words) > len(t.words) {
		for i, _ := range s.words {
			if i < len(t.words) {
				s.words[i] &= t.words[i]
			} else {
				s.words[i] = 0
			}
		}
	} else {
		for i, _ := range s.words {
			s.words[i] &= t.words[i]
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	if len(s.words) > len(t.words) {
		for i, _ := range s.words {
			if i < len(t.words) {
				s.words[i] &= ^t.words[i]
			}
		}
	} else {
		for i, _ := range s.words {
			s.words[i] &= ^t.words[i]
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	if len(s.words) > len(t.words) {
		for i, _ := range s.words {
			if i < len(t.words) {
				s.words[i] ^= t.words[i]
			}
		}
	} else {
		for i, _ := range s.words {
			s.words[i] ^= t.words[i]
		}
	}
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	var copy IntSet
	copy.words = make([]uint64, len(s.words))
	copy.words = s.words
	return &copy
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
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string
