package matrices

import (
	"testing"
	"tuples"
)

func TestNewMatrix4(t *testing.T) {
	m := NewMatrix(4, [][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	})
	if m.Data[0][0] != 1 {
		t.Fatal("m.Data[0][0]!=1")
	}
	if m.Data[0][3] != 4 {
		t.Fatal("m.Data[0][3]!=4")
	}
	if m.Data[1][0] != 5.5 {
		t.Fatal("m.Data[1][0]!=5.5")
	}
	if m.Data[1][2] != 7.5 {
		t.Fatal("m.Data[1][2]!=7.5")
	}
	if m.Data[2][2] != 11 {
		t.Fatal("m.Data[2][2]!=11")
	}
	if m.Data[3][0] != 13.5 {
		t.Fatal("m.Data[3][0]!=13.5")
	}
	if m.Data[3][2] != 15.5 {
		t.Fatal("m.Data[3][2]!=15.5")
	}
}

func TestNewMatrix2(t *testing.T) {
	m := NewMatrix(2, [][]float64{
		{-3, 5},
		{1, -2},
	})
	if m.Data[0][0] != -3 {
		t.Fatal("m.Data[0][0]!=-3")
	}
	if m.Data[0][1] != 5 {
		t.Fatal("m.Data[0][1]!=5")
	}
	if m.Data[1][0] != 1 {
		t.Fatal("m.Data[1][0]!=1")
	}
	if m.Data[1][1] != -2 {
		t.Fatal("m.Data[1][1]!=-2")
	}
}
func TestNewMatrix3(t *testing.T) {
	m := NewMatrix(3, [][]float64{
		{-3, 5, 0},
		{1, -2, -7},
		{0, 1, 1},
	})
	if m.Data[0][0] != -3 {
		t.Fatal("m.Data[0][0]!=-3")
	}
	if m.Data[1][1] != -2 {
		t.Fatal("m.Data[1][1]!=-2")
	}
	if m.Data[2][2] != 1 {
		t.Fatal("m.Data[2][2]!=1")
	}
}
func TestMatrixCompareSame(t *testing.T) {
	m1 := NewMatrix(4, [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	m2 := NewMatrix(4, [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	if !IsSame(m1, m2) {
		t.Fatal("m1 m2 should be same")
	}
}

func TestMatrixCompareDifferent(t *testing.T) {
	m1 := NewMatrix(4, [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	m2 := NewMatrix(4, [][]float64{
		{2, 3, 4, 5},
		{6, 7, 8, 9},
		{8, 7, 6, 5},
		{4, 3, 2, 1},
	})
	if IsSame(m1, m2) {
		t.Fatal("m1 m2 should be different")
	}
}
func TestMultiply(t *testing.T) {
	m1 := NewMatrix(4, [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	m2 := NewMatrix(4, [][]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	})
	n := Multiply(m1, m2)
	if !IsSame(n, NewMatrix(4, [][]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	})) {
		t.Fatal("matrix multiply failed")
	}
}

func TestMultiplyTuple(t *testing.T) {
	m := NewMatrix(4, [][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	})
	n := &tuples.Tuple{X: 1, Y: 2, Z: 3, W: 1}
	res := MultiplyTuple(m, n)
	if !res.Equal(&tuples.Tuple{X: 18, Y: 24, Z: 33, W: 1}) {
		t.Fatal("matrix multiply with tuple failed")
	}
}

func TestMultiplyIdentityMatrix(t *testing.T) {
	m := NewMatrix(4, [][]float64{
		{0, 1, 2, 4},
		{1, 2, 4, 8},
		{2, 4, 8, 16},
		{4, 8, 16, 32},
	})
	n := EyeMatrix(4)
	res := Multiply(m, n)
	if !IsSame(res, m) {
		t.Fatal("matrix multiply with identify matrix is not equal to itself")
	}
}

func TestMultiplyTupleWithIdentityMatrix(t *testing.T) {
	m := EyeMatrix(4)
	n := &tuples.Tuple{X: 1, Y: 2, Z: 3, W: 4}
	res := MultiplyTuple(m, n)
	if !res.Equal(n) {
		t.Fatal("identity matrix multiply with tuple failed")
	}
}

func TestTranspose(t *testing.T) {
	m := NewMatrix(4, [][]float64{
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8},
	})
	n := Transpose(m)
	if !IsSame(n, NewMatrix(4, [][]float64{
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8},
	})) {
		t.Fatal("transpose failed")
	}
}
func TestTransposeIdentityMatrix(t *testing.T) {
	m := EyeMatrix(4)
	n := Transpose(m)
	if !IsSame(n, m) {
		t.Fatal("transpose identity matrix is not equal to itself")
	}
}
func TestDeterminant2(t *testing.T) {
	m := NewMatrix(2, [][]float64{
		{1, 5},
		{-3, 2},
	})
	if Determinant(m) != 17 {
		t.Fatal("determinant 2x2 matrix failed")
	}
}
func TestSubmatrix3(t *testing.T) {
	m := NewMatrix(3, [][]float64{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	})
	n := Submatrix(m, 0, 2)
	if !IsSame(n, NewMatrix(2, [][]float64{
		{-3, 2},
		{0, 6},
	})) {
		t.Fatal("submatrix failed")
	}
}
func TestSubmatrix4(t *testing.T) {
	m := NewMatrix(4, [][]float64{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	})
	n := Submatrix(m, 2, 1)
	if !IsSame(n, NewMatrix(3, [][]float64{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	})) {
		t.Fatal("submatrix failed")
	}
}
func TestMinor(t *testing.T) {
	m := NewMatrix(3, [][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	b := Submatrix(m, 1, 0)
	val := Determinant(b)
	if val != Minor(m, 1, 0) {
		t.Fatal("minor failed")
	}
}
func TestCofactor(t *testing.T) {
	m := NewMatrix(3, [][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	if Minor(m, 0, 0) != Cofactor(m, 0, 0) {
		t.Fatal("cofactor failed")
	}
	if Minor(m, 1, 0) == Cofactor(m, 1, 0) {
		t.Fatal("cofactor failed")
	}
}

func TestDeterminant3(t *testing.T) {
	m := NewMatrix(3, [][]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	})
	if Cofactor(m, 0, 0) != 56 {
		t.Fatal("determinant failed")
	}
	if Cofactor(m, 0, 1) != 12 {
		t.Fatal("determinant failed")
	}
	if Cofactor(m, 0, 2) != -46 {
		t.Fatal("determinant failed")
	}
	if Determinant(m) != -196 {
		t.Fatal("determinant failed")
	}
}
func TestDeterminant4(t *testing.T) {
	m := NewMatrix(4, [][]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	})
	if Cofactor(m, 0, 0) != 690 {
		t.Fatal("determinant failed")
	}
	if Cofactor(m, 0, 1) != 447 {
		t.Fatal("determinant failed")
	}
	if Cofactor(m, 0, 2) != 210 {
		t.Fatal("determinant failed")
	}
	if Cofactor(m, 0, 3) != 51 {
		t.Fatal("determinant failed")
	}
	if Determinant(m) != -4071 {
		t.Fatal("determinant failed")
	}
}
func TestInvertible(t *testing.T) {
	m := NewMatrix(4, [][]float64{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6},
	})
	det := Determinant(m)
	if det == 0 {
		t.Fatal("the matrix is invertible but the answer is not")
	}
	if det != -2120 {
		t.Fatal("determinant failed")
	}

}

func TestInvertible2(t *testing.T) {
	m := NewMatrix(4, [][]float64{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0},
	})
	det := Determinant(m)
	if det != 0 {
		t.Fatal("determinant failed and A is not invertible but the answer is invertible")
	}
}

func TestInverse(t *testing.T) {
	m := NewMatrix(4, [][]float64{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4},
	})
	b, err := Inverse(m)
	if err != nil {
		t.Fatal("inverse failed")
	}
	det := Determinant(m)
	if det != 532 {
		t.Fatal("determinant failed")
	}
	if Cofactor(m, 2, 3) != -160 {
		t.Fatal("cofactor failed")
	}
	if !tuples.AlmostEqual(b.Data[3][2], -160.0/532.0, tuples.Eps) {
		t.Fatal("inverted matrix wrong,b:", b)
	}
	if Cofactor(m, 3, 2) != 105 {
		t.Fatal("cofactor failed")
	}
	if !tuples.AlmostEqual(b.Data[2][3], 105.0/532.0, tuples.Eps) {
		t.Fatal("inverted matrix wrong,b:", b)
	}
	if !IsSame(b, NewMatrix(4, [][]float64{
		{0.21805, 0.45113, 0.24060, -0.04511},
		{-0.80827, -1.45677, -0.44361, 0.52068},
		{-0.07895, -0.22368, -0.05263, 0.19737},
		{-0.52256, -0.81391, -0.30075, 0.30639},
	})) {
		t.Fatal("inverted matrix wrong,b:", b)
	}
}

func TestInverse2(t *testing.T) {
	m := NewMatrix(4, [][]float64{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	})
	b, err := Inverse(m)
	if !IsSame(NewMatrix(4, [][]float64{
		{-0.15385, -0.15385, -0.28205, -0.53846},
		{-0.07692, 0.12308, 0.02564, 0.03077},
		{0.35897, 0.35897, 0.43590, 0.92308},
		{-0.69231, -0.69231, -0.76923, -1.92308},
	}), b) || err != nil {
		t.Fatal("inverse failed")
	}
}
func TestInverse3(t *testing.T) {
	m := NewMatrix(4, [][]float64{
		{9, 3, 0, 9},
		{-5, -2, -6, -3},
		{-4, 9, 6, 4},
		{-7, 6, 6, 2},
	})
	b, err := Inverse(m)
	if !IsSame(NewMatrix(4, [][]float64{
		{-0.04074, -0.07778, 0.14444, -0.22222},
		{-0.07778, 0.03333, 0.36667, -0.33333},
		{-0.02901, -0.14630, -0.10926, 0.12963},
		{0.17778, 0.06667, -0.26667, 0.33333},
	}), b) || err != nil {
		t.Fatal("inverse failed")
	}
}
func TestMutiplyAProductByItsInverse(t *testing.T) {
	a := NewMatrix(4, [][]float64{
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1},
	})
	b := NewMatrix(4, [][]float64{
		{8, 2, 2, 2},
		{3, -1, 7, 0},
		{7, 0, 5, 4},
		{6, -2, 0, 5},
	})
	c := Multiply(a, b)
	inv, err := Inverse(b)
	if err != nil {
		t.Fatal("calculate inverse failed")
	}
	if !IsSame(Multiply(c, inv), a) {
		t.Fatal("c*b^-1!=a")
	}
}
