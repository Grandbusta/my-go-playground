package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
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

func isEven(n int) bool {
	return n%2 == 0
}

func pageCount(n, p int) {
	// front := 0
	// back := 0
	// hash := make([]interface{}, 0)
	// for i := 1; i <= n; i++ {
	// 	if i == 1 {
	// 		hash = append(hash, 1)
	// 	} else {
	// 		if reflect.TypeOf(hash[len(hash)-1]) == "[]int" && len(hash[len(hash)-1]) {
	// 			hash[len(hash)-1] = append(hash[len(hash)-1], i)
	// 		}
	// 		hash = append(hash, []int{i})
	// 	}
	// }
	// fmt.Println(hash)

}

func plusMinus(arr []int32) {
	positive := 0
	negative := 0
	zero := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			positive++
		}
		if arr[i] < 0 {
			negative++
		}
		if arr[i] == 0 {
			zero++
		}
	}
	getPrecision := func(val int32) string {
		return fmt.Sprintf("%.6f", float64(val)/float64(len(arr)))
	}
	fmt.Println(getPrecision(int32(positive)))
	fmt.Println(getPrecision(int32(negative)))
	fmt.Println(getPrecision(int32(zero)))
}

func staircase(n int32) {
	for i := 0; i < int(n); i++ {
		res := strings.Repeat(" ", int(n)-1-i) + strings.Repeat("#", i+1)
		fmt.Println(res)
	}
}

func main() {
	miniMaxSum([]int32{5, 3, 1, 7, 9})
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
	fmt.Println(timeConversion("12:00:00PM"))
	fmt.Println(gradingStudents([]int32{73, 67, 38, 33}))
	pageCount(5, 2)
	plusMinus([]int32{1, 1, 0, -1, -1})
	staircase(6)
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
