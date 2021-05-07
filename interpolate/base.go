package interpolate

import (
	"fmt"

	"github.com/53jk1/pok2"
)

// Base type provides the base functionality for any interpolation type
type Base struct {
	XYPairs []pok2.CoordinatePair
	X       []float64
	Y       []float64
}

// Fit receives two float64 slices - for the x and y coordinates where x[i] and y[i] represent a coordinate pair in a grid.
// It returns the error if the X and Y sizes do not match.
func (b *Base) Fit(x, y []float64) error {
	if len(x) != len(y) {
		return fmt.Errorf("X and Y sizes do not match")
	}
	b.X = x
	b.Y = y
	b.XYPairs = pok2.SlicesToCoordinatePairs(x, y)
	pok2.SortCoordinatePairs(b.XYPairs)
	return nil
}
