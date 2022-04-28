package main

import (
	"fmt"
	"sort"
	"strconv"
)

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

func timeConversion(s string) string {
	ext := s[len(s)-2:]
	hour := s[:2]
	if ext == "PM" && hour != "12" {
		temp, _ := strconv.Atoi(hour)
		hour = strconv.Itoa(12 + temp)
	} else if ext == "AM" && hour == "12" {
		hour = "00"
	}
	return hour + s[2:8]
}

func main() {
	miniMaxSum([]int32{5, 3, 1, 7, 9})
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
	fmt.Println(timeConversion("12:00:00PM"))
}
