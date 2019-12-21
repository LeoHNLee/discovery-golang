// Learn Golang with [염재현, "디스커버리 Go 언어", 한빛미디어, 2016]
package discovery_golang

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	chapter          string = "chapter 3 <string and data sturcture>"
	inputStringSlice        = []string{"A", "B", "C"}
)

// Deep Copy the Slice
func stringSliceCopy(src []string) []string {
	dest := make([]string, len(src))
	for i := range src {
		dest[i] = src[i]
	}
	return dest
}

// Insert string in the Slice
// Have To Error Handling: out of index, etc
// append(): 용량 초과시, capacity 자동 확장
// slice[pos]: 인덱스(length)를 기준으로 out of index check -> 따라서 capacity에 여유가 있더라도 length를 무시하고 출력 불가
// slice[pos:]에서 pos <= capacity인 경우 동작하며, pos <= length이면 빈 slice를 리턴
func stringSliceInsert(src []string, pos int, s ...string) []string {
	length := len(s)
	src = append(src[:pos+length], src[pos:]...)
	for i := range s {
		src[pos+i] = s[i]
	}
	return src
}

// Insert string in the Slice
// Using Copy Method
func stringSliceInsert2(src []string, pos int, s ...string) []string {
	length := len(s)
	src = append(src, s...)
	copy(src[pos+length:], src[pos:])
	for i := range s {
		src[pos+i] = s[i]
	}
	return src
}

// Delete continuous position string in the Slice on Ordered
func stringSliceDeleteOrdered(src []string, pos int, num int) []string {
	num = pos + num
	src = append(src[:pos], src[num:]...)
	return src
}

// Delete continuous position string in the Slice Unordered
// if you use POINTER SLICE, you have to initialize disabled to nil
// if you use STRUCTURE SLICE that include POINTER, you have to initialize disabled to T{}
func stringSliceDeleteUnordered(src []string, pos int, num int) []string {
	length := len(src) - num
	start := length
	num = pos + num
	if num > start {
		start = num
	}
	copy(src[pos:num], src[start:])
	// for i := length; i < len(src); i++ {
	// 	src[i] = nil
	// }
	src = src[:length]
	return src
}

// Calculator build on Stack
func Calc(expression string) int {
	var (
		operations []string
		nums       []int
	)
	pop := func() int {
		length := len(nums) - 1
		last := nums[length]
		nums = nums[:length]
		return last
	}
	// strings.Index는 string에 대해서만 동작한다. []string에 대해서는 error
	operate := func(ops string) {
		for len(operations) > 0 {
			length := len(operations) - 1
			operation := operations[length]
			if strings.Index(ops, operation) < 0 {
				return
			}
			operations = operations[:length]
			if operation == "(" {
				return
			}
			b, a := pop(), pop()
			switch operation {
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, a-b)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, a/b)
			}
		}
	}
	for _, token := range strings.Split(expression, " ") {
		switch token {
		case "(":
			operations = append(operations, token)
		case "+", "-":
			operate("+-*/")
			operations = append(operations, token)
		case "*", "/":
			operate("*/")
			operations = append(operations, token)
		case ")":
			operate("+-*/(")
		default:
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}
	operate("+-*/")
	return nums[0]
}

// Map
// var m map[int]int // no init, can't use
// m := make(map[string]string)
// m := map[int]int{}
// value, ok := m[key] // do not exist the key, return default value of type of map value
// map[key] = value
// delete(m, key)
// Unsafety as Thread (safe only read)

// Count value in map
func mapCounter(s string) map[rune]int {
	charCount := map[rune]int{}
	for _, r := range s {
		charCount[r]++
	}
	return charCount
}

// set
// weakness: can't verify existance of key in set at O(1)

// Using Map to Set
func hasDupeRune(s string) bool {
	runeSet := map[rune]struct{}{} // substitutes struct for bool in order to remove overhead
	for _, r := range s {
		if _, exists := runeSet[r]; exists {
			return true
		}
		runeSet[r] = struct{}{}
	}
	return false
}

// io
func openIntFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	var num int
	if _, err := fmt.Fscanf(f, "%d/n", &num); err == nil {
		return err
	} else {
		return err
	}
}

// output
func writeIntFile(filename string, num int) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := fmt.Fprintf(f, "%d\n", num); err != nil {
		return err
	} else {
		return err
	}
}

// Write the text
func writeText(w io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}

// Read the text file
func readText(r io.Reader, lines *[]string) error {
	scanner := bufio.NewScanner(r) // bufio: ignore \n, read by line, can tokenize
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// Write adjacency list
func writeAdjList(w io.Writer, adjList [][]int) error {
	size := len(adjList)
	if _, err := fmt.Fprintf(w, "%d", size); err != nil {
		return err
	}
	for i := 0; i < size; i++ {
		lsize := len(adjList[i])
		if _, err := fmt.Fprintf(w, "\n%d", lsize); err != nil {
			return err
		}
		for j := 0; j < lsize; j++ {
			if _, err := fmt.Fprintf(w, "%d", adjList[i][j]); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
	return nil
}

// Read adjacency list
func readAdjList(r io.Reader, adjList *[][]int) error {
	var size int
	if _, err := fmt.Fscanf(r, "%d", &size); err != nil {
		return err
	}
	*adjList = make([][]int, size)
	for i := 0; i < size; i++ {
		var lsize int
		if _, err := fmt.Fscanf(r, "\n%d", &lsize); err != nil {
			return err
		}
		(*adjList)[i] = make([]int, lsize)
		for j := 0; j < lsize; j++ {
			if _, err := fmt.Fscanf(r, " %d", &(*adjList)[i][j]); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fscanf(r, "\n"); err != nil {
		return err
	}
	return nil
}
