// Package mergesort is a merge sort algorithm implementation using memory buffer.
package mergesort

// BufferedMergeSort implements classic merge sort with optional
// buffer to limit memory consumption
type BufferedMergeSort struct {
	Buffer      []int
	IsRecursive bool
}

// Sort calls recursive or bottom-up merge sort implementation
func (s *BufferedMergeSort) Sort(data []int) {
	if s.Buffer == nil {
		s.Buffer = make([]int, len(data)/2)
	}

	m := &merger{s.Buffer}
	if s.IsRecursive {
		s.sortRecursively(data, 0, len(data)-1, m)
	} else {
		s.sortNonRecursively(data, 0, len(data)-1, m)
	}
}

func (s *BufferedMergeSort) sortRecursively(data []int, left, right int, m *merger) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	s.sortRecursively(data, left, mid, m)
	s.sortRecursively(data, mid+1, right, m)
	m.Merge(data, left, mid, right)
}

func (s *BufferedMergeSort) sortNonRecursively(data []int, left, right int, m *merger) {
	n := len(data)
	for w := 1; w < n; w = 2 * w {
		for i := 0; i < n; i = i + 2*w {
			m.Merge(data, i, min(i+w-1, n), min(i+2*w-1, n-1))
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
