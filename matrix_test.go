package pok2_test

import (
	"fmt"
	"testing"

	"github.com/53jk1/pok2"

	"github.com/stretchr/testify/assert"
)

func TestCompareMatrices(t *testing.T) {

	cases := map[string]struct {
		matrix1 pok2.Matrix
		matrix2 pok2.Matrix
		isEqual bool
	}{
		"basic matrix comparison": {
			matrix1: pok2.Matrix{
				{1, 2},
				{3, 4},
			},
			matrix2: pok2.Matrix{
				{1, 2},
				{3, 4},
			},
			isEqual: true,
		},
		"wrong dimensions matrix comparison": {
			matrix1: pok2.Matrix{
				{1, 2},
			},
			matrix2: pok2.Matrix{
				{1, 2},
				{3, 4},
			},
			isEqual: false,
		},
		"different matrices comparison": {
			matrix1: pok2.Matrix{
				{1, 2},
				{3, 5},
			},
			matrix2: pok2.Matrix{
				{1, 2},
				{3, 4},
			},
			isEqual: false,
		},
		"comparing matrix to nil": {
			matrix1: pok2.Matrix{
				{1, 2},
				{3, 5},
			},
			matrix2: nil,
			isEqual: false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			isEqual := c.matrix1.IsEqual(c.matrix2)
			assert.Equal(t, c.isEqual, isEqual)
		})
	}
}

func TestMatrixInsertCol(t *testing.T) {
	cases := map[string]struct {
		matrix         pok2.Matrix
		column         pok2.Vector
		index          int
		expectedResult pok2.Matrix
		expectedError  error
	}{
		"adding column at 0th index": {
			matrix: pok2.Matrix{
				{1, 2},
				{3, 4},
			},
			column: pok2.Vector{1, 1},
			index:  0,
			expectedResult: pok2.Matrix{
				{1, 1, 2},
				{1, 3, 4},
			},
			expectedError: nil,
		},
		"adding column at 1st index": {
			matrix: pok2.Matrix{
				{1, 2},
				{3, 4},
			},
			column: pok2.Vector{1, 1},
			index:  1,
			expectedResult: pok2.Matrix{
				{1, 1, 2},
				{3, 1, 4},
			},
			expectedError: nil,
		},
		"adding column at 2nd index": {
			matrix: pok2.Matrix{
				{1, 2},
				{3, 4},
			},
			column: pok2.Vector{1, 1},
			index:  2,
			expectedResult: pok2.Matrix{
				{1, 2, 1},
				{3, 4, 1},
			},
			expectedError: nil,
		},
		"adding column with wrong dimensions": {
			matrix: pok2.Matrix{
				{1, 2},
				{3, 4},
			},
			column:         pok2.Vector{1, 1, 4},
			index:          0,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Column dimensions must match"),
		},
		"adding column at incorrect index": {
			matrix: pok2.Matrix{
				{1, 2},
				{3, 4},
			},
			column:         pok2.Vector{1, 1},
			index:          -1,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Index cannot be less than 0"),
		},
		"adding column at index which is too large": {
			matrix: pok2.Matrix{
				{1, 2},
				{3, 4},
			},
			column:         pok2.Vector{1, 1},
			index:          3,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Index cannot be greater than number of columns + 1"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := c.matrix.InsertCol(c.index, c.column)
			assert.Equal(t, result, c.expectedResult)
			assert.Equal(t, err, c.expectedError)
		})
	}
}

func TestAddMatrices(t *testing.T) {
	cases := map[string]struct {
		matrix1        pok2.Matrix
		matrix2        pok2.Matrix
		expectedResult pok2.Matrix
		expectedError  error
	}{
		"basic matrix addition": {
			matrix1: pok2.Matrix{
				{1, 2},
				{3, 4},
			},
			matrix2: pok2.Matrix{
				{4, 3},
				{2, 1},
			},
			expectedResult: pok2.Matrix{
				{5, 5},
				{5, 5},
			},
			expectedError: nil,
		},
		// Wrong dimensions
		// {
		// 	matrix1: pok2.Matrix{
		// 		{1, 2},
		// 		{3, 4},
		// 	},
		// 	matrix2: pok2.Matrix{
		// 		{4, 3},
		// 	},
		// 	result:        nil,
		// 	expectedError: fmt.Errorf("Matrix dimensions must match"),
		// },
		// Adding two nils
		"adding two nils": {
			matrix1:        nil,
			matrix2:        nil,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Matrices cannot be nil"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			additionResult, err := c.matrix1.Add(c.matrix2)
			assert.Equal(t, additionResult, c.expectedResult)
			assert.Equal(t, err, c.expectedError)
		})
	}
}

