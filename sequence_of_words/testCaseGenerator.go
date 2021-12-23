package main

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

const length = 5_000
const numSequences = 20
const seqLen = 200
const filePath = "<fill_correct_path>/seq.09.in"

func main() {
	rand.Seed(time.Now().UnixNano())
	var all [length]string
	seqLengths := sequencesLengths()
	var sequences []string
	for _, seqLength := range seqLengths {
		sequences = append(sequences, generateSequenceOfLen(randStringRune(rand.Intn(2)+1), seqLength)...)
	}

	for i := 0; i < len(sequences); i++ {
		all[i] = sequences[i]
	}

	for i := len(sequences) - 1; i < len(all); i++ {
		all[i] = randStringRune(rand.Intn(10) + 1)
	}

	rand.Shuffle(len(all), func(i, j int) { all[i], all[j] = all[j], all[i] })

	writeToFile(all)
}

func sequencesLengths() []int {
	var seqLengths []int
	for i := 0; i < numSequences; i++ {
		seqLengths = append(seqLengths, rand.Intn(seqLen)+1)
	}
	return seqLengths
}

func generateSequenceOfLen(el string, n int) (res []string) {
	res = append(res, el)
	for i := 0; i < n; i++ {
		el = el + randStringRune(1)
		res = append(res, el)
	}
	return
}

func writeToFile(result [length]string) {
	f, _ := os.Create(filePath)
	w := bufio.NewWriter(f)
	isFirst := true
	for _, el := range result {
		if isFirst {
			w.WriteString(el)
			isFirst = false
		} else {
			w.WriteString("," + el)
		}
		w.Flush()
	}
}

func randStringRune(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
