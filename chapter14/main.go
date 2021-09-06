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

var samples = 500

func main() {
	//fig14_1()

	//fig14_3()

	//fig14_4()

	//fig14_5()

	fig14_6()

	beta := distuv.Beta{Alpha: 18 + 7, Beta: 14 + 3}
	fmt.Println("quantile 0.025 =", beta.Quantile(0.025))
	fmt.Println("quantile 0.975 =", beta.Quantile(0.975))

	beta = distuv.Beta{Alpha: 18 + 1000, Beta: 14 + 1000}
	fmt.Println("quantile 0.025 =", beta.Quantile(0.025))
	fmt.Println("quantile 0.975 =", beta.Quantile(0.975))

	beta = distuv.Beta{Alpha: 18 + 70, Beta: 14 + 30}
	fmt.Println("quantile 0.025 =", beta.Quantile(0.025))
	fmt.Println("quantile 0.975 =", beta.Quantile(0.975))
}

func fig14_1() {
	p, err := plot.New()
	if err != nil {
		log.Fatalln("plot.New()", err)
		panic(err)
	}

	p.Title.Text = "Beta(2,3) likelihood for possible converstion rate"
	p.X.Label.Text = "Conversion rate"
	p.X.Min = 0.0
	p.X.Max = 1.0
	p.Y.Label.Text = "Density"
	p.Y.Min = 0
	p.Y.Max = 2
	p.Legend.Top = true
	p.Legend.Left = true
	p.Add(plotter.NewGrid())

	beta := distuv.Beta{Alpha: 2, Beta: 3}
	line := plotter.NewFunction(func(x float64) float64 { return beta.Prob(x) })
	line.Samples = samples
	line.Color = color.RGBA{A: 255}
	line.Width = 1
	line.Dashes = []vg.Length{vg.Points(5), vg.Points(5), vg.Points(10), vg.Points(5)}

	p.Add(line)

	if err = p.Save(8*vg.Inch, 8*vg.Inch, "14_1.png"); err != nil {
		log.Fatalln("plot.Save()", err)
	}
}

func fig14_3() {
	p, err := plot.New()
	if err != nil {
		log.Fatalln("plot.New()", err)
		panic(err)
	}

	p.Title.Text = "Possible priors for email converstion rate"
	p.X.Label.Text = "Conversion rate"
	p.X.Min = 0.0
	p.X.Max = 0.15
	p.Y.Label.Text = "Density"
	p.Y.Min = 0
	p.Y.Max = 45
	p.Legend.Top = true
	p.Legend.Left = false
	p.Add(plotter.NewGrid())

	beta1 := distuv.Beta{Alpha: 1, Beta: 41}
	line1 := plotter.NewFunction(func(x float64) float64 { return beta1.Prob(x) })
	line1.Samples = samples
	line1.Color = color.RGBA{A: 255}
	line1.Width = 1
	p.Legend.Add("Beta(1,41)", line1)

	beta2 := distuv.Beta{Alpha: 2, Beta: 80}
	line2 := plotter.NewFunction(func(x float64) float64 { return beta2.Prob(x) })
	line2.Samples = samples
	line2.Color = color.RGBA{A: 255}
	line2.Width = 1
	line2.Dashes = []vg.Length{vg.Points(3), vg.Points(2)}
	p.Legend.Add("Beta(2,80)", line2)

	beta3 := distuv.Beta{Alpha: 5, Beta: 200}
	line3 := plotter.NewFunction(func(x float64) float64 { return beta3.Prob(x) })
	line3.Samples = samples
	line3.Color = color.RGBA{A: 255}
	line3.Width = 1
	line3.Dashes = []vg.Length{vg.Points(10), vg.Points(5)}
	p.Legend.Add("Beta(5,20)", line3)

	p.Add(line1, line2, line3)

	if err = p.Save(8*vg.Inch, 8*vg.Inch, "14_3.png"); err != nil {
		log.Fatalln("plot.Save()", err)
	}
}

func fig14_4() {
	p, err := plot.New()
	if err != nil {
		log.Fatalln("plot.New()", err)
		panic(err)
	}

	p.Title.Text = "Estimates of conversion rate with and without prior"
	p.X.Label.Text = "Conversion rate"
	p.X.Min = 0.0
	p.X.Max = 1.0
	p.Y.Label.Text = "Density"
	p.Y.Min = 0
	p.Y.Max = 15
	p.Legend.Top = true
	p.Legend.Left = false
	p.Add(plotter.NewGrid())

	beta1 := distuv.Beta{Alpha: 1 + 2, Beta: 40 + 3}
	line1 := plotter.NewFunction(func(x float64) float64 { return beta1.Prob(x) })
	line1.Samples = samples
	line1.Color = color.RGBA{B: 255, A: 255}
	line1.Width = 1

	beta2 := distuv.Beta{Alpha: 2, Beta: 3}
	line2 := plotter.NewFunction(func(x float64) float64 { return beta2.Prob(x) })
	line2.Samples = samples
	line2.Color = color.RGBA{A: 255}
	line2.Width = 1
	line2.Dashes = []vg.Length{vg.Points(3), vg.Points(2)}

	p.Add(line1, line2)
	p.Legend.Add("With Prior", line1)
	p.Legend.Add("No Prior", line2)

	if err = p.Save(8*vg.Inch, 8*vg.Inch, "14_4.png"); err != nil {
		log.Fatalln("plot.Save()", err)
	}
}

