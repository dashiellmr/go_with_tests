package arrays

import (
	"reflect"
	"slices"
	"testing"
)

func TestArrays(t *testing.T) {
	t.Run("testing with 5 values", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		got := Summation(nums)
		want := 15
		assertSummationEquality(t, got, want, nums)
	})
	t.Run("testing with less than 5 values", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got := Summation(nums)
		want := 6
		assertSummationEquality(t, got, want, nums)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("testing summation of two arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{6, 15}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("testing functionality of sumAllTails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}
		if !slices.Equal(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("testing summation of two empty arrays", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{})
		want := []int{0, 0}
		if !slices.Equal(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func assertSummationEquality(t testing.TB, got, want int, given []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q, given %v", got, want, given)
	}
}
