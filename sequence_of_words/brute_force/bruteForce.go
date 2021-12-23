package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println(longestSequence(readInput()))
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

	return strings.Split(res, ",")
}

func longestSequence(elements []string) int {
	var max int
	for _, el := range elements {
		cur := 1 + maxSeqLengthStartingWith(el, elements)
		if cur > max {
			max = cur
		}
	}
	return max
}

func maxSeqLengthStartingWith(s string, elements []string) int {
	var max int
	for _, el := range elements {
		if len(el) == len(s)+1 && strings.HasPrefix(el, s) {
			res := 1 + maxSeqLengthStartingWith(el, elements)
			if res > max {
				max = res
			}
		}
	}
	return max
}
