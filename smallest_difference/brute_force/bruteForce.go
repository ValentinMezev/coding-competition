package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr1, arr2 := readInput()
	fmt.Println(minDiff(arr1, arr2))
}

func minDiff(arr1, arr2 []int) int {
	min := math.MaxInt64
	for _, e1 := range arr1 {
		for _, e2 := range arr2 {
			if abs(e1 - e2) < min {
				min = abs(e1 - e2)
				if min == 0 {
					return min
				}
			}
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