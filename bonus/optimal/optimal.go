package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"

	"os"
	"strconv"
	"strings"
)

type MinHeap struct {
	els                         []int
	putAt, capacity, firstElIdx int
}

func (h *MinHeap) len() int {
	if h == nil {
		return 0
	}
	return h.putAt - 1
}

func (h *MinHeap) empty() bool {
	return h.len() == 0
}

func (h *MinHeap) getMin() (int, error) {
	if h.empty() {
		return 0, errors.New("cannot getMin from empty heap")
	}
	res := h.els[1]
	lastEl := h.putAt - 1
	if lastEl == 1 {
		h.putAt -= 1
		return res, nil
	}
	h.els[1] = h.els[lastEl]
	h.putAt -= 1
	err := h.heapifyDown()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (h *MinHeap) swap(i, j int) {
	tmp := h.els[i]
	h.els[i] = h.els[j]
	h.els[j] = tmp
}

func (h *MinHeap) heapifyUp() error {
	if h.empty() {
		return errors.New("cannot heapifyUp empty heap")
	}
	lastElPos := h.putAt - 1
	if lastElPos == 1 {
		return nil
	}
	for {
		parent := lastElPos / 2
		if h.els[lastElPos] < h.els[parent] {
			h.swap(lastElPos, parent)
		} else {
			break
		}
		if parent == 1 {
			break
		}
		lastElPos = parent
	}
	return nil
}

func (h *MinHeap) heapifyDown() error {
	if h.empty() {
		return errors.New("cannot heapifyDown empty heap")
	}
	index := h.firstElIdx

	half := h.len() / 2
	for index <= half {
		child := index * 2
		right := child + 1
		if right <= h.len() && h.els[right] < h.els[child] {
			child = right
		}

		if h.els[index] > h.els[child] {
			h.swap(child, index)
			index = child
		} else {
			break
		}
	}
	return nil
}

func (h *MinHeap) firstEl() int {
	return h.els[h.firstElIdx]
}

func (h *MinHeap) setAsFirstEl(x int) {
	h.els[h.firstElIdx] = x
}

func (h *MinHeap) add(x int) error {
	if h.putAt <= h.capacity {
		h.els[h.putAt] = x
		h.putAt += 1
		err := h.heapifyUp()
		if err != nil {
			return err
		}
	} else {
		if x > h.firstEl() {
			h.setAsFirstEl(x)
			err := h.heapifyDown()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	diffsBonus := readInput()

	bonus, _ := strconv.Atoi(diffsBonus[0])
	diffs := strings.Split(diffsBonus[1], ",")
	print(topDiffers(bonus, toInts(diffs)))
}

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)

	const maxCapacity = 200000000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	var res string
	for scanner.Scan() {
		res += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return strings.Split(res, ";")
}

func toInts(nums []string) (result []int) {
	for _, num := range nums {
		x, _ := strconv.Atoi(num)
		result = append(result, x)
	}
	return
}

func print(top []int) {
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(top), " ", ",", -1), "[]"))
}

func constructMinHeap(k int) *MinHeap {
	firstEl, putAt := 1, 1
	els := make([]int, k+1)
	return &MinHeap{
		els:        els,
		putAt:      putAt,
		firstElIdx: firstEl,
		capacity:   k,
	}
}

func topDiffers(k int, elements []int) []int {
	mh := constructMinHeap(k)
	for _, el := range elements {
		err := mh.add(el)
		if err != nil {
			os.Exit(1)
		}
	}
	res := make([]int, mh.len())
	i := mh.len() - 1

	for i >= 0 {
		min, err := mh.getMin()
		if err != nil {
			os.Exit(1)
		}
		res[i] = min
		i -= 1
	}
	return res
}
