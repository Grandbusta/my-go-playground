package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"time"
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

func gradingStudents(grades []int32) []int32 {
	new_grades := make([]int32, 0)

	for _, v := range grades {
		if v < 38 {
			new_grades = append(new_grades, v)
		} else {
			round := func(num int32) int32 {
				return int32(math.Ceil(float64(num)/5) * 5)
			}(v)
			if round-v < 3 {
				new_grades = append(new_grades, round)
			} else {
				new_grades = append(new_grades, v)
			}
		}
	}
	return new_grades
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	miniMaxSum([]int32{5, 3, 1, 7, 9})
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
	fmt.Println(timeConversion("12:00:00PM"))
	fmt.Println(gradingStudents([]int32{73, 67, 38, 33}))
	// for i := 0; i < 10; i++ {
	// 	go f(i)
	// }

	// c := make(chan string)
	// go pinger(c)
	// go ponger(c)
	// go printer(c)

	// var input string
	// fmt.Scanln(&input)
}
