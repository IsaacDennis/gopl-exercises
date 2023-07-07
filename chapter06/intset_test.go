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
