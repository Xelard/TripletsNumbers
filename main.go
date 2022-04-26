package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var result [][]int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	str := strings.Split(scanner.Text(), ",")
	numbers := getNumbers(str)
	l := len(numbers)

	for i := 0; i < l; i++ {
		j := i + 1
		k := l - 1
		for j < k {
			if numbers[j]+numbers[k] == -1*numbers[i] {
				var t []int
				t = append(t, numbers[i])
				t = append(t, numbers[j])
				t = append(t, numbers[k])
				result = append(result, t)

				for j+1 < k && numbers[j+1] == numbers[k] {
					j++
				}
				for k-1 > j && numbers[k-1] == numbers[k] {
					k--
				}
				j++
				k--
			} else if numbers[j]+numbers[k] < -1*numbers[i] {
				j++
			} else {
				k--
			}
		}
		for i+1 < l && numbers[i+1] == numbers[i] {
			i++
		}
	}

	for _, v := range result {
		print(v)
	}
}

func getNumbers(str []string) []int {
	var a []int
	for _, v := range str {
		n, _ := strconv.Atoi(v)
		a = append(a, n)
	}
	sort.Ints(a)
	return a
}

func print(result []int) {
	var resultStr []string
	for _, v := range result {
		resultStr = append(resultStr, strconv.Itoa(v))
	}
	fmt.Print(strings.Join(resultStr, ","))
	fmt.Print(" ")
}
