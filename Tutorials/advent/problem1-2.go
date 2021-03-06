package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func contains(list []int, newVal int) bool {
	count := 0
	for _, val := range list {
		if val == newVal {
			count++
		}
	}
	if count > 1 {
		return true
	}
	return false
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	check(err)

	nums := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Println(data[i])
		nums[i], _ = strconv.Atoi(string(data[i]))
	}
	fmt.Println(nums[0:5])

	// states := make([]int, 1) //values reached
	// current := 0             //current state value
	// index := 0               //current index in data
	// run := true
	// for run {
	// 	if index >= len(data) {
	// 		index = 0
	// 	}
	// 	current += int(data[index])
	// 	states = append(states, current)
	// 	run = !contains(states, current)
	// 	index++
	// 	if index > 5 {
	// 		break
	// 	}
	// }
	// fmt.Println("First repeat", current)
	// fmt.Println(states)
}
