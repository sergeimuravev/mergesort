package mergesort

type merger struct {
	buffer []int
}

// Merges two sorted parts of array using buffer to sequentially
// divide right sub-array into buffered chunks to be merged with the left sub-array
func (m *merger) Merge(data []int, left, mid, right int) {
	if mid < right {
		pivot := min(mid+len(m.buffer), right)
		m.mergeBuffered(data, left, mid, pivot)
		m.Merge(data, left, pivot, right)
	}
}

// Merges two sorted parts using min(len(a),len(b)) memory buffer
// w/o prior knowledge about width of sub-arrays
func (m *merger) mergeBuffered(data []int, left, mid, right int) {
	leftW := 1 + mid - left
	rightW := right - mid
	if leftW < rightW {
		m.mergeLeft(data, left, mid, right, m.buffer[:min(leftW, len(m.buffer))])
	} else {
		m.mergeRight(data, left, mid, right, m.buffer[:min(rightW, len(m.buffer))])
	}
}

// Merges two sorted parts using min(len(a),len(b)) memory buffer
// assuming that left part of array is shorter than the right one
func (m *merger) mergeLeft(data []int, left, mid, right int, buffer []int) {
	// Populate buffer initially
	copy(buffer, data[:len(buffer)])

	// Loop variables
	j := mid + 1
	k := 0

	for i := left; i <= right; i++ {
		// Merge buffer and data
		if k == len(buffer) {
			data[i] = data[j]
			j++
		} else if j > right {
			data[i] = buffer[k]
			k++
		} else if buffer[k] < data[j] {
			data[i] = buffer[k]
			k++
		} else {
			data[i] = data[j]
			j++
		}
	}
}

// Merges two sorted parts using min(len(a),len(b)) memory buffer
// assuming that right part of array is shorter than the right one
func (m *merger) mergeRight(data []int, left, mid, right int, buffer []int) {
	// Populate buffer initially
	copy(buffer, data[mid+1:])

	// Loop variables
	j := mid
	k := len(buffer) - 1

	for i := right; i >= left; i-- {
		// Merge buffer and data
		if k < 0 {
			data[i] = data[j]
			j--
		} else if j < left {
			data[i] = buffer[k]
			k--
		} else if buffer[k] > data[j] {
			data[i] = buffer[k]
			k--
		} else {
			data[i] = data[j]
			j--
		}
	}
}
