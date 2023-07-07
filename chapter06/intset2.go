package chapter06

import (
	"bytes"
	"fmt"
)

const intSize = 32 << (^uint(0) >> 63)

// Contains solutions of Exercises 6.1 to 6.4
type IntSet2 struct {
	words []uint
}

func (s *IntSet2) Has(x int) bool {
	word, bit := x/intSize, uint(x%intSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet2) Add(x int) {
	word, bit := x/intSize, uint(x%intSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Exercise 6.2 solution
func (s *IntSet2) AddAll(nums ...int) {
	for _, num := range nums {
		s.Add(num)
	}
}

func (s *IntSet2) UnionWith(t *IntSet2) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Exercise 6.4 solution
func (s *IntSet2) Elems() []int {
	var slice []int
	for i, word := range s.words {
		if word != 0 {
			for j := 0; j < intSize; j++ {
				if word&(1<<uint(j)) != 0 {
					slice = append(slice, intSize*i+j)
				}
			}
		}
	}
	return slice
}

func (s *IntSet2) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", intSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Exercise 6.1 solution
func (s *IntSet2) Len() int {
	var len int
	for _, word := range s.words {
		if word != 0 {
			for j := 0; j < intSize; j++ {
				if word&(1<<uint(j)) != 0 {
					len++
				}
			}
		}
	}
	return len
}

func (s *IntSet2) Remove(x int) {
	word, bit := x/intSize, uint(x%intSize)
	if word < len(s.words) {
		s.words[word] &= ^(1 << bit)
	}
}

func (s *IntSet2) Clear() {
	s.words = []uint{}
}

func (s *IntSet2) Copy() *IntSet2 {
	var copy []uint
	for _, word := range s.words {
		copy = append(copy, word)
	}
	return &IntSet2{
		words: copy,
	}
}
