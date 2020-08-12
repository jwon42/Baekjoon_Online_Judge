package main

import "fmt"

func main() {
	var testcase, k, n int
	var apart [15][15]int
	for i := 1; i <= 14; i++ {
		apart[0][i] = i
	}
	for i := 1; i <= 14; i++ {
		for j := 1; j <= 14; j++ {
			apart[i][j] = apart[i-1][j] + apart[i][j-1]
		}
	}
	fmt.Scanf("%d", &testcase)
	for i := 1; i <= testcase; i++ {
		fmt.Scanf("%d", &k)
		fmt.Scanf("%d", &n)
		fmt.Printf("%d\n", apart[k][n])
	}
}
