package main

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {

	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	expected := []bool{true, true, true, false, true}
	got := kidsWithCandies(candies, extraCandies)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}

	candies = []int{4, 2, 1, 1, 2}
	extraCandies = 1
	expected = []bool{true, false, false, false, false}
	got = kidsWithCandies(candies, extraCandies)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}

}
