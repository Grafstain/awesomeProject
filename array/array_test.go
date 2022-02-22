package array

import (
	"reflect"
	"testing"
)

func TestArraySum(t *testing.T) {

	t.Run("Суммирование элементов массива", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := ArraySum(numbers)
		want := 15
		if got != want {
			t.Errorf("given %v want %d got %d", numbers, want, got)
		}
	})
	t.Run("Суммирование всех элементов и вывод в новом слайсе", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9}, []int{5, 9})
		want := []int{3, 9, 14}
		checkSums(t, got, want)
	})
	t.Run("Суммирование хвостов всех срезов", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 1}, []int{0, 9, 6}, []int{5, 9, 2})
		want := []int{3, 15, 11}
		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func checkSums(t *testing.T, got []int, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v got %v", want, got)
	}
}
