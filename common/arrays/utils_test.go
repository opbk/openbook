package arrays

import "testing"

func TestIntToStringWithInt(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []string{"1", "2", "3", "4", "5"}
	for i, el := range IntToString(input) {
		if el != expected[i] {
			t.Fail()
		}
	}
}

func TestIntToStringWithInt64(t *testing.T) {
	input := []int64{int64(1), int64(2), int64(3), int64(4), int64(5)}
	expected := []string{"1", "2", "3", "4", "5"}
	for i, el := range Int64ToString(input) {
		if el != expected[i] {
			t.Fail()
		}
	}
}
