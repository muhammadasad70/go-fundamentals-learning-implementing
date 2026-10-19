package testing

// import (
// 	"testing"
// )

// func TestAdd(t *testing.T) {
// 	result := Add(20, 10)
// 	expected := 50

// 	if result != expected {
// 		t.Errorf("Expected %d got %d", expected, result)
// 	}
// 	t.Log("Testing Completed")

// }

// func TestAdd_2(t *testing.T) {
// 	result := Add(20, 10)
// 	expected := 50

// 	if result != expected {
// 		t.Fatalf("Expected %d got %d", expected, result)
// 	}
// 	t.Log("Testing Completed")

// }

// func TestTableDriven(t *testing.T) {
// 	tests := []struct {
// 		a        int
// 		b        int
// 		expected int
// 	}{
// 		{10, 20, 30},
// 		{5, 5, 10},
// 		{-10, 5, -5},
// 		{0, 0, 0},
// 		{-10, -20, -30},
// 	}

// 	for _, tc := range tests {
// 		result := Add(tc.a, tc.b)
// 		if result != tc.expected {
// 			t.Errorf("Add(%d,%d)=%d; wants %d", tc.a, tc.b, result, tc.expected)
// 		}
// 	}
// }
