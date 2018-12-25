package sort

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	for _, test := range testCases {
		actual := make([]int, len(test.unsorted))
		copy(actual, test.unsorted)
		QuickSort3Ways(actual)
		if reflect.DeepEqual(actual, test.sorted) {
			t.Logf("PASS: %s", test.massage)
		} else {
			t.Errorf("FAIL: %s\nMergeSort(%v)\nExpected: %#v, Actual: %#v",
				test.massage, test.unsorted, test.sorted, actual)
		}
	}
}
