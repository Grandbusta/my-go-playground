package main

import (
	"fmt"
	"sort"
)

// func intToRoman(num int) string {
// 	values := map[int]string{
// 		"I": 1,
// 		"V": 5,
// 		"X": 10,
// 		"L": 50,
// 		"C": 100,
// 		"D": 500,
// 		"M": 1000,
// 	}
// 	for k, v := range values {
// 		fmt.Println(k, v)
// 	}
// 	return ""
// }

type rect struct {
	width, height int
}

func (r rect) perimeter() int {
	return 2*r.height + 2*r.width
}

func miniMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	first_4 := arr[:4]
	last_4 := arr[len(arr)-4:]
	var first_sum int
	var last_sum int
	for i := 0; i < len(first_4); i++ {
		first_sum += int(first_4[i])
		last_sum += int(last_4[i])
	}
	fmt.Println(first_sum, last_sum)
}

func birthdayCakeCandles(candles []int32) int32 {
	hash := make(map[int32]int)
	for i := 0; i < len(candles); i++ {
		if hash[candles[i]] == 0 {
			hash[candles[i]] = 1
		} else {
			hash[candles[i]]++
		}
	}
	highest := 0
	for _, v := range hash {
		if highest < v {
			highest = v
		}
	}
	return int32(highest)
}

func main() {
	// res := intToRoman(2)
	// fmt.Println(res)
	rectangle := rect{2, 4}
	fmt.Println(rectangle.perimeter())
	miniMaxSum([]int32{5, 3, 1, 7, 9})
	miniMaxSum([]int32{256741038, 623958417, 467905213, 714532089, 938071625})
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
}
