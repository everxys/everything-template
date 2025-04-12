package util

import (
	"everything-template/pkg/util"
	"testing"
)

func TestEqualSlice(t *testing.T) {
	if !util.EqualSlice([]int{1, 2, 3}, []int{1, 2, 3}) {
		t.Error("Expected slices to be equal")
	}
}

func TestEqualMap(t *testing.T) {
	if !util.EqualMap(map[string]int{"a": 1}, map[string]int{"a": 1}) {
		t.Error("Expected maps to be equal")
	}
}
