package utils

import "testing"

func TestEqual(t *testing.T) {
	tests := []struct {
		a      []string
		b      []string
		result bool
	}{
		{nil, nil, true},
		{[]string{"a"}, []string{"a"}, true},
		{[]string{"a", "b"}, []string{"a", "b"}, true},
		{[]string{"b", "a"}, []string{"a", "b"}, false},
		{[]string{"a", "b"}, []string{"a"}, false},
		{[]string{"a"}, []string{"a", "b"}, false},
		{[]string{"A"}, []string{"a"}, false},
		{[]string{"a"}, []string{}, false},
		{[]string{""}, []string{}, false},
	}

	for _, test := range tests {
		result := Equal(test.a, test.b)
		if result != test.result {
			var check string
			if test.result {
				check = "equal"
			} else {
				check = "not equal"
			}

			t.Errorf("Expected two slices to be %s, they weren't", check)
		}
	}
}
