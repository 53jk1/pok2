package main_test

import (
	"fmt"
	"testing"

	"github.com/53jk1/pok2"
	"github.com/stretchr/testify/assert"
)

func TestVectorDim(t *testing.T) {
	cases := map[string]struct {
		vector         pok2.Vector
		expectedResult int
	}{
		"basic test": {
			vector:         pok2.Vector{1, 2, 3},
			expectedResult: 3,
		},
		"empty vector": {
			vector:         pok2.Vector{},
			expectedResult: 0,
		},
		"single element vector": {
			vector:         pok2.Vector{1},
			expectedResult: 1,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, c.expectedResult, c.vector.Dim())
		})
	}

}

func TestVectorAreDimsEqual(t *testing.T) {
	cases := map[string]struct {
		vector1        pok2.Vector
		vector2        pok2.Vector
		expectedResult bool
	}{
		"basic equality test": {
			vector1:        pok2.Vector{1, 2, 3},
			vector2:        pok2.Vector{3, 1, 0},
			expectedResult: true,
		},
		"different dimensions": {
			vector1:        pok2.Vector{1, 2},
			vector2:        pok2.Vector{1, 2, 3},
			expectedResult: false,
		},
		"testing with empty vector": {
			vector1:        pok2.Vector{1},
			vector2:        pok2.Vector{},
			expectedResult: false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, c.expectedResult, c.vector1.AreDimsEqual(c.vector2))
		})
	}
}

func TestAddVectors(t *testing.T) {
	cases := map[string]struct {
		vector1        pok2.Vector
		vector2        pok2.Vector
		expectedResult pok2.Vector
		expectedError  error
	}{
		"add vectors basic test": {
			vector1:        pok2.Vector{1, 2, 3},
			vector2:        pok2.Vector{3, 1, 0},
			expectedResult: pok2.Vector{4, 3, 3},
			expectedError:  nil,
		},
		"add with wrong dimensions": {
			vector1:        pok2.Vector{1, 2},
			vector2:        pok2.Vector{1, 2, 3},
			expectedResult: nil,
			expectedError:  fmt.Errorf("Dimensions must match"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := c.vector1.Add(c.vector2)

			assert.Equal(t, c.expectedResult, result)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestSubtractVectors(t *testing.T) {
	cases := map[string]struct {
		vector1        pok2.Vector
		vector2        pok2.Vector
		expectedResult pok2.Vector
		expectedError  error
	}{
		"subtract vectors basic test": {
			vector1:        pok2.Vector{1, 2, 3},
			vector2:        pok2.Vector{3, 1, 0},
			expectedResult: pok2.Vector{-2, 1, 3},
			expectedError:  nil,
		},
		"subtract with wrong dimensions": {
			vector1:        pok2.Vector{1, 2},
			vector2:        pok2.Vector{1, 2, 3},
			expectedResult: nil,
			expectedError:  fmt.Errorf("Dimensions must match"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := c.vector1.Subtract(c.vector2)

			assert.Equal(t, c.expectedResult, result)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestVectorDotProduct(t *testing.T) {
	cases := map[string]struct {
		vector1        pok2.Vector
		vector2        pok2.Vector
		expectedResult float64
		expectedError  error
	}{
		"basic dot product test": {
			vector1:        pok2.Vector{1, 2, 3},
			vector2:        pok2.Vector{4, 5, 6},
			expectedResult: 32,
			expectedError:  nil,
		},
		"dot product with wrong dimensions": {
			vector1:        pok2.Vector{1, 2},
			vector2:        pok2.Vector{1, 2, 3},
			expectedResult: 0,
			expectedError:  fmt.Errorf("Dimensions must match"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := c.vector1.Dot(c.vector2)

			assert.Equal(t, c.expectedResult, result)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestMultiplyVectorByScalar(t *testing.T) {
	cases := map[string]struct {
		vector         pok2.Vector
		scalar         float64
		expectedResult pok2.Vector
	}{
		"basic multiply vector by scalar": {
			vector:         pok2.Vector{1, 2, 3},
			scalar:         5,
			expectedResult: pok2.Vector{5, 10, 15},
		},
		"multiply vector by 0": {
			vector:         pok2.Vector{1, 2},
			scalar:         0,
			expectedResult: pok2.Vector{0, 0},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := c.vector.MultiplyByScalar(c.scalar)
			assert.Equal(t, c.expectedResult, result)
		})
	}

}

func TestDivideVectorByScalar(t *testing.T) {
	cases := map[string]struct {
		vector         pok2.Vector
		scalar         float64
		expectedResult pok2.Vector
		expectedError  error
	}{
		"basic vector division by scalar": {
			vector:         pok2.Vector{1, 2, 3},
			scalar:         5.0,
			expectedResult: pok2.Vector{0.2, 0.4, 0.6},
			expectedError:  nil,
		},
		"vector division by 0": {
			vector:         pok2.Vector{1, 2},
			scalar:         0,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Cannot divide by zero"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := c.vector.DivideByScalar(c.scalar)
			assert.Equal(t, c.expectedResult, result)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestVectorPower(t *testing.T) {
	cases := map[string]struct {
		vector         pok2.Vector
		power          float64
		expectedResult pok2.Vector
	}{
		"vector elements squared": {
			vector:         pok2.Vector{1, 2, 3},
			power:          2,
			expectedResult: pok2.Vector{1, 4, 9},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := c.vector.Power(c.power)
			assert.Equal(t, c.expectedResult, result)
		})
	}
}

func TestVectorIsSimilar(t *testing.T) {
	cases := map[string]struct {
		vector1        pok2.Vector
		vector2        pok2.Vector
		tolerance      float64
		expectedResult bool
	}{
		"test non-similar vectors": {
			vector1:        pok2.Vector{1.2, 2.5},
			vector2:        pok2.Vector{1, 2},
			tolerance:      0.01,
			expectedResult: false,
		},
		"test similar vectors": {
			vector1:        pok2.Vector{1.000000001, 2.0000000001},
			vector2:        pok2.Vector{1, 2},
			tolerance:      0.1,
			expectedResult: true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := c.vector1.IsSimilar(c.vector2, c.tolerance)
			assert.Equal(t, c.expectedResult, result)
		})
	}
}
