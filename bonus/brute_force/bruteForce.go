package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	diffsBonus := readInput()
	bonus, _ := strconv.Atoi(strings.TrimSpace(diffsBonus[0]))
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

func topDiffers(k int, elements []int) []int {
	sort.Slice(elements, func(i, j int) bool {
		return elements[i] > elements[j]
	})
	return elements[:k]
}
