package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const filePath = "<fill_correct_path>/seq.09.in"
const k = 100_000
const numLength = 5_000_000_000


func main() {
	rand.Seed(time.Now().UnixNano())
	var arr []int
	for i := 0; i < k; i++ {
		x := rand.Intn(numLength)
		arr = append(arr, x)
	}
	writeToFile(arr, arr[0] + arr[k/2])

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
