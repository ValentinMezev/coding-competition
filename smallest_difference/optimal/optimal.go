package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	arr1, arr2 := readInput()
	fmt.Println(minDiff(arr1, arr2))
}

func minDiff(arr1, arr2 []int) int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	min := math.MaxInt64
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if abs(arr1[i] - arr2[j]) < min {
			min = abs(arr1[i] - arr2[j])
			if min == 0 {
				return min
			}
		}
		if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}
	return min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readInput() ([]int, []int) {
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

	arrays := strings.Split(res, ";")
	arr1 := strings.Split(arrays[0], ",")
	arr2 := strings.Split(arrays[1], ",")
	return toInts(arr1), toInts(arr2)
}

func toInts(nums []string) (result []int) {
	for _, num := range nums {
		x, _ := strconv.Atoi(num)
		result = append(result, x)
	}
	return
}