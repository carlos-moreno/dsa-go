package main

import (
	"fmt"
	"time"
)

func main() {
	timeStartFindSum := time.Now()
	fmt.Println("Resultado findSum:", findSum(1000000000))
	timeEndFindSum := time.Now()
	fmt.Println("Tempo de execução de findSum:", timeEndFindSum.Sub(timeStartFindSum).Milliseconds())

	fmt.Println("==================================================")

	timeStartFindSumIterate := time.Now()
	fmt.Println("Resultado findSumIterate:", findSumIterate(1000000000))
	timeEndFindSumIterate := time.Now()
	fmt.Println("Tempo de execução de findSumIterate:", timeEndFindSumIterate.Sub(timeStartFindSumIterate).Milliseconds())
}

func findSum(n int) int {
	return n * (n + 1) / 2
}

func findSumIterate(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}
