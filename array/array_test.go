package array

import (
	"reflect"
	"testing"
)

func TestPrepend(t *testing.T) {
	arr := []string{"a", "b", "c"}
	res := Prepend(arr, "d")
	ans := []string{"d", "a", "b", "c"}

	if !reflect.DeepEqual(res, ans) {
		t.Errorf("Incorrect: %s. Wanted: %s", res, ans)
	}

	res = Prepend(arr, "d", "e", "f")
	ans = []string{"d", "e", "f", "a", "b", "c"}

	if !reflect.DeepEqual(res, ans) {
		t.Errorf("Incorrect: %s. Wanted: %s", res, ans)
	}
}
