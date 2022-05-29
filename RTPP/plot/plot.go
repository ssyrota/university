package plot

import (
	"github.com/Arafatk/glot"
	qh "github.com/sergeycrisp/university/RTPP/quick_hull"
)

// Draw plot to the file
func MakePlot(shellAllPoints, shellCornedPoints []qh.Point, output string) error {
	// Set up plot
	dimensions := 2
	persist := false
	debug := false
	plot, err := glot.NewPlot(dimensions, persist, debug)
	if err != nil {
		return err
	}

	// Draw all points
	allPoints := [][]float64{{}, {}}
	for _, p := range shellAllPoints {
		allPoints[0] = append(allPoints[0], p.X)
		allPoints[1] = append(allPoints[1], p.Y)
	}
	plot.AddPointGroup("All points", "points", allPoints)

	// Draw corner shell points line
	cornerPoints := [][]float64{{}, {}}
	for _, p := range shellCornedPoints {
		cornerPoints[0] = append(cornerPoints[0], p.X)
		cornerPoints[1] = append(cornerPoints[1], p.Y)
	}
	cornerPoints[0] = append(cornerPoints[0], cornerPoints[0][0])
	cornerPoints[1] = append(cornerPoints[1], cornerPoints[1][0])
	plot.AddPointGroup("Corner points", "lines", cornerPoints)

	// Set min and max values for axes
	_, maxX := qh.MinMax(allPoints[0])
	plot.SetXrange(-int(maxX*0.1), int(maxX*1.1))
	_, maxY := qh.MinMax(allPoints[1])
	plot.SetYrange(-int(maxX*0.1), int(maxY*1.1))

	// Save png
	err = plot.SavePlot(output)
	if err != nil {
		return err
	}

	return nil
}
