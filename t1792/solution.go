package t1792

import "container/heap"

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	pq := make(PriorityQueue, 0, len(classes))
	heap.Init(&pq)

	// Add all classes to the priority queue
	for _, c := range classes {
		class := []int{c[0], c[1]}
		heap.Push(&pq, class)
	}

	// Add extra students to classes that will benefit most
	for i := 0; i < extraStudents; i++ {
		c := heap.Pop(&pq).([]int)
		c[0]++ // increment pass count
		c[1]++ // increment total count
		heap.Push(&pq, c)
	}

	// Calculate the average ratio
	totalRatio := 0.0
	for _, c := range pq {
		totalRatio += float64(c[0]) / float64(c[1])
	}
	return totalRatio / float64(len(classes))
}

type PriorityQueue [][]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// Calculate the improvement in ratio when adding one student
	// We want to prioritize classes that will have the highest improvement
	// Improvement = (pass+1)/(total+1) - pass/total
	// We use a max heap, so we return true when i should come before j

	passI, totalI := pq[i][0], pq[i][1]
	passJ, totalJ := pq[j][0], pq[j][1]

	// Calculate improvement for class i
	improvementI := float64(passI+1)/float64(totalI+1) - float64(passI)/float64(totalI)
	// Calculate improvement for class j
	improvementJ := float64(passJ+1)/float64(totalJ+1) - float64(passJ)/float64(totalJ)

	return improvementI > improvementJ
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
