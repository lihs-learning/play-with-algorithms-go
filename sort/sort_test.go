package sort

import (
	"reflect"
	"runtime"
	"testing"
)

var functionsToTest = []func(arr []int){
	BubbleSort,
	InsertSort,
	SelectSort,
	MergeSort,
	QuickSort2Ways,
	QuickSort3Ways,
}

func Test(t *testing.T) {
	for _, fn := range functionsToTest {
		runSort(fn, t)
	}
}

func runSort(fn func(arr []int), t *testing.T) {
	fnName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	for _, test := range testCases {
		actual := make([]int, len(test.unsorted))
		copy(actual, test.unsorted)
		fn(actual)
		if reflect.DeepEqual(actual, test.sorted) {
			t.Logf("PASS: %s %s", fnName, test.massage)
		} else {
			t.Errorf("FAIL: %s\n%s(%v)\nExpected: %#v, Actual: %#v",
				test.massage, fnName, test.unsorted, test.sorted, actual)
		}
	}
}
