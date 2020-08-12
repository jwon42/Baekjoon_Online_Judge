package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type info struct {
	age  int
	name string
}

func main() {
	var cnt int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &cnt)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var member = make([]info, cnt)
	for i := 0; i < cnt; i++ {
		fmt.Fscanln(reader, &member[i].age, &member[i].name)
	}
	sort.SliceStable(member, func(i, j int) bool {
		return member[i].age < member[j].age
	})
	for i := 0; i < cnt; i++ {
		fmt.Fprintln(writer, member[i].age, member[i].name)
	}
}
