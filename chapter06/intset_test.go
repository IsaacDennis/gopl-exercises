package chapter06

import "testing"

func TestIntSetHas(t *testing.T) {
	var input []int = []int{2, 4, 128, 5}
	set := IntSet{}
	for _, num := range input {
		set.Add(num)
	}
	for _, num := range input {
		if !set.Has(num) {
			t.Errorf("Error: expected %d to be contained by set.", num)
		}
	}
}

func TestIntSetAddAll(t *testing.T) {
	var input []int = []int{2, 4, 128, 5}
	set := IntSet{}
	set.AddAll(2, 4, 128, 5)
	for _, num := range input {
		if !set.Has(num) {
			t.Errorf("Error: expected %d to be contained by set.", num)
		}
	}
}

func TestIntSetElems(t *testing.T) {
	var expected = []int{1, 2, 3, 4, 5, 6, 10}
	set := IntSet{}
	set.AddAll(expected...)
	elems := set.Elems()
	if !CmpSlices(elems, expected) {
		t.Errorf("Error: expected %v, got %v", expected, elems)
	}
}

func TestIntSetLen(t *testing.T) {
	var input = []int{1, 2, 3, 5, 123}
	var expected int = 5
	set := IntSet{}
	set.AddAll(input...)
	len := set.Len()
	if len != expected {
		t.Errorf("Error: expected %d, got %d", expected, len)
	}
}

func TestIntSetRemove(t *testing.T) {
	var input = []int{1, 2, 3, 4, 5, 6}
	set := IntSet{}
	set.AddAll(input...)
	set.Remove(2)
	if set.Has(2) {
		t.Errorf("Error: expected %d to be removed", 2)
	}
}

// True if slices are equal (including ints being in the same order)
func CmpSlices(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	} else {
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				return false
			}
		}
		return true
	}
}
