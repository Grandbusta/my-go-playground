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

// func climbingLeaderboard(ranked []int32, player []int32) []int32 {
// 	h := make(map[int32]string, 0)
// 	for i := 0; i < len(ranked); i++ {
// 		if _, ok := h[ranked[i]]; ok {

// 		}
// 	}
// }

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	rank_arr := make([]int, 0)
	for i := 0; i < len(ranked); i++ {
		if i == 0 {
			rank_arr = append(rank_arr, 1)
		} else {
			if ranked[i-1] == ranked[i] {
				rank_arr = append(rank_arr, rank_arr[i-1])
			} else {
				rank_arr = append(rank_arr, rank_arr[i-1]+1)
			}
		}
	}
	idx := 0
	ridx := len(ranked) - 1
	new_ranks := make([]int, 0)
	for idx < len(player) {
		if player[idx] < ranked[ridx] {
			new_ranks = append(new_ranks, rank_arr[ridx]+1)
			idx++
			ridx--
		} else if player[idx] == ranked[ridx] {
			new_ranks = append(new_ranks, rank_arr[ridx])
			idx++
			ridx--
		} else {
			if player[idx] < ranked[ridx-1] {
				new_ranks = append(new_ranks, rank_arr[ridx]+1)
				idx++
				ridx--
			} else if player[idx] == ranked[ridx-1] {
				new_ranks = append(new_ranks, rank_arr[ridx])
				idx++
				ridx--
			} else {
				new_ranks = append(new_ranks, rank_arr[ridx]-1)
				idx++
				ridx--
			}
		}
		fmt.Println(rank_arr[ridx])
	}
	fmt.Println(rank_arr)
	return []int32{3}
}

func sumDigits(n int) int {
	sum := 0
	for n > 0 {
		// find the remainder of n, add it to sum
		sum += n % 10
		// divide n by 10 to get the main and reassign
		n /= 10
	}
	return sum
}

func diffPairs(arr []int, k int) {
	// first approach: brute force
	// idx := 0
	// count := 0
	// for idx < len(arr) {
	// 	for i := idx + 1; i < len(arr); i++ {
	// 		if arr[i]-arr[idx] == k || arr[idx]-arr[i] == k {
	// 			count++
	// 		}
	// 	}
	// 	idx++
	// }
	// fmt.Println(count)

	// Optimized approach
	hash := make(map[int]int, 0)
	count := 0
	for i := 0; i < len(arr); i++ {
		hash[arr[i]] = arr[i]
	}
	for i := 0; i < len(arr); i++ {
		if _, ok := hash[arr[i]-k]; ok {
			count++
		}
	}
	fmt.Println(count)

}

func Permutations(a string, b string) {
	// get the length of a, delare a pointer variable starting from 0, get the next len(a) values from pointer.
	// get the 4 values. run a loop to check if each of the values is a, if any of it does not exist, break the loop, else increase count. increase p
	// a_len := len(a)
	a_hash := make(map[string]int, 0)
	for i := 0; i < len(a); i++ {
		if _, ok := a_hash[string(a[i])]; ok {
			a_hash[string(a[i])]++
		} else {
			a_hash[string(a[i])] = 1
		}
	}
	all_perm := make([]string, 0)
	check_perm := func(value string) (bool, string) {
		status := true
		c_hash := make(map[string]int, 0)
		for i := 0; i < len(value); i++ {
			val := string(value[i])
			if _, ok := a_hash[string(a[i])]; ok {
				c_hash[val]++
			} else {
				c_hash[val] = 1
			}
		}
		for k, v := range c_hash {
			if v != a_hash[k] {
				status = false
				break
			}
		}

		return status, value
	}
	count := 0
	for i := 0; i <= len(b)-len(a); i++ {
		b_values := b[i : i+4]
		status, value := check_perm(b_values)
		if status {
			all_perm = append(all_perm, value)
			count++
		}
	}
	fmt.Println(all_perm)
	fmt.Println(count)
}

func encryption(s string) string {
	// Write your code here
	n_spaces := strings.Replace(s, " ", "", -1)
	s_len := len(n_spaces)
	sq_len := math.Sqrt(float64(s_len))
	// row:=math.Floor(sq_len)
	column := math.Ceil(sq_len)
	grid := make([]string, 0)
	for i := 0; i < len(n_spaces); i += int(column) {
		v := ""
		if i+int(column) > len(n_spaces) {
			v = n_spaces[i:]
		} else {
			v = n_spaces[i : i+int(column)]
		}
		grid = append(grid, v)
	}
	res := make([]string, 0)
	j := 0
	for j < int(column) {
		d := ""
		for i := 0; i < len(grid); i++ {
			val := ""
			if j >= len(grid[i]) {
				val = ""
			} else {
				val = string(grid[i][j])
			}
			d += val
		}
		res = append(res, d)
		j++
	}
	return strings.Join(res, " ")
}

func main() {
	// miniMaxSum([]int32{5, 3, 1, 7, 9})
	// fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
	// fmt.Println(timeConversion("12:00:00PM"))
	// fmt.Println(gradingStudents([]int32{73, 67, 38, 33}))
	// pageCount(5, 2)
	// plusMinus([]int32{1, 1, 0, -1, -1})
	// staircase(6)
	// climbingLeaderboard([]int32{10, 90, 90, 80}, []int32{70, 80, 105})
	// fmt.Println(sumDigits(22))
	diffPairs([]int{1, 7, 5, 9, 2, 12, 3}, 2)
	Permutations("abbc", "cbabadcbbabbcbabaabccbabc")
	fmt.Println(encryption("haveaniceday"))
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
