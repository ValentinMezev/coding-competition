package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const filePath = "<fill_correct_path>/seq.09.in"

func main() {
	k := 1_000_000
	rand.Seed(time.Now().UnixNano())
	var res []int
	for i := 0; i < k; i++ {
		x := rand.Intn(1_000_000)
		res = append(res, x)
	}
	writeToFile(res, 5)
}

func writeToFile(result []int, k int) {
	f, _ := os.Create(filePath)
	w := bufio.NewWriter(f)
	isFirst := true
	w.WriteString(strconv.Itoa(k) + ";")
	for _, el := range result {
		if isFirst {
			w.WriteString(strconv.Itoa(el))
			isFirst = false
		} else {
			w.WriteString("," + strconv.Itoa(el))
		}

	}
	w.Flush()
}

