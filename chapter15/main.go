package main

import (
	"fmt"
	"image/color"
	"log"
	"sort"
	"time"

	xrand "golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var samples = 500
var s = xrand.NewSource(uint64(time.Now().UnixNano()))
var trials = 100000

func main() {
	//fig15_2()

	priorAlpha := 3.0
	priorBeta := 10.0 - priorAlpha

	betaA := distuv.Beta{Alpha: 36 + priorAlpha, Beta: 114 + priorBeta, Src: s}
	betaB := distuv.Beta{Alpha: 50 + priorAlpha, Beta: 100 + priorBeta, Src: s}
	aSamples := betaSamples(betaA, trials)
	bSamples := betaSamples(betaB, trials)
	fmt.Println("B > A =", bSuper(aSamples, bSamples))

	//fig15_3(aSamples, bSamples)

	//fig15_4(aSamples, bSamples)

	priorAlpha = 300.0
	priorBeta = 700.0
	betaA = distuv.Beta{Alpha: 36 + priorAlpha, Beta: 114 + priorBeta, Src: s}
	betaB = distuv.Beta{Alpha: 50 + priorAlpha, Beta: 100 + priorBeta, Src: s}
	aSamples = betaSamples(betaA, trials)
	bSamples = betaSamples(betaB, trials)
	fmt.Println("15.1 B > A =", bSuper(aSamples, bSamples))

	betaA = distuv.Beta{Alpha: 36 + 30, Beta: 114 + 70, Src: s}
	betaB = distuv.Beta{Alpha: 50 + 20, Beta: 100 + 80, Src: s}
	aSamples = betaSamples(betaA, trials)
	bSamples = betaSamples(betaB, trials)
	fmt.Println("15.2 B > A =", bSuper(aSamples, bSamples))

	ex15_3()
}

func ex15_3() {
	aTrueRate := 0.25
	bTrueRate := 0.3
	numSamples := 0
	super := -1.0
	for super < 0.95 {
		numSamples += 100
		aTrue, aFalse := runif(numSamples/2, aTrueRate)
		bTrue, bFalse := runif(numSamples/2, bTrueRate)
		betaA := distuv.Beta{Alpha: float64(aTrue + 300), Beta: float64(aFalse + 700), Src: s}
		betaB := distuv.Beta{Alpha: float64(bTrue + 300), Beta: float64(bFalse + 700), Src: s}
		aSamples := betaSamples(betaA, trials)
		bSamples := betaSamples(betaB, trials)
		super = bSuper(aSamples, bSamples)
	}
	fmt.Println("15.3 director of marketing samples =", numSamples)

	numSamples = 0
	super = -1.0
	for super < 0.95 {
		numSamples += 100
		aTrue, aFalse := runif(numSamples/2, aTrueRate)
		bTrue, bFalse := runif(numSamples/2, bTrueRate)
		betaA := distuv.Beta{Alpha: float64(aTrue + 30), Beta: float64(aFalse + 70), Src: s}
		betaB := distuv.Beta{Alpha: float64(bTrue + 20), Beta: float64(bFalse + 80), Src: s}
		aSamples := betaSamples(betaA, trials)
		bSamples := betaSamples(betaB, trials)
		super = bSuper(aSamples, bSamples)
	}
	fmt.Println("15.3 lead designer samples =", numSamples)
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

func fig15_2() {
	p, err := plot.New()
	if err != nil {
		log.Fatalln("plot.New()", err)
		panic(err)
	}

	p.Title.Text = "Parameter estimation variants A and B"
	p.X.Label.Text = "Conversion rate"
	p.X.Min = 0.09
	p.X.Max = 0.51
	p.Y.Label.Text = "Density"
	p.Y.Min = -0.5
	p.Y.Max = 12.5
	p.Legend.Top = true
	p.Legend.Left = true

	ts := []plot.Tick{
		{Value: 0.1, Label: "0.1"}, {Value: 0.2, Label: "0.2"},
		{Value: 0.3, Label: "0.3"}, {Value: 0.4, Label: "0.4"}, {Value: 0.5, Label: "0.5"},
	}
	p.X.Tick.Marker = plot.ConstantTicks(ts)
	p.Y.Tick.Marker = plot.ConstantTicks([]plot.Tick{
		{Value: 0, Label: "0"}, {Value: 2, Label: "2"}, {Value: 4, Label: "4"},
		{Value: 6, Label: "6"}, {Value: 8, Label: "8"}, {Value: 10, Label: "10"},
		{Value: 12, Label: "12"},
	})

	grid := plotter.NewGrid()
	grid.Horizontal.Color = color.RGBA{A: 0}
	grid.Vertical.Color = color.RGBA{A: 0}
	p.Add(grid)

	priorAlpha := 3.0
	priorBeta := 10.0 - priorAlpha

	betaA := distuv.Beta{Alpha: 36 + priorAlpha, Beta: 114 + priorBeta}
	lineA := plotter.NewFunction(func(x float64) float64 { return betaA.Prob(x) })
	lineA.Samples = samples
	lineA.XMin = 0.1
	lineA.XMax = 0.5
	lineA.Color = color.RGBA{B: 255, A: 255}
	lineA.Width = 1

	betaB := distuv.Beta{Alpha: 50 + priorAlpha, Beta: 100 + priorBeta}
	lineB := plotter.NewFunction(func(x float64) float64 { return betaB.Prob(x) })
	lineB.Samples = samples
	lineB.XMin = 0.1
	lineB.XMax = 0.5
	lineB.Color = color.RGBA{R: 255, A: 255}
	lineB.Width = 1

	p.Add(lineA, lineB)
	p.Legend.Add("A", lineA)
	p.Legend.Add("B", lineB)

	if err = p.Save(8*vg.Inch, 8*vg.Inch, "15_2.png"); err != nil {
		log.Fatalln("plot.Save()", err)
	}
}

func fig15_3(aSamples, bSamples []float64) {
	v := make(plotter.Values, len(aSamples))
	for i := range v {
		v[i] = bSamples[i] / aSamples[i]
	}
	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Histogram of bSamples/aSamples"
	p.X.Label.Text = "bSamples/aSamples"
	p.X.Min = 0.5
	p.X.Max = 3.5
	ts := []plot.Tick{
		{Value: 0.5, Label: "0.5"}, {Value: 1.0, Label: "1.0"},
		{Value: 1.5, Label: "1.5"}, {Value: 2.0, Label: "2.0"},
		{Value: 2.5, Label: "2.5"}, {Value: 3.0, Label: "3.0"},
		{Value: 3.5, Label: "3.5"},
	}
	p.X.Tick.Marker = plot.ConstantTicks(ts)
	p.Y.Label.Text = "Frequency"

	h, err := plotter.NewHist(v, 15)
	if err != nil {
		panic(err)
	}
	p.Add(h)

	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "15_3.png"); err != nil {
		panic(err)
	}
}

func fig15_4(aSamples, bSamples []float64) {
	var v []float64
	for i := range aSamples {
		v = append(v, bSamples[i]/aSamples[i])
	}
	sort.Float64s(v)

	p, err := plot.New()
	if err != nil {
		log.Fatalln("plot.New()", err)
		panic(err)
	}

	p.Title.Text = "ecdf(bSamples/aSamples)"
	p.X.Label.Text = "Cumulative probability"
	p.X.Min = 0.25
	p.X.Max = 3.75
	p.Y.Label.Text = "Improvement"
	ts := []plot.Tick{
		{Value: 0.5, Label: "0.5"}, {Value: 1.0, Label: "1.0"},
		{Value: 1.5, Label: "1.5"}, {Value: 2.0, Label: "2.0"},
		{Value: 2.5, Label: "2.5"}, {Value: 3.0, Label: "3.0"},
		{Value: 3.5, Label: "3.5"},
	}
	p.X.Tick.Marker = plot.ConstantTicks(ts)
	p.Y.Min = 0.0
	p.Y.Max = 1.0
	p.Y.Tick.Marker = plot.ConstantTicks([]plot.Tick{
		{Value: 0, Label: "0.0"}, {Value: 0.25, Label: "0.25"},
		{Value: 0.5, Label: "0.5"}, {Value: 0.75, Label: "0.75"},
		{Value: 1.0, Label: "1.0"},
		{Value: 0.1, Label: ""}, {Value: 0.2, Label: ""},
		{Value: 0.3, Label: ""}, {Value: 0.4, Label: ""},
		{Value: 0.6, Label: ""}, {Value: 0.7, Label: ""},
		{Value: 0.8, Label: ""}, {Value: 0.9, Label: ""},
	})
	p.Legend.Top = true
	p.Legend.Left = true

	p.Add(plotter.NewGrid())

	line := plotter.NewFunction(func(q float64) float64 { return stat.CDF(q, stat.Empirical, v, nil) })
	line.Samples = samples
	line.XMin = 0.25
	line.XMax = 3.75
	line.Width = 1

	p.Add(line)

	if err = p.Save(8*vg.Inch, 8*vg.Inch, "15_4.png"); err != nil {
		log.Fatalln("plot.Save()", err)
	}
}
