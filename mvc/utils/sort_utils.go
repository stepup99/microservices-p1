package utils

import "sort"

func BubbleSort(elms []int) []int {
	keepRunning := true
	for keepRunning {
		keepRunning = false
		for i := 0; i < len(elms)-1; i++ {
			if elms[i] > elms[i+1] {
				elms[i], elms[i+1] = elms[i+1], elms[i]
				keepRunning = true
			}
		}
	}
	return elms
}

func GenerateArr(size int) []int {
	elms := make([]int, size)
	i := 0
	for j := size - 1; j >= 0; j-- {
		elms[i] = j
		i++
	}
	return elms
}

func InbuiltSort(elms []int) []int {
	sort.Ints(elms)
	return elms
}

func Sort(elems []int) {
	if len(elems) < 1000 {
		BubbleSort(elems)
		return
	}
	InbuiltSort(elems)
	return
}
