package main

import (
	"log"

	"github.com/zacksfF/Data-Engineering/Data-pipeline/internal/wine"
	"golang.org/x/text/message/pipeline"
)

func main() {
	p := new(pipeline.Pipeline)
	p.Gather("https://archive.ics.uci.edu/ml/machine-learning-databases/iris/iris.data", "data/external/iris.csv").
		Organize("data/external/iris.csv", wine.Organize).
		Evaluate("reports/plot_scatter.png", wine.PlotScatter).
		Validate(nil, wine.PrintStats)
	if err := p.Error(); err != nil {
		log.Fatal(err)
	}
}
