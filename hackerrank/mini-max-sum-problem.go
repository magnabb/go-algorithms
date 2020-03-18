// https://www.hackerrank.com/challenges/mini-max-sum/problem

/*
Given five positive integers,
find the minimum and maximum values that can be calculated by summing exactly four of the five integers.

Then print the respective minimum and maximum values as a single line of two space-separated long integers.

For example arr = [1,3,5,7,9] Our minimum sum is 1+3+5+7=16 and our maximum sum is 3+5+7+9=24. We would print
16 24
*/
package hackerrank

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int64) {
	var calculated = make([]int64, len(arr))

	for idx := range arr {
		for innerIdx, innerVal := range arr {
			if innerIdx == idx {
				continue
			}

			calculated[idx] += innerVal
		}
	}

	min, max := min(calculated), max(calculated)
	fmt.Println(min, max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int64

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int64(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func min(arr []int64) int64 {
	min := arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
	}

	return min
}

func max(arr []int64) int64 {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	return max
}