func TestSubtractMatrices(t *testing.T) {
	cases := map[string]struct {
		matrix1        pok2.Matrix
		matrix2        pok2.Matrix
		expectedResult pok2.Matrix
		expectedError  error
	}{
		"basic matrix subtraction": {
			matrix1: pok2.Matrix{
				{10, 5},
				{3, 1},
			},
			matrix2: pok2.Matrix{
				{1, 1},
				{1, 1},
			},
			expectedResult: pok2.Matrix{
				{9, 4},
				{2, 0},
			},
			expectedError: nil,
		},
		"matrix subtraction with negative result": {
			matrix1: pok2.Matrix{
				{3, 2},
				{3, 1},
			},
			matrix2: pok2.Matrix{
				{4, 3},
				{4, 2},
			},
			expectedResult: pok2.Matrix{
				{-1, -1},
				{-1, -1},
			},
			expectedError: nil,
		},
		"matrix subtraction with wrong dimensions": {
			matrix1: pok2.Matrix{
				{1, 2},
				{3, 4},
			},
			matrix2: pok2.Matrix{
				{4, 3},
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("Matrix dimensions must match"),
		},
		"matrix subtraction with two nils": {
			matrix1:        nil,
			matrix2:        nil,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Matrices cannot be nil"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			additionResult, err := c.matrix1.Subtract(c.matrix2)
			assert.Equal(t, additionResult, c.expectedResult)
			assert.Equal(t, err, c.expectedError)
		})
	}
}

func TestCol(t *testing.T) {
	cases := map[string]struct {
		matrix         pok2.Matrix
		i              int
		expectedResult pok2.Vector
		expectedError  error
	}{
		"getting column at index 1": {
			matrix: pok2.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			i:              1,
			expectedResult: pok2.Vector{2, 5},
			expectedError:  nil,
		},
		"getting column at wrong index": {
			matrix: pok2.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			i:              -5,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Index cannot be negative"),
		},
		"getting column at index which is too large": {
			matrix: pok2.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			i:              5,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Index cannot be greater than the length"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			column, err := c.matrix.Col(c.i)
			assert.Equal(t, column, c.expectedResult)
			assert.Equal(t, err, c.expectedError)
		})
	}
}

func TestRow(t *testing.T) {
	cases := map[string]struct {
		matrix         pok2.Matrix
		expectedResult pok2.Vector
		i              int
		expectedError  error
	}{
		"getting the row at index 1": {
			matrix: pok2.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			i:              1,
			expectedResult: pok2.Vector{4, 5, 6},
			expectedError:  nil,
		},
		"getting the row at wrong index": {
			matrix: pok2.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			i:              -5,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Index cannot be negative"),
		},
		"getting the row at index which is too large": {
			matrix: pok2.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			i:              5,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Index cannot be greater than the length"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			column, err := c.matrix.Row(c.i)
			assert.Equal(t, column, c.expectedResult)
			assert.Equal(t, err, c.expectedError)
		})
	}

}

