package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var i int
	var max int

	for _, el := range strings.Split(res, ",") {
		if max < len(el) {
			max = len(el)
		}
		i++
	}
	fmt.Println(i)
	fmt.Println(max)
}
