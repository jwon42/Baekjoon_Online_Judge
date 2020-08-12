package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var count int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fscanln(reader, &count)
	var numbers = make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Fscanln(reader, &numbers[i])
	}
	sort.Ints(numbers)
	for _, idx := range numbers {
		fmt.Fprintln(writer, idx)
	}
}
