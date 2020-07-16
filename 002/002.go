package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	var iteration int //getting the number of iterations
	_, err := fmt.Scanf("%d", &iteration)
	if err != nil {
		log.Fatal(err)
	}
	//Looping for the number of iterations
	for i := 0; i < iteration; i++ {
		num := big.NewInt(0)
		_, err := fmt.Scan(num)
		if err != nil {
			log.Fatal(err)
		}
		sum := big.NewInt(0)
		a := big.NewInt(0)
		b := big.NewInt(1)
		c := big.NewInt(0)
		div := big.NewInt(2)
		for {
			a.Add(a, b)
			if a.Cmp(num) < 1 {
				if c.Mod(a, div); c.Cmp(big.NewInt(0)) == 0 {
					sum.Add(a, sum)
				}
			} else {
				break
			}
			a, b = b, a
		}

		fmt.Println(sum)
	}
}