func TestTransposeMatrix(t *testing.T) {
	cases := map[string]struct {
		matrix         pok2.Matrix
		expectedResult pok2.Matrix
		expectedError  error
	}{
		"basic matrix transpose": {
			matrix: pok2.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			expectedResult: pok2.Matrix{
				{1, 4},
				{2, 5},
				{3, 6},
			},
			expectedError: nil,
		},
		"second basic matrix transpose": {
			matrix: pok2.Matrix{
				{1, 4},
				{2, 5},
				{3, 6},
			},
			expectedResult: pok2.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			expectedError: nil,
		},
		"transposing one-dimensional matrix": {
			matrix: pok2.Matrix{
				{1, 4},
			},
			expectedResult: pok2.Matrix{
				{1},
				{4},
			},
			expectedError: nil,
		},
		// Inconsistent dimensions
		// {
		// 	matrix: pok2.Matrix{
		// 		{1, 4},
		// 		{2},
		// 	},
		// 	expectedResult:    nil,
		// 	expectedError: fmt.Errorf("Inconsistent dimensions"),
		// },
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			transposed, err := c.matrix.Transpose()
			assert.Equal(t, transposed, c.expectedResult)
			assert.Equal(t, err, c.expectedError)
		})
	}

}

func TestMatrixMultiplication(t *testing.T) {
	cases := map[string]struct {
		matrix1        pok2.Matrix
		matrix2        pok2.Matrix
		expectedResult pok2.Matrix
		expectedError  error
	}{
		"basic matrix multiplication": {
			matrix1: pok2.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			matrix2: pok2.Matrix{
				{1, 1},
				{2, 3},
				{5, 2},
			},
			expectedResult: pok2.Matrix{
				{20, 13},
				{44, 31},
			},
			expectedError: nil,
		},
		"second matrix multiplication": {
			matrix1: pok2.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			matrix2: pok2.Matrix{
				{1, 4},
				{2, 5},
				{3, 6},
			},
			expectedResult: pok2.Matrix{
				{14, 32},
				{32, 77},
			},
			expectedError: nil,
		},
		"multiplying matrices with wrong dimensions": {
			matrix1: pok2.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			matrix2: pok2.Matrix{
				{1, 4},
				{2, 5},
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("The number of columns of the 1st matrix must equal the number of rows of the 2nd matrix"),
		},
		"multiplying matrix with identity matrix": {
			matrix1: pok2.Matrix{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			matrix2: pok2.Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expectedResult: pok2.Matrix{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expectedError: nil,
		},
		"multiplying matrices with different but correct dimensions": {
			matrix1: pok2.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			matrix2: pok2.Matrix{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			},
			expectedResult: pok2.Matrix{
				{6, 12, 18},
				{15, 30, 45},
			},
			expectedError: nil,
		},
		"multiplying 1D matrix with 2D one": {
			matrix1: pok2.Matrix{
				{3, 4, 2},
			},
			matrix2: pok2.Matrix{
				{13, 9, 7, 15},
				{8, 7, 4, 6},
				{6, 4, 0, 3},
			},
			expectedResult: pok2.Matrix{
				{83, 63, 37, 75},
			},
			expectedError: nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			multiplied, err := c.matrix1.MultiplyBy(c.matrix2)
			assert.Equal(t, multiplied, c.expectedResult)
			assert.Equal(t, err, c.expectedError)
		})
	}
}

func TestMatrixLeftDivide(t *testing.T) {
	cases := map[string]struct {
		matrix1        pok2.Matrix
		matrix2        pok2.Matrix
		expectedResult pok2.Matrix
		expectedError  error
	}{
		"basic matrix left divide": {
			matrix1: pok2.Matrix{
				{2},
				{4},
			},
			matrix2: pok2.Matrix{
				{4},
				{4},
			},
			expectedResult: pok2.Matrix{
				{1.2},
			},
			expectedError: nil,
		},
		"second matrix left divide": {
			matrix1: pok2.Matrix{
				{1, 2},
				{2, 2},
			},
			matrix2: pok2.Matrix{
				{3, 2},
				{1, 1},
			},
			expectedResult: pok2.Matrix{
				{-2, -1},
				{2.5, 1.5},
			},
			expectedError: nil,
		},
		"left divide with wrong dimensions": {
			matrix1: pok2.Matrix{
				{1, 2},
				{2, 2},
			},
			matrix2: pok2.Matrix{
				{3, 2},
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("The number of columns of the 1st matrix must equal the number of rows of the 2nd matrix"),
		},
		"left divide - singular matrix": {
			matrix1: pok2.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			matrix2: pok2.Matrix{
				{1, 1},
				{1, 1},
				{1, 1},
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("Matrix is singular"),
		},
		"left divide with ones column": {
			matrix1: pok2.Matrix{
				{1, 1.3},
				{1, 2.1},
				{1, 3.7},
				{1, 4.2},
			},
			matrix2: pok2.Matrix{
				{2.2},
				{5.8},
				{10.2},
				{11.8},
			},
			expectedResult: pok2.Matrix{
				{-1.5225601452564645},
				{3.1938266000907847},
			},
			expectedError: nil,
		},
		"second left divide with ones column": {
			matrix1: pok2.Matrix{
				{1, 0.3},
				{1, 0.8},
				{1, 1.2},
				{1, 1.7},
				{1, 2.4},
				{1, 3.1},
				{1, 3.8},
				{1, 4.5},
				{1, 5.1},
				{1, 5.8},
				{1, 6.5},
			},
			matrix2: pok2.Matrix{
				{8.61},
				{7.94},
				{7.55},
				{6.85},
				{6.11},
				{5.17},
				{4.19},
				{3.41},
				{2.63},
				{1.77},
				{0.89},
			},
			expectedResult: pok2.Matrix{
				{8.99987709451432},
				{-1.246552501126634},
			},
			expectedError: nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			leftDivided, err := c.matrix1.LeftDivide(c.matrix2)

			isSimilar := leftDivided.IsSimilar(c.expectedResult, 1e-4)
			assert.Equal(t, true, isSimilar)
			assert.Equal(t, err, c.expectedError)
		})
	}
}

func TestMatrixInverse(t *testing.T) {
	cases := map[string]struct {
		matrix         pok2.Matrix
		expectedResult pok2.Matrix
		expectedError  error
	}{
		"simple matrix inverse": {
			matrix: pok2.Matrix{
				{4, 7},
				{2, 6},
			},
			expectedResult: pok2.Matrix{
				{0.6, -0.7},
				{-0.2, 0.4},
			},
			expectedError: nil,
		},
		"inverting non-square matrix": {
			matrix: pok2.Matrix{
				{4, 7},
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("Cannot invert non-square Matrix"),
		},
		"inverting singular matrix": {
			matrix: pok2.Matrix{
				{2, 4},
				{6, 12},
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("Matrix is singular"),
		},
		"second simple matrix inverse": {
			matrix: pok2.Matrix{
				{3, 0, 2},
				{2, 0, -2},
				{0, 1, 1},
			},
			expectedResult: pok2.Matrix{
				{0.2, 0.2, 0},
				{-0.2, 0.3, 1},
				{0.2, -0.3, 0},
			},
			expectedError: nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			inverted, err := c.matrix.Invert()
			isSimilar := inverted.IsSimilar(c.expectedResult, 1e-10)
			assert.Equal(t, true, isSimilar)
			assert.Equal(t, err, c.expectedError)
		})
	}
}
func TestMatrixIsSimilar(t *testing.T) {
	cases := map[string]struct {
		matrix1        pok2.Matrix
		matrix2        pok2.Matrix
		expectedResult bool
	}{
		"basic matrix similarity test": {
			matrix1: pok2.Matrix{
				{1, 2},
				{3, 4},
			},
			matrix2: pok2.Matrix{
				{1.000000001, 2.0000000001},
				{3, 4},
			},
			expectedResult: true,
		},
		"matrix similarity with wrong dimensions": {
			matrix1: pok2.Matrix{
				{1, 2},
			},
			matrix2: pok2.Matrix{
				{1, 2},
				{3, 4},
			},
			expectedResult: false,
		},
		"matrix similarity with non-similar matrices": {
			matrix1: pok2.Matrix{
				{1.2, 2.5},
				{3.2, 5.4},
			},
			matrix2: pok2.Matrix{
				{1, 2},
				{3, 4},
			},
			expectedResult: false,
		},
		"passing nil as a second matrix": {
			matrix1: pok2.Matrix{
				{1, 2},
				{3, 5},
			},
			matrix2:        nil,
			expectedResult: false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			isSimilar := c.matrix1.IsSimilar(c.matrix2, 0.1)
			assert.Equal(t, c.expectedResult, isSimilar)
		})
	}
}
