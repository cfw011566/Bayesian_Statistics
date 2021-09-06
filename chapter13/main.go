package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	beta := distuv.Beta{Alpha: 300, Beta: 39700}
	fmt.Println("P(0 - 0.0065) =", beta.CDF(0.0065))
	fmt.Println("P(0.0085 - 1) =", beta.Survival(0.0085))
	fmt.Println("median =", beta.Quantile(0.5))
	fmt.Println("quantile 0.9 =", beta.Quantile(0.9))
	fmt.Println("quantile 0.999 =", beta.Quantile(0.999))

	fmt.Println()
	fmt.Println("Exercise 13.2")
	exercise2()

	fmt.Println()
	fmt.Println("Exercise 13.3")
	exercise3()
}

func exercise2() {
	snowfall := []float64{7.8, 9.4, 10.0, 7.9, 9.4, 7.0, 7.0, 7.1, 8.9, 7.4}
	mean, std := stat.MeanStdDev(snowfall, nil)
	norm := distuv.Normal{Mu: mean, Sigma: std}
	fmt.Println("quantile 0.0005 =", norm.Quantile(0.0005))
	fmt.Println("quantile 0.9995 =", norm.Quantile(0.9995))
}

func exercise3() {
	candyBeta := distuv.Beta{Alpha: 10, Beta: 20}
	lower := candyBeta.Quantile(0.025)
	upper := candyBeta.Quantile(0.975)
	fmt.Println("quantile 0.025 =", lower)
	fmt.Println("quantile 0.975 =", upper)
	fmt.Println(int(lower*40), "-", int(upper*40))
	betaMean := candyBeta.Mean()
	betaStdDev := candyBeta.StdDev()
	betaLower := betaMean - 2*betaStdDev
	betaUpper := betaMean + 2*betaStdDev
	fmt.Println("Candy Beta Mean =", betaMean)
	fmt.Println("Candy Beta StdDev =", betaStdDev)
	fmt.Println("95% confidence interval from 2 sigma =", betaLower, "-", betaUpper)
	fmt.Println(int(betaLower*40), "-", int(betaUpper*40))

	fmt.Println("Binomial")
	candyBinom := distuv.Binomial{N: 30, P: 10.0 / 30.0}
	binomMean := candyBinom.Mean()
	binomStdDev := candyBinom.StdDev()
	binomLower := binomMean - 2*binomStdDev
	binomUpper := binomMean + 2*binomStdDev
	fmt.Println("Candy Binomial Mean =", binomMean)
	fmt.Println("Candy Binomial StdDev =", binomStdDev)
	fmt.Println("95% confidence interval from 2 sigma =", binomLower, "-", binomUpper)
	for i := 1; i <= 40; i++ {
		fmt.Println(i, "=", candyBinom.CDF(float64(i)))
	}
}
