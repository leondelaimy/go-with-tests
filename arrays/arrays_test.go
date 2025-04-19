package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum slice of any size", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}

		got := Sum(slice)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, slice)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum all slices of any size", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		slice2 := []int{1, 2}

		got := SumAll(slice, slice2)
		want := []int{15, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}