package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	norm := distuv.Normal{Mu: 0, Sigma: 1}
	fmt.Println("12.1 =", norm.Survival(5))

	temperatures := []float64{100.0, 99.8, 101.0, 100.5, 99.7}
	fmt.Println("mu =", mu(temperatures))
	fmt.Println("sigma =", sigma(temperatures))
	norm = distuv.Normal{Mu: mu(temperatures), Sigma: sigma(temperatures)}
	fmt.Println("12.2 =", norm.Survival(100.4))

	distances := []float64{4.9 * 2.5 * 2.5, 4.9 * 3 * 3, 4.9 * 3.5 * 3.5, 4.9 * 4 * 4, 4.9 * 2 * 2}
	norm = distuv.Normal{Mu: mu(distances), Sigma: sigma(distances)}
	fmt.Println("12.3 =", norm.Survival(500))
	fmt.Println("12.4 =", norm.CDF(0))
}

func mu(data []float64) float64 {
	sum := 0.0
	for _, f := range data {
		sum = sum + f
	}
	return sum / float64(len(data))
}

func sigma(data []float64) float64 {
	mean := mu(data)
	sum := 0.0
	for _, f := range data {
		sum = sum + (f-mean)*(f-mean)
	}
	return math.Sqrt(sum / float64(len(data)))
}
