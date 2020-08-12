package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var numbers = make([]int, 9)
	var sum int
	for i := 0; i < 9; i++ {
		fmt.Fscanln(reader, &numbers[i])
		sum += numbers[i]
	}
	for i := 0; i < 9; i++ {
		for j := i + 1; j < 9; j++ {
			if sum-numbers[i]-numbers[j] == 100 {
				numbers[i], numbers[j] = 99, 99
				sort.Ints(numbers)
				for _, i := range numbers[:7] {
					fmt.Fprintln(writer, i)
				}
			}
		}
	}
}
