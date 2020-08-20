package mergesort_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/sergeimuravev/mergesort"
)

func TestMergeSort(t *testing.T) {
	for _, recursive := range []bool{true, false} {
		data := make([]int, 0)
		for i := 0; i < 32; i++ {
			data = append(data, rand.Intn(10))
		}

		t.Logf("Before sorting: %v", data)
		ms := mergesort.BufferedMergeSort{nil, recursive}
		ms.Sort(data)
		t.Logf("After sorting: %v", data)

		if !sort.IsSorted(sort.IntSlice(data)) {
			t.Error("Sequence not sorted.")
			break
		}
	}
}

func BenchmarkRecursive256B(b *testing.B) { benchmarkMergeSort(b, 256, true) }
func BenchmarkRecursive1K(b *testing.B)   { benchmarkMergeSort(b, 1024, true) }
func BenchmarkRecursive4K(b *testing.B)   { benchmarkMergeSort(b, 1024*4, true) }
func BenchmarkRecursive32K(b *testing.B)  { benchmarkMergeSort(b, 1024*32, true) }
func BenchmarkRecursive256K(b *testing.B) { benchmarkMergeSort(b, 1024*256, true) }
func BenchmarkRecursive512K(b *testing.B) { benchmarkMergeSort(b, 1024*512, true) }
func BenchmarkRecursive1M(b *testing.B)   { benchmarkMergeSort(b, 1024*1024, true) }

func BenchmarkNonRecursive256B(b *testing.B) { benchmarkMergeSort(b, 256, false) }
func BenchmarkNonRecursive1K(b *testing.B)   { benchmarkMergeSort(b, 1024, false) }
func BenchmarkNonRecursive4K(b *testing.B)   { benchmarkMergeSort(b, 1024*4, false) }
func BenchmarkNonRecursive32K(b *testing.B)  { benchmarkMergeSort(b, 1024*32, false) }
func BenchmarkNonRecursive256K(b *testing.B) { benchmarkMergeSort(b, 1024*256, false) }
func BenchmarkNonRecursive512K(b *testing.B) { benchmarkMergeSort(b, 1024*512, false) }
func BenchmarkNonRecursive1M(b *testing.B)   { benchmarkMergeSort(b, 1024*1024, false) }

func benchmarkMergeSort(
	b *testing.B,
	bufferSize int,
	isRecursive bool) {

	const size = 1e6 // sample set size
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		data := make([]int, 0)
		for i := 0; i < size; i++ {
			data = append(data, rand.Intn(1000))
		}
		b.StartTimer()

		ms := mergesort.BufferedMergeSort{make([]int, bufferSize), isRecursive}
		ms.Sort(data)

		b.StopTimer()
		if !sort.IsSorted(sort.IntSlice(data)) {
			panic("Sequence not sorted.")
		}
		b.StartTimer()
	}
}
