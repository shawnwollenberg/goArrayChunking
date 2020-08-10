package Arraychunking

import (
	"reflect"
	"testing"
)

func TestArraychunking(t *testing.T) {
	want := [][]int{{1, 2}, {3, 4}, {5}}
	if got := Arraychunking([]int{1, 2, 3, 4, 5}, 2); !reflect.DeepEqual(want, got) {
		t.Errorf("Arraychunking() = %q, want %q", got, want)
	}
}
