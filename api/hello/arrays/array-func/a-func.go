package main

import "fmt"

func calculateAvg(arr [6]int, size int) int {
	var sum int
	for _, value := range arr {
		sum += value
	}
	return sum / size

}

func main() {

	scores := [6]int{90, 80, 70, 60, 50, 40}
	avg := calculateAvg(scores, 6)
	fmt.Printf("Average Score: %d", avg)
}
