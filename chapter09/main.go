package main

import (
	"fmt"
	"image/color"
	"log"

	"gonum.org/v1/gonum/stat/distuv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	fmt.Println("vim-go")

	//betaPNG()

	// 9.1
	coin := distuv.Beta{Alpha: 6, Beta: 1}
	fmt.Println("9.1 =", coin.CDF(0.6)-coin.CDF(0.4))

	// 9.2
	prior := 10.0
	for ; ; prior += 1.0 {
		coin = distuv.Beta{Alpha: 6 + prior, Beta: 1 + prior}
		prob := coin.CDF(0.6) - coin.CDF(0.4)
		fmt.Println("prior =", prior, "prob. =", prob)
		if prob > 0.95 {
			break
		}
	}
	fmt.Println("9.2 find the prior that coin is fair", prior)

	// 9.3
	moreHeads := 1.0
	for ; ; moreHeads += 1.0 {
		coin = distuv.Beta{Alpha: 6 + prior + moreHeads, Beta: 1 + prior}
		prob := coin.CDF(0.6) - coin.CDF(0.4)
		fmt.Println("more heads =", moreHeads, "prob. =", prob)
		if prob < 0.5 {
			break
		}
	}
	fmt.Println("9.3 find more heads that coin is unfair", moreHeads)
}

func betaPng() {
	p, err := plot.New()
	if err != nil {
		log.Fatalln("plot.New()", err)
		panic(err)
	}

	p.Title.Text = "Beta"
	p.X.Label.Text = "Probability of success"
	p.X.Min = 0.7
	p.X.Max = 0.75
	p.Y.Label.Text = "Density"
	p.Y.Min = 0
	p.Y.Max = 160
	p.Legend.Top = true
	p.Legend.Left = true

	han := distuv.Beta{Alpha: 20002, Beta: 7401}
	line := plotter.NewFunction(func(x float64) float64 { return han.Prob(x) })
	line.Color = color.RGBA{G: 255, A: 255}
	line.Width = 3
	p.Legend.Add("Liner Model", line)

	p.Add(line)

	if err = p.Save(8*vg.Inch, 8*vg.Inch, "beta.png"); err != nil {
		log.Fatalln("plot.Save()", err)
	}
}
