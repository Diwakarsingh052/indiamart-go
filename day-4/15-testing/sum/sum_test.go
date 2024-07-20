package sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSumInt(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	want := 15
	got := SumInt(x)

	// Checking if the function's output matches the expected output.
	// If they are not equal, the test will fail and print the error message with the expected and got values.
	if got != want {
		t.Errorf("sum of 1 to 5 should be %v; got %v", want, got)
		//t.Fatalf()
	}

	// create a case where input is nil
	want = 0
	got = SumInt(nil)
	if got != want {
		t.Errorf("sum of nil should be %v; got %v", want, got)
		//t.Fatalf()
	}

}

func TestSumIntTableTest(t *testing.T) {

	tt := []struct {
		name    string // Name of the test
		numbers []int  // Input numbers for the test
		want    int    // Expected result for the test
	}{
		// each instance of struct in the slice would act like a test case
		{
			name:    "one to five",
			numbers: []int{1, 2, 3, 4, 5},
			want:    15,
		},
		{
			name:    "nil slice", // Test case where a nil slice is passed, expecting sum as 0
			numbers: nil,
			want:    0,
		},
		{
			name:    "one minus one", // Test case where 1 and -1 are summed
			numbers: []int{1, -1},
			want:    0, // Expect the sum to be 0 in this case
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := SumInt(tc.numbers)
			require.Equal(t, tc.want, got)
		})
	}
}
