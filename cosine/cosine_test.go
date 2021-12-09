package cosine

import "testing"

func TestCosine(t *testing.T) {

	// define vars
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{1, 2, 3, 4, 5, 6}

	cos, err := CalculateCosine(a, b)
	if err != nil {
		t.Error(err)
	}
	if cos < 0.99 {
		t.Error("Expected cosine similarity of 1, got instead", cos)
	}

	a = []int{1, 2, 3, 4, 5, 6}
	b = []int{1, 2, 3, 4, 5, 9}

	_, err = CalculateCosine(a, b)

	if err != nil {
		t.Error(err)
	}

	a = []int{1, 2, 3, 4, 5, 6}
	b = []int{10, 11, 12, 13, 14, 15}

	cos, err = CalculateCosine(a, b)

	if err != nil {
		t.Error(err)
	}
	if cos < 0.90 {
		t.Error("Expected similarity of greater than 0.90, got instead ", cos)
	}
}
