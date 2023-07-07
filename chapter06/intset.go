package chapter06

import (
	"bytes"
	"fmt"
)

// Contains solutions of Exercises 6.1 to 6.4
type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Exercise 6.2 solution
func (s *IntSet) AddAll(nums ...int) {
	for _, num := range nums {
		s.Add(num)
	}
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Exercise 6.4 solution
func (s *IntSet) Elems() []int {
	var slice []int
	for i, word := range s.words {
		if word != 0 {
			for j := 0; j < 64; j++ {
				if word&(1<<uint(j)) != 0 {
					slice = append(slice, 64*i+j)
				}
			}
		}
	}
	return slice
}

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

// Exercise 6.1 solution
func (s *IntSet) Len() int {
	var len int
	for _, word := range s.words {
		if word != 0 {
			for j := 0; j < 64; j++ {
				if word&(1<<uint(j)) != 0 {
					len++
				}
			}
		}
	}
	return len
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &= ^(1 << bit)
	}
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	var copy []uint64
	for _, word := range s.words {
		copy = append(copy, word)
	}
	return &IntSet{
		words: copy,
	}
}
