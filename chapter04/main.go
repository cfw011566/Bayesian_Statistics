package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	coin := distuv.Binomial{N: 24, P: 0.5}
	fmt.Println("12 head in 24 toss =", coin.CDF(12)-coin.CDF(11))

	gacha := distuv.Binomial{N: 100, P: 0.00720}
	fmt.Println("at least one Survival(0) =", gacha.Survival(0))
	fmt.Println()
	for i := 0.0; i < 2.0; i += 1.0 {
		fmt.Println("CDF", i, gacha.CDF(i))
		fmt.Println("Survival", i, gacha.Survival(i))
	}
}
