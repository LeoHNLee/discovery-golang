// Test Golang with [염재현, "디스커버리 Go 언어", 한빛미디어, 2016]
package discovery_golang

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func Exmaple_string() {
	inputStringSlice := []string{"A", "B", "C"}
	fmt.Println(stringSliceCopy(inputStringSlice)) // stringSliceCopy
	insertStringSlice := stringSliceInsert(inputStringSlice, 2, "D", "E", "F")
	fmt.Println(insertStringSlice)                                         // stringSliceInsert
	fmt.Println(stringSliceInsert2(inputStringSlice, 3, "D", "E", "F"))    // stringSliceInsert2
	deleteStringSlice := stringSliceDeleteOrdered(insertStringSlice, 1, 3) // stringSliceDeleteOrdered
	fmt.Println(deleteStringSlice)
	fmt.Println(stringSliceDeleteUnordered(insertStringSlice, 1, 3))
	fmt.Println(Calc("3 * ( ( 3 + 1 ) * 3 ) / 2"))
	// Output
	// ["A" "B" "C"]
	// ["A" "B" "D" "E" "F" "C"]
	// ["A" "B" "C" "D" "E" "F"]
	// ["A" "E" "F"]
	// ["A" "E" "F"]
	// 18
}

// Check the Map
func Exmaple_map() {
	codeCount := mapCounter("abcbcc")
	correctCount := map[rune]int{97: 1, 98: 2, 99: 3}
	mapCount := reflect.DeepEqual(
		correctCount,
		codeCount,
	)
	fmt.Println(mapCount)
	// Output
	// true
}

// Below two Bad Method of map testing Cuz' they can't check that key pairs of maps are not same

// func TestCount(t *testing.T) {
// 	codeCount := mapCounter("abcbcc")
// 	if len(codeCount) != 3{
// 		t.Error("codeCount:", codeCount)
// 		t.Fatal("count should be 3 but:", len(codeCount))
// 	}
// 	if codeCount['a'] != 1 || codeCount['b'] !=2 || codeCount["c"] != 3 {
// 		t.Error("codeCount mismatch:", codeCount)
// 	}
// }

// func ExampleCount(){
// 	codeCount := mapCounter("abc")
// 	for _, key := range []rune{"a", "b", "c"} {
// 		fmt.Println(string(key), codeCount[key])
// 	}
// }

// IO
func Example_io() {
	// Write
	lines := []string{
		"test1",
		"test2",
	}
	if err := writeText(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}

	// Read
	r := strings.NewReader("test1\ntest2\n")
	copy(lines, lines[2:])
	if err := readText(r, &lines); err != nil {
		fmt.Println(err)
	}

	// // Adj List
	// adjList := [][]int {
	// 	{1},
	// 	{0},
	// }
	// w := bytes.NewBuffer(nil)
	// if err := writeAdjList(w, adjList); err != nil {
	// 	fmt.Println(err.Error())
	// }
	// expected := "2\n1 1\n1 0\n"
	// if expected != w.String() {

	// }

	// Output
	// test1
	// test2
	// [test1 test2]
}
