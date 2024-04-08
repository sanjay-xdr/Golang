// https://playbook.one2n.in/go-bootcamp/go-projects/basic-number-filtering/basic-number-filtering-exercise

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Print("Basic Number Filtering Exercise")
	// c := math.Ceil(1.49)
	// fmt.Printf("%.1f", c)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Print(numbers)

	filteredNumberAny(numbers, isEven, isPrime)
	filteredNumberAll(numbers, isPrime, isEven)
}

type Condition func(int) bool

func isEven(number int) bool {
	return number%2 == 0
}

func isOdd(number int) bool {
	return number%2 != 0
}

func isPrime(number int) bool {

	if number <= 1 {
		return false
	}

	sqrt_val := int(math.Sqrt(float64(number)))

	for i := 2; i <= sqrt_val; i++ {

		if number%i == 0 {
			return false
		}

	}
	return true
}

func filteredNumberAll(numbers []int, condition ...Condition) {

	finalNumbers := []int{}

	for _, num := range numbers {
		allConditionsMet := true

		for _, cond := range condition {

			if !cond(num) {
				allConditionsMet = false
				break
			}
		}

		if allConditionsMet {

			finalNumbers = append(finalNumbers, num)
		}

	}

	fmt.Println("Final Result")

	fmt.Print(finalNumbers)

}

func filteredNumberAny(numbers []int, condition ...Condition) {

	finalNumbers := []int{}

	for _, num := range numbers {
		anyConditionMet := false
		for _, cond := range condition {

			if cond(num) {
				anyConditionMet = true
				break
			}
		}

		if anyConditionMet {
			finalNumbers = append(finalNumbers, num)
		}

	}

	fmt.Println("Final Result")

	fmt.Print(finalNumbers)

}
