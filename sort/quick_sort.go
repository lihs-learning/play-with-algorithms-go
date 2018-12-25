package sort

import (
	"math/rand"
	"time"
)

func QuickSort2Ways(arr []int) {
	if len(arr) < 16 {
		InsertSort(arr)
		return
	}

	randNum := randInt(len(arr) - 1)

	arr[0], arr[randNum] = arr[randNum], arr[0]
	v := arr[0]

	ltTail, gtHead := 1, len(arr)-1

	for {
		for ltTail < len(arr) && arr[ltTail] < v {
			ltTail++
		}
		for gtHead > 1 && arr[gtHead] > v {
			gtHead--
		}
		if ltTail > gtHead {
			break
		}
		arr[ltTail], arr[gtHead] = arr[gtHead], arr[ltTail]
		ltTail++
		gtHead--
	}

	arr[0], arr[gtHead] = arr[gtHead], arr[0]

	QuickSort2Ways(arr[:gtHead])
	QuickSort2Ways(arr[gtHead:])
}

func QuickSort3Ways(arr []int) {
	// 注意如果没有插入排序，需要考虑很多边界情况
	if len(arr) < 16 {
		InsertSort(arr)
		return
	}

	randNum := randInt(len(arr) - 1)

	arr[0], arr[randNum] = arr[randNum], arr[0]

	ltTail, gtHead := 0, len(arr)
	i := 1

	v := arr[0]

	for i < gtHead {
		switch {
		case arr[i] == v:
			i++
		case arr[i] < v:
			arr[i], arr[ltTail+1] = arr[ltTail+1], arr[i]
			ltTail++
			i++
		case arr[i] > v:
			arr[i], arr[gtHead-1] = arr[gtHead-1], arr[i]
			gtHead--
		}
	}

	arr[0], arr[ltTail] = arr[ltTail], arr[0]

	QuickSort3Ways(arr[:ltTail+1])
	QuickSort3Ways(arr[gtHead:])
}

func randInt(n int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(n)
}
