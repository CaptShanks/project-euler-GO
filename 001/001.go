ckage main

import (
	"fmt"
	"log"
)

func main() {
	var n int //iterations
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("error in reading")
		log.Fatal(err)
	}

	//Slice to store the sum
	var sum []int
	//Looping for n
	for i := 0; i < n; i++ {
		var num int
		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			fmt.Println("error in reading 2")
			log.Fatal(err)
		}

		mul5Num := (num - 1) / 5
		sum5Natural := 5 * (mul5Num * (mul5Num + 1) / 2)

		mul3Num := (num - 1) / 3
		sum3Natural := 3 * (mul3Num * (mul3Num + 1) / 2)
                 
		mul15Num := (num - 1) / 15
		sum15Natural := 15 * (mul15Num * (mul15Num + 1) / 2)
		//subtracting the sum of multiples of 15 as they have been added twice since multiples of 15 are multiples of 3 and 5 both.
		finalSum := sum3Natural + sum5Natural - sum15Natural
		sum = append(sum, finalSum)
	}

	//printing the sum values

	for i := 0; i < len(sum); i++ {
		fmt.Println(sum[i])
	}
}

