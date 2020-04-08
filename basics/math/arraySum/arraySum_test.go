package arraySum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := ArraySum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := ArraySum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumNArrays([]int{1, 2, 3}, []int{0, 9})
	want := []int{6, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	t.Run("collection of 2 arrays", func(t *testing.T) {
		got := SumArrayTails([]int{1, 2, 3}, []int{0, 9})
		want := []int{5, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("collection of arrays including an empty one", func(t *testing.T) {
		got := SumArrayTails([]int{1, 2, 3}, []int{0, 9}, []int{})
		want := []int{5, 9, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func TestSumArrayValuesAtPosition(t *testing.T) {

	t.Run("collection of different sized arrays", func(t *testing.T) {
		got := SumArrayValuesAtPosition([]int{1, 2, 3}, []int{0, 9})
		want := []int{1, 11, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("collection of different sized arrays including an empty one", func(t *testing.T) {
		got := SumArrayValuesAtPosition([]int{1, 2, 3}, []int{0, 9}, []int{})
		want := []int{1, 11, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
