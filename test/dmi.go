// +build ignore

package main

// Test for the DMI interaction, inspired by Thiaville et al. EPL 100 2012.
// In this simulation, the DMI interaction converts a Bloch wall to a Néel wall.

import (
	. "code.google.com/p/mx3/engine"
	"fmt"
	"log"
	"math"
)

func main() {
	Init()
	defer Close()

	const (
		Nz, Ny, Nx = 1, 128, 128
		Sz, Sx, Sy = 0.6e-9, 250e-9, 250e-9
		cx, cy, cz = Sx / Nx, Sy / Ny, Sz / Nz
	)

	SetMesh(Nx, Ny, Nz, cx, cy, cz)

	Msat = Const(1100e3)
	Aex = Const(16e-12)
	Alpha = Const(3)
	Ku1 = ConstVector(0, 0, 1.27E6)

	M.Set(TwoDomain(0, 0, 1, 1, 1, 0, 0, 0, -1)) // up-down domains with wall between Bloch and Néél type
	M.SaveAs("m_init.dump")

	DMI = Const(0.1e-3) // J/m2.

	Solver.MaxErr = 1e-4
	Solver.Maxdt = 5e-13
	Table.Autosave(10e-12)
	M.Autosave(1000e-12)
	Run(5e-9)

	avg := M.Average()
	fmt.Println("m (D=0.1e-3 J/m2):", avg[X], avg[Y], avg[Z])
	expect(avg[X], 0.059)
	expect(avg[Y], 0.035)
	fmt.Println("OK")
}

func expect(have, want float64) {
	if math.Abs(have-want) > 1e-3 {
		log.Fatalln("have:", have, "want:", want)
	}
}
