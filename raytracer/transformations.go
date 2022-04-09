package raytracer

import (
	"math"
)

func Translation(x, y, z float64) *Matrix {
	ret := EyeMatrix(4)
	ret.Data[0][3] = x
	ret.Data[1][3] = y
	ret.Data[2][3] = z
	return ret
}
func Scaling(x, y, z float64) *Matrix {
	ret := EyeMatrix(4)
	ret.Data[0][0] = x
	ret.Data[1][1] = y
	ret.Data[2][2] = z
	return ret
}
func RotationX(rad float64) *Matrix {
	ret := EyeMatrix(4)
	ret.Data[1][1] = math.Cos(rad)
	ret.Data[1][2] = -math.Sin(rad)
	ret.Data[2][1] = math.Sin(rad)
	ret.Data[2][2] = math.Cos(rad)
	return ret
}
func RotationY(rad float64) *Matrix {
	ret := EyeMatrix(4)
	ret.Data[0][0] = math.Cos(rad)
	ret.Data[0][2] = math.Sin(rad)
	ret.Data[2][0] = -math.Sin(rad)
	ret.Data[2][2] = math.Cos(rad)
	return ret
}
func RotationZ(rad float64) *Matrix {
	ret := EyeMatrix(4)
	ret.Data[0][0] = math.Cos(rad)
	ret.Data[0][1] = -math.Sin(rad)
	ret.Data[1][0] = math.Sin(rad)
	ret.Data[1][1] = math.Cos(rad)
	return ret
}
func Shearing(xy, xz, yx, yz, zx, zy float64) *Matrix {
	ret := EyeMatrix(4)
	ret.Data[0][1] = xy
	ret.Data[0][2] = xz
	ret.Data[1][0] = yx
	ret.Data[1][2] = yz
	ret.Data[2][0] = zx
	ret.Data[2][1] = zy
	return ret
}
func ViewTransform(from, to, up *Tuple) *Matrix {
	forward := Normalize(Subtract(to, from))
	left := Cross(forward, Normalize(up))
	trueUp := Cross(left, forward)
	orientation := NewMatrix(4, [][]float64{
		{left.X, left.Y, left.Z, 0},
		{trueUp.X, trueUp.Y, trueUp.Z, 0},
		{-forward.X, -forward.Y, -forward.Z, 0},
		{0, 0, 0, 1},
	})
	return orientation.Multiply(Translation(-from.X, -from.Y, -from.Z))
}
