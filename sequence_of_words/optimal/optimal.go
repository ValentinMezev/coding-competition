package main

import (
	"bufio"
	"errors"
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

type children map[uint8]*Node

type Node struct {
	final bool
	ch    children
}

func (n *Node) longestSeq() (bool, int) {
	if len(n.ch) == 0 {
		if n.final {
			return true, 1
		} else {
			return false, 1
		}
	}
	var maxFlag bool
	var maxL int
	childs := (map[uint8]*Node)(n.ch)
	for _, v := range childs {
		flag, l := v.longestSeq()
		var cur int
		if !n.final {
			flag = false
			cur = l
		} else {
			if flag {
				cur = 1 + l
			} else {
				cur = l
			}
		}

		if cur > maxL {
			maxL = cur
			maxFlag = flag
		}
	}
	return maxFlag, maxL
}

func (t *children) add(s string) error {
	if t == nil {
		return errors.New("cannot add to nil trie")
	}
	if len(s) == 0 {
		return nil
	}

	c := s[0]
	childs := (map[uint8]*Node)(*t)
	if _, ok := childs[c]; !ok {
		childs[c] = &Node{
			final: false,
			ch:    children{},
		}
	}

	if len(s) == 1 {
		childs[c].final = true
		return nil
	}

	return childs[c].ch.add(s[1:])
}

func longestSequence(elements []string) int {
	childs := &children{}
	for _, el := range elements {
		childs.add(el)
	}
	iterChilds := (map[uint8]*Node)(*childs)
	var max int
	for _, el := range iterChilds {
		_, v := el.longestSeq()
		if v > max {
			max = v
		}
	}
	return max
}
