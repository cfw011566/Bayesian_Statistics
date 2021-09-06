package main

import (
	"fmt"
	"time"

	xrand "golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

var s = xrand.NewSource(uint64(time.Now().UnixNano()))
var trials = 100000

type Param struct {
	aTrueRate   float64
	aPriorAlpha float64
	aPriorBeta  float64
	bTrueRate   float64
	bPriorAlpha float64
	bPriorBeta  float64
}

func main() {
	param1 := Param{aTrueRate: 0.25, aPriorAlpha: 300, aPriorBeta: 700,
		bTrueRate: 0.3, bPriorAlpha: 300, bPriorBeta: 700}
	fmt.Println("15.3 director of marketing samples =", samples(param1))

	param2 := Param{aTrueRate: 0.25, aPriorAlpha: 30, aPriorBeta: 70,
		bTrueRate: 0.3, bPriorAlpha: 20, bPriorBeta: 80}
	fmt.Println("15.3 lead designer samples =", samples(param2))

	for i := 0; i < 10; i++ {
		fmt.Println(samples(param1), samples(param2))
	}
}

func samples(p Param) int {
	numSamples := 200
	super := -1.0
	for super < 0.95 {
		numSamples += 100
		aTrue, aFalse := runif(numSamples/2, p.aTrueRate)
		bTrue, bFalse := runif(numSamples/2, p.bTrueRate)
		betaA := distuv.Beta{Alpha: float64(aTrue) + p.aPriorAlpha, Beta: float64(aFalse) + p.aPriorBeta, Src: s}
		betaB := distuv.Beta{Alpha: float64(bTrue) + p.bPriorAlpha, Beta: float64(bFalse) + p.bPriorBeta, Src: s}
		aSamples := betaSamples(betaA, trials)
		bSamples := betaSamples(betaB, trials)
		//super = bSuper(aSamples, bSamples)
		s1 := bSuper(aSamples, bSamples)
		aTrue, aFalse = runif(numSamples/2, p.aTrueRate)
		bTrue, bFalse = runif(numSamples/2, p.bTrueRate)
		betaA = distuv.Beta{Alpha: float64(aTrue) + p.aPriorAlpha, Beta: float64(aFalse) + p.aPriorBeta, Src: s}
		betaB = distuv.Beta{Alpha: float64(bTrue) + p.bPriorAlpha, Beta: float64(bFalse) + p.bPriorBeta, Src: s}
		aSamples = betaSamples(betaA, trials)
		bSamples = betaSamples(betaB, trials)
		s2 := bSuper(aSamples, bSamples)
		super = (s1 + s2) / 2
		//fmt.Println("s1", s1, "s2", s2, super)
	}
	return numSamples
}

func runif(n int, rate float64) (int, int) {
	uniform := distuv.Uniform{Min: 0.0, Max: 1.0, Src: s}
	nTrue := 0
	nFalse := 0
	for i := 0; i < n; i++ {
		if uniform.Rand() < rate {
			nTrue++
		} else {
			nFalse++
		}
	}
	return nTrue, nFalse
}

func betaSamples(beta distuv.Beta, n int) []float64 {
	var samples []float64
	for i := 0; i < n; i++ {
		samples = append(samples, beta.Rand())
	}
	return samples
}

func bSuper(aSamples, bSamples []float64) float64 {
	trials := len(aSamples)
	bSuper := 0
	for i := 0; i < trials; i++ {
		if bSamples[i] > aSamples[i] {
			bSuper++
		}
	}
	return float64(bSuper) / float64(trials)
}
