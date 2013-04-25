package engine

// Utilities for setting magnetic configurations.
// TODO: use [3][][][]float32, hide data.Slice API.
// Requires careful ZYX translation.

import (
	"code.google.com/p/mx3/data"
	"code.google.com/p/mx3/util"
)

// Make a vortex magnetization with given circulation and core polarization (+1 or -1)
// Example:
// 	M.Set(Vortex(1, 1)) // counterclockwise, core up
func Vortex(circ, pol int) *data.Slice {
	util.Argument(circ == 1 || circ == -1)
	util.Argument(pol == 1 || pol == -1)

	m := data.NewSlice(3, &mesh)
	v := m.Vectors()
	cy, cz := len(v[0][0])/2, len(v[0][0][0])/2
	for i := range v[0] {
		for j := range v[0][i] {
			for k := range v[0][0][j] {
				y := j - cy
				x := k - cz
				v[X][i][j][k] = 0
				v[Y][i][j][k] = float32(x * circ)
				v[Z][i][j][k] = float32(-y * circ)
			}
		}
		v[Z][i][cy][cz] = 0.
		v[Y][i][cy][cz] = 0.
		v[X][i][cy][cz] = float32(pol)
	}
	return m
}

// Make a 2-domain configuration with domain wall.
// (mx1, my1, mz1) and (mx2, my2, mz2) are the magnetizations in the left and right domain, respectively.
// (mxwall, mywall, mzwall) is the magnetization in the wall.
// E.g.:
// 	M.Set(TwoDomain(1,0,0,  0,1,0,  -1,0,0)) // head-to-head domains with transverse (Néel) wall
// 	M.Set(TwoDomain(1,0,0,  0,0,1,  -1,0,0)) // head-to-head domains with perpendicular (Bloch) wall
// 	M.Set(TwoDomain(0,0,1,  1,0,0,   0,0,-1))// up-down domains with Bloch wall
func TwoDomain(mx1, my1, mz1, mxwall, mywall, mzwall, mx2, my2, mz2 float64) *data.Slice {
	m1 := [3]float32{float32(mz1), float32(my1), float32(mx1)}
	m2 := [3]float32{float32(mz2), float32(my2), float32(mx2)}
	mw := [3]float32{float32(mzwall), float32(mywall), float32(mxwall)}
	m := data.NewSlice(3, &mesh)
	Nz := mesh.Size()[0]
	Ny := mesh.Size()[1]
	Nx := mesh.Size()[2]
	util.Argument(Nx >= 4)
	v := m.Vectors()
	for c := range mw {
		for i := 0; i < Nz; i++ {
			for j := 0; j < Ny; j++ {
				for k := 0; k < Nx/2; k++ {
					v[c][i][j][k] = m1[c]
				}
				for k := Nx / 2; k < Nx; k++ {
					v[c][i][j][k] = m2[c]
				}
				v[c][i][j][Nx/2-2] += mw[c]
				v[c][i][j][Nx/2-1] = mw[c]
				v[c][i][j][Nx/2] = mw[c]
				v[c][i][j][Nx/2+1] += mw[c]
			}
		}
	}
	return m
}

// convert mumax's internal ZYX convention to userspace XYZ.
//func convertXYZ(arr [][][][]float32) *host.Array {
//	s := arr.Size3D
//	n := arr.NComp()
//	a := arr.Array
//	transp := host.NewArray(n, []int{s[Z], s[Y], s[X]})
//	t := transp.Array
//	for c := 0; c < n; c++ {
//		c2 := swapIndex(c, n)
//		for i := 0; i < s[X]; i++ {
//			for j := 0; j < s[Y]; j++ {
//				for k := 0; k < s[Z]; k++ {
//					t[c2][k][j][i] = a[c][i][j][k]
//				}
//			}
//		}
//	}
//	return transp
//}

//// Returns a function that returns the vector value for all i,j,k.
//func Uniform(x, y, z float32) func(i, j, k int) [3]float32 {
//	v := [3]float32{x, y, z}
//	return func(i, j, k int) [3]float32 {
//		return v
//	}
//}
//
//// Sets value at index i,j,k to f(i,j,k).
//func SetAll(array [3][][][]float32, f func(i, j, k int) [3]float32) {
//	n := core.SizeOf(array[0])
//	i2, j2, k2 := n[0], n[1], n[2]
//	SetRegion(array, 0, 0, 0, i2, j2, k2, f)
//}
//
//// Sets the region between (i1, j1, k1), (i2, j2, k2) to f(i,j,k).
//func SetRegion(array [3][][][]float32, i1, j1, k1, i2, j2, k2 int, f func(i, j, k int) [3]float32) {
//	for i := i1; i < i2; i++ {
//		for j := j1; j < j2; j++ {
//			for k := k1; k < k2; k++ {
//				v := f(i, j, k)
//				for c := range array {
//					array[c][i][j][k] = v[c]
//				}
//			}
//		}
//	}
//}
