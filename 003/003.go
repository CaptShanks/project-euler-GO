package main

import "fmt"

func main() {
	var num int
	fmt.Scanf("%d", &num)

	//Iterating based on the above number
	for i := 0; i < num; i++ {
		//reading the value
		var (
			input        int
			primeNumbers []int = make([]int, 2)
		)
		primeNumbers[0] = 2
		primeNumbers[1] = 3
		fmt.Scanf("%d", &input)
		numberPrimeClosestToInput := (input + 1) / 6

		for i := 1; i <= numberPrimeClosestToInput; i++ {
			val := 6*i + 1
			val2 := 6*i - 1
			primeNumbers = append(primeNumbers, val2)
			primeNumbers = append(primeNumbers, val)

		}
		for j := len(primeNumbers) - 1; j > 0; j-- {
			var count int = 0
			if input%primeNumbers[j] == 0 {
				for i := 2; i < primeNumbers[j]; i++ {
					if primeNumbers[j]%i == 0 {
						count++
						break
					}
				}
				if count == 0 {
					fmt.Println(primeNumbers[j])
					break
				}
			}
		}
	}
}
