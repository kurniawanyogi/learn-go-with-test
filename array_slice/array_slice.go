package main

import "fmt"

func main() {
	// arrays are fixed capacity, its means [5]int different from [4]int, and make compile error
	//fmt.Println(Sum([5]int{1, 2, 3, 4, 5}))

	// slices can have any size
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))

	fmt.Println(SumAll([]int{}, []int{1, 0}))

	fmt.Println(SumAllTails([]int{}, []int{1, 2, 3}))
}

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAll(slices ...[]int) []int {
	var groupOfTotal []int
	for _, slice := range slices {
		groupOfTotal = append(groupOfTotal, Sum(slice))
	}
	return groupOfTotal
}

// untuk sum isi slices selain element pertama
func SumAllTails(slices ...[]int) []int {
	var groupOfTotal []int
	for _, slice := range slices {
		if len(slice) != 0 {
			groupOfTotal = append(groupOfTotal, Sum(slice[1:]))
		} else {
			groupOfTotal = append(groupOfTotal, 0)
		}

	}
	return groupOfTotal
}