func fig14_5() {
	p, err := plot.New()
	if err != nil {
		log.Fatalln("plot.New()", err)
		panic(err)
	}

	p.Title.Text = "Estimates of conversion rate with and without prior"
	p.X.Label.Text = "Conversion rate"
	p.X.Min = 0.0
	p.X.Max = 1.0
	p.Y.Label.Text = "Density"
	p.Y.Min = 0
	p.Y.Max = 15
	p.Legend.Top = true
	p.Legend.Left = false
	p.Add(plotter.NewGrid())

	beta1 := distuv.Beta{Alpha: 1 + 25, Beta: 40 + 75}
	line1 := plotter.NewFunction(func(x float64) float64 { return beta1.Prob(x) })
	line1.Samples = samples
	line1.Color = color.RGBA{B: 255, A: 255}
	line1.Width = 1

	beta2 := distuv.Beta{Alpha: 25, Beta: 75}
	line2 := plotter.NewFunction(func(x float64) float64 { return beta2.Prob(x) })
	line2.Samples = samples
	line2.Color = color.RGBA{A: 255}
	line2.Width = 1
	line2.Dashes = []vg.Length{vg.Points(3), vg.Points(2)}

	p.Add(line1, line2)
	p.Legend.Add("With Prior", line1)
	p.Legend.Add("No Prior", line2)

	if err = p.Save(8*vg.Inch, 8*vg.Inch, "14_5.png"); err != nil {
		log.Fatalln("plot.Save()", err)
	}
}

type myTicks struct {
	steps  int
	major  int
	format string
}

func (t myTicks) Ticks(min, max float64) []plot.Tick {
	step := (max - min) / float64(t.steps)
	tks := make([]plot.Tick, t.steps+1)
	for i := 0; i <= t.steps; i++ {
		tks[i].Value = min + step*float64(i)
		if i%t.major == 0 {
			tks[i].Label = fmt.Sprintf(t.format, tks[i].Value)
		}
	}
	return tks[:]
}

func fig14_6() {
	p, err := plot.New()
	if err != nil {
		log.Fatalln("plot.New()", err)
		panic(err)
	}

	p.Title.Text = "Estimates of conversion rate with more data with and without prior"
	p.X.Label.Text = "Conversion rate"
	p.X.Min = 0.0
	p.X.Max = 1.0
	p.Y.Label.Text = "Density"
	p.Y.Min = 0
	p.Y.Max = 25
	p.Legend.Top = true
	p.Legend.Left = false

	ts := []plot.Tick{
		{Value: 0, Label: "0"}, {Value: 0.1, Label: "0.1"}, {Value: 0.2, Label: "0.2"},
		{Value: 0.25, Label: "0.25"},
		{Value: 0.3, Label: "0.3"}, {Value: 0.4, Label: "0.4"}, {Value: 0.5, Label: "0.5"},
		{Value: 0.6, Label: ""}, {Value: 0.7, Label: ""}, {Value: 0.8, Label: ""},
		{Value: 0.9, Label: ""}, {Value: 1, Label: "1"},
		{Value: 0.05, Label: ""}, {Value: 0.15, Label: ""}, {Value: 0.35, Label: ""}, {Value: 0.45, Label: ""},
	}
	p.X.Tick.Marker = plot.ConstantTicks(ts)
	p.Y.Tick.Marker = myTicks{25, 5, "%.0f"}
	p.Y.Tick.Color = color.RGBA{G: 255, A: 255}
	p.Y.Color = color.RGBA{G: 128, A: 255}

	grid := plotter.NewGrid()
	grid.Horizontal.Color = color.RGBA{R: 128, B: 128, A: 255}
	grid.Horizontal.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	p.Add(grid)

	priorAlpha := 86.0
	priorBeta := 300.0 - priorAlpha

	beta1 := distuv.Beta{Alpha: priorAlpha, Beta: priorBeta}
	line1 := plotter.NewFunction(func(x float64) float64 { return beta1.Prob(x) })
	line1.Samples = samples
	line1.XMin = 0.0
	line1.XMax = 1.0
	line1.Color = color.RGBA{A: 255}
	line1.Width = 1
	line1.Dashes = []vg.Length{vg.Points(3), vg.Points(2)}

	beta2 := distuv.Beta{Alpha: 1 + priorAlpha, Beta: 40 + priorBeta}
	line2 := plotter.NewFunction(func(x float64) float64 { return beta2.Prob(x) })
	line2.Samples = samples
	line2.XMin = 0.0
	line2.XMax = 1.0
	line2.Color = color.RGBA{B: 255, A: 255}
	line2.Width = 1

	beta3 := distuv.Beta{Alpha: 5 + priorAlpha, Beta: 200 + priorBeta}
	line3 := plotter.NewFunction(func(x float64) float64 { return beta3.Prob(x) })
	line3.Samples = samples
	line3.XMin = 0.0
	line3.XMax = 0.5
	line3.Color = color.RGBA{R: 255, A: 255}
	line3.Width = 1

	p.Add(line1, line2, line3)
	p.Legend.Add("No Prior", line1)
	p.Legend.Add("With Prior", line2)
	p.Legend.Add("With Prior2", line3)

	if err = p.Save(8*vg.Inch, 8*vg.Inch, "14_6.png"); err != nil {
		log.Fatalln("plot.Save()", err)
	}
}
