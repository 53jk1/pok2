package lagrange

import (
	"fmt"
)

// Lagrange zapewnia podstawową funkcjonalność interpolacji lagrange.
// Biorąc pod uwagę wycinki X i Y float64, można oszacować wartość funkcji w żądanym punkcie.
type Lagrange struct {
	interpolate.Base
}

// New zwraca nowy obiekt Lagrange'a.
func New() *Lagrange {
	lg := &Lagrange{}
	return lg
}

func (lg *Lagrange) Interpolate(val float64) float64 {
	var est float64

	for i := 0; i < len(lg.X); i++ {
		prod := lg.Y[i]
		for j := 0; j < len(lg.X); j++ {
			if i != j {
				prod = prod * (val - lg.X[j]) / (lg.X[i] - lg.X[j])
			}
		}
		est += prod
	}

	return est
}

func (lg *Lagrange) Validate(val float64) error {

	for i := 0; i < len(lg.X); i++ {
		for j := 0; j < len(lg.X); j++ {
			if i != j {
				if lg.X[i]-lg.X[j] == 0 {
					return fmt.Errorf("Istnieją co najmniej 2 takie same wartości X. Spowoduje to dzielenie przez zero w interpolacji Lagrange'a")
				}
			}
		}
	}

	if val < lg.XYPairs[0].X {
		return fmt.Errorf("Wartość do interpolacji jest zbyt mała i nie mieści się w zakresie")
	}

	if val > lg.XYPairs[len(lg.XYPairs)-1].X {
		return fmt.Errorf("Wartość do interpolacji jest zbyt duża i nie mieści się w zakresie")
	}

	return nil
}
