package datastructure

import (
	"container/heap"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPriorityQueue(t *testing.T) {
	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &Item{
		Value:    "123",
		Priority: 1,
	})
	heap.Push(&pq, &Item{
		Value:    1234,
		Priority: 3,
	})
	heap.Push(&pq, &Item{
		Value:    "abc",
		Priority: 2,
	})
	for _, item := range pq {
		fmt.Printf("%d(%d): %v\n", item.Priority, item.index, item.Value)
	}

	item := &Item{
		Value:    0.1,
		Priority: 0,
	}
	heap.Push(&pq, item)
	pq.Update(item, 0.2, 10)

	fmt.Println("--------")
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%d(%d): %v\n", item.Priority, item.index, item.Value)
	}
}

func TestPriorityQueueTimeCost(t *testing.T) {
	pq := make(PriorityQueue, 0)

	start := time.Now()
	for i := 0; i < 1000000; i++ {
		heap.Push(&pq, &Item{
			Value:    rand.Int(),
			Priority: rand.Int(),
		})
	}
	fmt.Printf("1m push cost: %v\n", time.Since(start))

	start = time.Now()
	for pq.Len() > 0 {
		heap.Pop(&pq)
	}
	fmt.Printf("1m pop cost: %v\n", time.Since(start))

	start = time.Now()
	item := &Item{
		Value:    rand.Int(),
		Priority: rand.Int(),
	}
	heap.Push(&pq, item)
	fmt.Printf("one push cost: %v\n", time.Since(start))

	start = time.Now()
	pq.Update(item, rand.Int(), rand.Int())
	fmt.Printf("update cost: %v\n", time.Since(start))

	for i := 0; i < 1000000; i++ {
		heap.Push(&pq, &Item{
			Value:    rand.Int(),
			Priority: rand.Int(),
		})
	}
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		heap.Push(&pq, &Item{
			Value:    rand.Int(),
			Priority: rand.Int(),
		})
	}
	fmt.Printf("2nd 1m push cost: %v\n", time.Since(start))
}
