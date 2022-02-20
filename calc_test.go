package main

import (
	"strconv"
	"testing"
)

func TestCalc(t *testing.T) {
	{
		sum := calc([]int{})
		if sum != 0 {
			t.Fatal("expected 0, but was " + strconv.Itoa(sum))
		}
	}
	{
		sum := calc([]int{1})
		if sum != 1 {
			t.Fatal("expected 1, but was " + strconv.Itoa(sum))
		}
	}
	{
		sum := calc([]int{1, 2, 3})
		if sum != 6 {
			t.Fatal("expected 6, but was " + strconv.Itoa(sum))
		}
	}
}
