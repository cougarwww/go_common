package algorithm

import (
	"testing"
)

func Test_MaoPaoSort(t *testing.T) {
	intArray := MaoPaoSort([]int{1, 3, 2, 5})
	t.Log(intArray)
}

func Test_SelectSort(t *testing.T) {
	intArray := SelectSort([]int{1, 3, 2, 5})
	t.Log(intArray)
}

func Test_InsertSort(t *testing.T) {
	intArray := InsertSort([]int{1, 3, 2, 5})
	t.Log(intArray)
}

func Test_ShellSort(t *testing.T) {
	intArray := ShellSort([]int{1, 3, 2, 5,7,9,2})
	t.Log(intArray)
}
func Test_HeapSort(t *testing.T) {
	intArray := HeapSort([]int{8,5,0,3,7,1,2})
	t.Log(intArray)
}
