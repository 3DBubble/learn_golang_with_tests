package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	checkGotEqualsWant := func(t testing.TB, got, want int, nums []int) {
		t.Helper()
		if got != want {
			t.Errorf("given %v, got %d, want %d", nums, got, want)
		}
	}

	t.Run("collection of 5 integers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		got := Sum(nums)
		want := 15
		checkGotEqualsWant(t, got, want, nums)
	})

	t.Run("collection of vairable integers", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got := Sum(nums)
		want := 6
		checkGotEqualsWant(t, got, want, nums)
	})
}

func TestSumAllTails(t *testing.T) {
	checkDeepEqual := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("Multiple Slices of integers", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}
		checkDeepEqual(t, got, want)
	})

	t.Run("Empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 5, 6})
		want := []int{0, 11}
		checkDeepEqual(t, got, want)
	})
}
