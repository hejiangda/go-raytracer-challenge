package raytracer

import (
	"errors"
)

type Matrix struct {
	Data [][]float64
	Dim  int
}

func NewMatrix(dim int, vals [][]float64) *Matrix {
	m := new(Matrix)
	m.Dim = dim
	m.Data = make([][]float64, dim)
	for i := range m.Data {
		m.Data[i] = make([]float64, dim)
	}
	copy(m.Data, vals)
	return m
}
func ZeroMatrix(dim int) *Matrix {
	m := new(Matrix)
	m.Dim = dim
	m.Data = make([][]float64, dim)
	for i := range m.Data {
		m.Data[i] = make([]float64, dim)
	}
	return m
}
func EyeMatrix(dim int) *Matrix {
	m := ZeroMatrix(dim)
	for i := 0; i < dim; i++ {
		m.Data[i][i] = 1
	}
	return m
}
func (m *Matrix) Equal(m2 *Matrix) bool {
	if m.Dim != m2.Dim || m.Dim == 0 || m2.Dim == 0 {
		return false
	}
	if len(m.Data) != len(m2.Data) {
		return false
	}
	if (m.Data == nil) || (m2.Data == nil) {
		return false
	}
	for i, v := range m.Data {
		for j, w := range v {
			if !AlmostEqual(w, m2.Data[i][j], Eps) {
				return false
			}
		}
	}
	return true
}
func (m *Matrix) Multiply(m1 *Matrix) *Matrix {
	return Multiply(m, m1)
}
func Multiply(m1, m2 *Matrix) *Matrix {
	m := ZeroMatrix(4)
	dim := m1.Dim

	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			calVal := func() float64 {
				var ret float64
				for d := 0; d < dim; d++ {
					ret += m1.Data[i][d] * m2.Data[d][j]
				}
				return ret
			}
			m.Data[i][j] = calVal()
		}
	}
	return m
}

func MultiplyTuple(m *Matrix, t *Tuple) *Tuple {
	res := &Tuple{}
	dim := m.Dim
	getTupleVal := func(key int) float64 {
		var ret float64
		switch key {
		case 0:
			ret = t.X
		case 1:
			ret = t.Y
		case 2:
			ret = t.Z
		case 3:
			ret = t.W
		}
		return ret
	}
	setTupleVal := func(key int, val float64) {
		switch key {
		case 0:
			res.X = val
		case 1:
			res.Y = val
		case 2:
			res.Z = val
		case 3:
			res.W = val
		}
	}
	for i := 0; i < dim; i++ {
		var tmpSum float64
		for j := 0; j < dim; j++ {
			tj := getTupleVal(j)
			tmpSum += m.Data[i][j] * tj
		}
		setTupleVal(i, tmpSum)
	}
	return res
}
func Transpose(m *Matrix) *Matrix {
	ret := ZeroMatrix(m.Dim)
	for i := 0; i < m.Dim; i++ {
		for j := 0; j < m.Dim; j++ {
			ret.Data[j][i] = m.Data[i][j]
		}
	}
	return ret
}
func Determinant(m *Matrix) float64 {
	dim := m.Dim
	var det float64
	if dim == 2 {
		det = m.Data[0][0]*m.Data[1][1] - m.Data[0][1]*m.Data[1][0]
	} else {
		for col := 0; col < dim; col++ {
			det += m.Data[0][col] * Cofactor(m, 0, col)
		}
	}

	return det
}
func Submatrix(m *Matrix, row, col int) *Matrix {
	subDim := m.Dim - 1
	ret := ZeroMatrix(subDim)
	var k int
	for i := 0; i < m.Dim; i++ {
		if i == row {
			continue
		}
		for j := 0; j < m.Dim; j++ {
			if j == col {
				continue
			}
			ret.Data[k/subDim][k%subDim] = m.Data[i][j]
			k++
		}
	}
	return ret
}
func Minor(m *Matrix, row, col int) float64 {
	b := Submatrix(m, row, col)
	return Determinant(b)
}
func Cofactor(m *Matrix, row, col int) float64 {
	val := Minor(m, row, col)
	if (row+col)%2 == 1 {
		val = -val
	}
	return val
}
func Inverse(m *Matrix) (ret *Matrix, err error) {
	if Determinant(m) == 0 {
		return ret, errors.New("the matrix is not invertible")
	}
	ret = ZeroMatrix(m.Dim)
	det := Determinant(m)
	for row := 0; row < m.Dim; row++ {
		for col := 0; col < m.Dim; col++ {
			c := Cofactor(m, row, col)
			ret.Data[col][row] = c / det
		}
	}
	return
}
