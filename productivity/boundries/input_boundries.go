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
	file, err := os.Open("<fill_correct_path>/seq.09.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

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

	els := strings.Split(res, ";")
	var i int
	min := 10147455745
	max := 0
	for _, el := range strings.Split(els[1], ",") {
		x, _ := strconv.Atoi(el)
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
		i++
	}
	fmt.Println(i)
	fmt.Println(min)
	fmt.Println(max)
}