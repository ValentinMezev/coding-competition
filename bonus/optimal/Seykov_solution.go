package main

import (
	"bufio"
	"fmt"
	"log"
	// "math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()
	target, _ := strconv.Atoi(input[0])
	// fmt.Println(target)
	solve(input[1], target)
}
func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 90000000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	var res string
	for scanner.Scan() {
		res += scanner.Text()
		break
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

func solve(s string, k int) {
	num := 0
	hist := make([]int, 1000000)
	for _, c := range s {
		if c == ','{
			hist[num]++
			num = 0
		} else {
			num *= 10
			num += int(c - '0')
			// fmt.Println(prev)
		}
	}
	hist[num]++

	printed := 0
	for i := 999999; i >= 0; i-- {

		if hist[i] > 0 {
			// fmt.Printf("%d: %d\n", i, hist[i])
			for j := hist[i]; j > 0; j-- {
				fmt.Printf("%d", i )
				if printed < k - 1 {
					fmt.Print(",")
				}
				printed++
				if printed >= k {
					return
				}
			}

		}
	}

}
