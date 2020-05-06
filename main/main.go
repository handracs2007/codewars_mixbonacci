// https://www.codewars.com/kata/5811aef3acdf4dab5e000251/train/go

package main

import "fmt"

var fiboMap = make(map[int]int)
var padoMap = make(map[int]int)
var pellMap = make(map[int]int)
var jacMap = make(map[int]int)
var triMap = make(map[int]int)
var tetraMap = make(map[int]int)

func fibonacci(n int) int {
	var cache, exist = fiboMap[n]
	if exist {
		return cache
	}

	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func padovan(n int) int {
	var cache, exist = padoMap[n]
	if exist {
		return cache
	}

	if n == 0 {
		return 1
	} else if n == 1 || n == 2 {
		return 0
	}

	return padovan(n-2) + padovan(n-3)
}

func pell(n int) int {
	var cache, exist = pellMap[n]
	if exist {
		return cache
	}

	if n < 2 {
		return n
	}

	return 2*pell(n-1) + pell(n-2)
}

func jacobsthal(n int) int {
	var cache, exist = jacMap[n]
	if exist {
		return cache
	}

	if n < 2 {
		return n
	}

	return jacobsthal(n-1) + 2*jacobsthal(n-2)
}

func tribonacci(n int) int {
	var cache, exist = triMap[n]
	if exist {
		return cache
	}

	if n == 0 || n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
}

func tetranacci(n int) int {
	var cache, exist = tetraMap[n]
	if exist {
		return cache
	}

	if n == 0 || n == 1 || n == 2 {
		return 0
	} else if n == 3 {
		return 1
	}

	return tetranacci(n-1) + tetranacci(n-2) + tetranacci(n-3) + tetranacci(n-4)
}

func Mixbonacci(pattern []string, length int) []int64 {
	var result = make([]int64, 0)

	if length > 0 && len(pattern) > 0 {
		result = make([]int64, length)
		var commandCountMap = make(map[string]int)

		var funcMap = make(map[string]func(int) int)
		funcMap["fib"] = fibonacci
		funcMap["pad"] = padovan
		funcMap["pel"] = pell
		funcMap["jac"] = jacobsthal
		funcMap["tri"] = tribonacci
		funcMap["tet"] = tetranacci

		var cacheMap = make(map[string]map[int]int)
		cacheMap["fib"] = fiboMap
		cacheMap["pad"] = padoMap
		cacheMap["pel"] = pellMap
		cacheMap["jac"] = jacMap
		cacheMap["tri"] = triMap
		cacheMap["tet"] = tetraMap

		for i := 0; i < length; i++ {
			var command = pattern[i%len(pattern)]

			var count, _ = commandCountMap[command]
			commandCountMap[command] = count + 1

			var output = funcMap[command](count)
			result[i] = int64(output)

			// store the result in the map
			cacheMap[command][count] = output
		}
	}

	return result
}

func main() {
	fmt.Println(Mixbonacci([]string{"fib", "pad", "pel"}, 6))
}
