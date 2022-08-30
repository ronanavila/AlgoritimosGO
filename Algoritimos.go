// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

func main() {

	array := []int{8, 10, 5, 1, -5, 7, 9, 3}

	fmt.Println(sumTwoNumber2(array, 13))

}



func sumTwoNumber3(array []int, num int) []int {

	sum := []int{}

	if len(array) < 1 {
		return sum
	}
	
	i := 0
	j := len(array)-1

	sort.Ints(array)

	for i <j {
		value := array[i] + array[j]
		if value == num{
			sum = append(sum, i, j)
			return sum
		}
		
		if value < num{
			i++
		}

		if value > num{
			j--
		}
	}
	
	return sum
}

func sumTwoNumber2(array []int, num int) []int {
	sum := []int{}

	if len(array) < 1 {
		return sum
	}
	hash := make(map[int]int)

	for i := 0; i < len(array); i++ {
		v1, v2 := hash[array[i]]
		if v2 {
			sum = append(sum, i, v1)
			return sum
		} else {
			hash[num-array[i]] = i
		}		

	}
	return sum
}

func sumTwoNumber(array []int, num int) []int {
	sum := []int{}

	if len(array) < 1 {
		return sum
	}
	for i := 0; i < len(array)-1; i++ {
		for j := 1; j < len(array); j++ {
			if (array[i] + array[j]) == num {

				sum = append(sum, i, j)
				return sum
			}
		}
	}
	return sum
}
