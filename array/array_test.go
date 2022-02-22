package array

import (
	"testing"
)

func TestArraySum(t *testing.T) {
	t.Run("Суммирование элементов массива", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := ArraySum(numbers)
		want := 15
		if got != want {
			t.Errorf("given %v want %d got %d", numbers, want, got)
		}
	})
}
