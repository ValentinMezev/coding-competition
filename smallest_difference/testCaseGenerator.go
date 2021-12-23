package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const filePath = "<fill_correct_path>/seq.09.in"

func randInt(min, max int) int {
	return rand.Intn(max - min) + min
}

func main() {
	desiredDiff := 137
	k := 50_000
	rand.Seed(time.Now().UnixNano())
	els := make(map[int]bool)
	var arr1 []int
	for i := 0; i < k; i++ {
		x := randInt(-200_000, 100_000)
		els[x] = true
		arr1 = append(arr1, x)
	}
	var arr2 []int
	for i := 0; i < k; {
		x := randInt(-100_000, 200_000)
		isOk := true
		for i := 0; i < desiredDiff; i++ {
			if els[x - i] || els[x + i] {
				isOk = false
				break
			}
		}
		if !isOk {
			continue
		}
		i++
		arr2 = append(arr2, x)
	}
	writeToFile(arr1, arr2)
}

func writeToFile(arr1, arr2 []int) {
	f, _ := os.Create(filePath)
	w := bufio.NewWriter(f)
	isFirst := true
	for _, el := range arr1 {
		if isFirst {
			w.WriteString(strconv.Itoa(el))
			isFirst = false
		} else {
			w.WriteString("," + strconv.Itoa(el))
		}

	}
	w.WriteString(";")
	isFirst = true
	for _, el := range arr2 {
		if isFirst {
			w.WriteString(strconv.Itoa(el))
			isFirst = false
		} else {
			w.WriteString("," + strconv.Itoa(el))
		}

	}
	w.Flush()
}

