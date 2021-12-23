package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	numsAndTarget := readInput()
	target, _ := strconv.Atoi(strings.TrimSpace(numsAndTarget[0]))
	fmt.Println(productivity(toInts(strings.Split(numsAndTarget[1], ",")), target))
}

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)

	const maxCapacity = 2000000
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

func productivity(nums []int, target int) bool {
	for i, num := range nums {
		for _, n := range nums[i+1:] {
			if (target - num) == n {
				return true
			}
		}
	}
	return false
}
