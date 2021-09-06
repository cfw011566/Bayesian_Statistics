package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	black_box := distuv.Beta{Alpha: 14, Beta: 27}
	fmt.Println("The probability of getting two coins from the box is 0.5 =",
		black_box.CDF(0.5))

	gacha := distuv.Beta{Alpha: 5, Beta: 1195}
	fmt.Println("The chance of pulling a Bradley Efron is greater than 0.005 =",
		gacha.Survival(0.005))

	coin := distuv.Beta{Alpha: 4, Beta: 6}
	fmt.Println("5.1 =", coin.Survival(0.6))

	coin = distuv.Beta{Alpha: 9, Beta: 11}
	fmt.Println("5.2 =", coin.CDF(0.55)-coin.CDF(0.45))

	coin = distuv.Beta{Alpha: 109, Beta: 111}
	fmt.Println("5.3 =", coin.CDF(0.55)-coin.CDF(0.45))
}
