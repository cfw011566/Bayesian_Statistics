package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"gonum.org/v1/gonum/stat/distuv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

var samples = 500

func main() {
	binom := distuv.Binomial{N: 100, P: 0.5}
	fmt.Println("(24,100,0.5) =", binom.CDF(24))
	f := math.Pow(0.5, 24) * math.Pow(0.5, 76) / math.Pow(0.05, 24) / math.Pow(0.95, 76)
	fmt.Println(f)
	f = math.Pow(0.2, 24) * math.Pow(0.8, 76) / math.Pow(0.5, 24) / math.Pow(0.5, 76)
	fmt.Println(f)

	//fig19_1()
	//fig19_5()
	ex19_1()
}

func fig19_1() {
	p, err := plot.New()
	if err != nil {
		log.Fatalln("plot.New()", err)
		panic(err)
	}

	p.Title.Text = ""
	p.X.Label.Text = "Hypotheses"
	p.X.Min = 0.0
	p.X.Max = 1.0
	p.Y.Label.Text = "bfs"
	p.Y.Min = 0
	p.Y.Max = 1500000
	//p.Legend.Top = true
	//p.Legend.Left = true
	p.Add(plotter.NewGrid())

	bayesFactor := func(t, b float64) float64 {
		top := math.Pow(t, 24) * math.Pow(1-t, 76)
		bottom := math.Pow(b, 24) * math.Pow(1-b, 76)
		return top / bottom
	}
	line := plotter.NewFunction(func(x float64) float64 {
		return bayesFactor(x, 0.5)
	})
	line.Samples = samples
	line.Width = 1

	p.Add(line)

	if err = p.Save(8*vg.Inch, 8*vg.Inch, "19_1.png"); err != nil {
		log.Fatalln("plot.Save()", err)
	}
}

func fig19_5() {
	p, err := plot.New()
	if err != nil {
		log.Fatalln("plot.New()", err)
		panic(err)
	}

	p.Title.Text = "Beta(24,76) for our hypotheses"
	p.X.Label.Text = "Hypotheses"
	p.X.Min = 0.0
	p.X.Max = 1.0
	p.Y.Label.Text = "Density"
	p.Y.Min = 0
	p.Y.Max = 10
	p.Add(plotter.NewGrid())

	beta := distuv.Beta{Alpha: 24, Beta: 76}
	line := plotter.NewFunction(func(x float64) float64 { return beta.Prob(x) })
	line.Samples = samples
	line.Color = color.RGBA{A: 255}
	line.Width = 1

	p.Add(line)

	if err = p.Save(8*vg.Inch, 8*vg.Inch, "19_5.png"); err != nil {
		log.Fatalln("plot.Save()", err)
	}
}

func ex19_1() {
	p, err := plot.New()
	if err != nil {
		log.Fatalln("plot.New()", err)
		panic(err)
	}

	p.Title.Text = ""
	p.X.Label.Text = "Hypotheses"
	p.X.Min = 0.0
	p.X.Max = 1.0
	p.Y.Label.Text = "bfs"
	p.Y.Min = 0
	p.Y.Max = 10
	p.Legend.Top = true
	//p.Legend.Left = true
	p.Add(plotter.NewGrid())

	bayesFactor := func(t, b float64) float64 {
		top := math.Pow(t, 24) * math.Pow(1-t, 76)
		bottom := math.Pow(b, 24) * math.Pow(1-b, 76)
		return top / bottom
	}

	n := 200
	pts1 := make(plotter.XYs, n)
	pts2 := make(plotter.XYs, n)
	sum1 := 0.0
	sum2 := 0.0
	steps := 1.0 / float64(n)
	for i := 0; i < n; i++ {
		x := float64(i) * steps
		y1 := bayesFactor(x, 0.5)
		y2 := bayesFactor(x, 0.24)
		pts1[i].X = x
		pts1[i].Y = y1
		pts2[i].X = x
		pts2[i].Y = y2
		sum1 += y1
		sum2 += y2
	}
	for i := 0; i < n; i++ {
		pts1[i].Y = pts1[i].Y / sum1 * float64(n)
		pts2[i].Y = pts2[i].Y / sum2 * float64(n)
	}

	line, err := plotter.NewLine(pts1)
	if err != nil {
		log.Panic(err)
	}
	line.Color = color.RGBA{A: 255}
	p.Legend.Add("H1: P(prize) = 0.5", line)

	s, err := plotter.NewScatter(pts2)
	if err != nil {
		log.Panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{A: 255}
	s.GlyphStyle.Radius = vg.Points(3)
	s.GlyphStyle.Shape = draw.RingGlyph{}
	p.Legend.Add("H1: P(prize) = 0.24", s)

	p.Add(line, s)

	if err = p.Save(8*vg.Inch, 8*vg.Inch, "ex19_1.png"); err != nil {
		log.Fatalln("plot.Save()", err)
	}
}
