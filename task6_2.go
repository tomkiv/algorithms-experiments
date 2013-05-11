package main

import (
  "bufio"
	"container/heap"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Item int

// A PriorityQueue implements heap.Interface and holds Items.
type HeapH []Item
type HeapL []Item

func (pq HeapH) Len() int { return len(pq) }

func (pq HeapH) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i] < pq[j]
}

func (pq HeapH) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *HeapH) Push(x interface{}) {
	a := *pq
	n := len(a)
	a = a[0 : n+1]
	item := x.(Item)
	a[n] = item
	*pq = a
}

func (pq *HeapH) Pop() interface{} {
	a := *pq
	n := len(a)
	item := a[n-1]
	*pq = a[0 : n-1]
	return item
}

func (pq HeapL) Len() int { return len(pq) }

func (pq HeapL) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i] > pq[j]
}

func (pq HeapL) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *HeapL) Push(x interface{}) {
	a := *pq
	n := len(a)
	a = a[0 : n+1]
	item := x.(Item)
	a[n] = item
	*pq = a
}

func (pq *HeapL) Pop() interface{} {
	a := *pq
	n := len(a)
	item := a[n-1]
	*pq = a[0 : n-1]
	return item
}

func balance(hl *HeapL, hh *HeapH) {
	llen := hl.Len()
	hlen := hh.Len()
	//fmt.Println(llen, hlen)
	for {
		if llen == hlen || llen-1 == hlen {
			break
		}
		if hlen > llen {
			heap.Push(hl, heap.Pop(hh).(Item))
		} else {
			heap.Push(hh, heap.Pop(hl).(Item))
		}
		llen = hl.Len()
		hlen = hh.Len()
	}
	if hlen > 0 && llen > 0 {
		v1 := heap.Pop(hl).(Item)
		v2 := heap.Pop(hh).(Item)
		if v1 > v2 {
			heap.Push(hh, v1)
			heap.Push(hl, v2)
		} else {
			heap.Push(hl, v1)
			heap.Push(hh, v2)
		}
	}
}

func readFile() {
	fi, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	var sum int
	sum = 0
	hh := make(HeapH, 0, 10000)
	hl := make(HeapL, 0, 10000)
	defer fi.Close()
	r := bufio.NewReader(fi)
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.Trim(line, " \n\t\x0D")
		number, err := strconv.Atoi(line)
		if err != nil {
			break
		}
		//fmt.Println("-------")
		item := Item(number)
		heap.Push(&hl, item)
		balance(&hl, &hh)
		//fmt.Println("hl=", hl)
		//fmt.Println("hh=", hh)

		it := heap.Pop(&hl).(Item)
		//fmt.Println(it)
		sum = sum + int(it)
		heap.Push(&hl, it)
	}
	fmt.Println(sum % 10000)
}

var fileName = flag.String("file", "", "File name")

func main() {
	flag.Parse()
	readFile()
}
