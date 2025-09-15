package sorting

import (
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("BubbleSort => [4, 1, 2, 5, 3] -> [1, 2, 3, 4, 5]", func(t *testing.T) {
		arr := []int{4, 1, 2, 5, 3}

		want := []int{1, 2, 3, 4, 5}
		if got := BubbleSort(arr); !slices.Equal(got, want) {
			t.Errorf("BubbleSort failed, expected %v but got %v", want, got)
		}
	})
	t.Run("BubbleSort => [5, 4, 3, 2, 1] -> [1, 2, 3, 4, 5]", func(t *testing.T) {
		arr := []int{5, 4, 3, 2, 1}

		want := []int{1, 2, 3, 4, 5}
		if got := BubbleSort(arr); !slices.Equal(got, want) {
			t.Errorf("BubbleSort failed, expected %v but got %v", want, got)
		}
	})
	t.Run("BubbleSort => [1, 2, 3, 4, 5] -> [1, 2, 3, 4, 5]", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}

		want := []int{1, 2, 3, 4, 5}
		if got := BubbleSort(arr); !slices.Equal(got, want) {
			t.Errorf("BubbleSort failed, expected %v but got %v", want, got)
		}
	})
}

func TestInsertionSort(t *testing.T) {
	t.Run("InsertionSort => [4, 1, 2, 5, 3] -> [1, 2, 3, 4, 5]", func(t *testing.T) {
		arr := []int{4, 1, 2, 5, 3}

		want := []int{1, 2, 3, 4, 5}
		if got := BubbleSort(arr); !slices.Equal(got, want) {
			t.Errorf("InsertionSort failed, expected %v but got %v", want, got)
		}
	})
	t.Run("InsertionSort => [5, 4, 3, 2, 1] -> [1, 2, 3, 4, 5]", func(t *testing.T) {
		arr := []int{5, 4, 3, 2, 1}

		want := []int{1, 2, 3, 4, 5}
		if got := InsertionSort(arr); !slices.Equal(got, want) {
			t.Errorf("InsertionSort failed, expected %v but got %v", want, got)
		}
	})
	t.Run("InsertionSort => [1, 2, 3, 4, 5] -> [1, 2, 3, 4, 5]", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}

		want := []int{1, 2, 3, 4, 5}
		if got := InsertionSort(arr); !slices.Equal(got, want) {
			t.Errorf("InsertionSort failed, expected %v but got %v", want, got)
		}
	})
}
