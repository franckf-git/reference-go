package addition

import "testing"

func TestAdd(t *testing.T) {
	var want int = 4
	var got int = Add(2, 2)
	if want != got {
		t.Fatalf(`test give us : %v - and we want %v`, got, want)
	}
}
