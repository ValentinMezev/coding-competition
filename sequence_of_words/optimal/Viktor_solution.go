package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := readInput()
	fmt.Println(bonus(strings.Split(input[0], ",")))
}

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 20000000
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

func bonus(words []string) int {
	ans := 0
	m := make(map[string]struct{})

	for _, word := range words {
		m[word] = struct{}{}
	}

	for _, word := range words {
		last := word
		curr := 1
		for {
			last = last[0 : len(last)-1]
			if _, ok := m[last]; ok {
				curr++
			} else {
				break
			}
		}
		if curr > ans{
			ans = curr
		}
	}
	return ans
}
