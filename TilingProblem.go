package pftp

import "fmt"

// m is minimum 2X2
func tileremainingsquares(m [][]int, tiletype, sxpos, sypos int) {
	for i := range m {
		for j := range m[i] {
			if !(i == sxpos && j == sypos) {
				m[i][j] = tiletype
			}

		}
	}
}

func TileIt(m [][]int, statuexpos, statueypos int) [][]int {

	debglvl := 1

	m[statuexpos][statueypos] = 42

	if len(m) == 2 {
		tileremainingsquares(m, 79, statuexpos, statueypos)
		return m
	}

	halflen := len(m) / 2

	m1 := make([][]int, halflen, halflen)
	m2 := make([][]int, halflen, halflen)
	m3 := make([][]int, halflen, halflen)
	m4 := make([][]int, halflen, halflen)

	m1 = populatem1(halflen, m1, m)

	m2 = populatem2(halflen, m2, m)

	m3 = populatem3(halflen, m3, m)

	m4 = populatem4(halflen, m4, m)

	fmt.Println("statuexpos = ", statuexpos, " , statueypos = ", statueypos)

	bottomhalf, righthalf := findstatuequadrant(statuexpos, halflen, statueypos, debglvl)

	var newm1x, newm1y int
	var newm2x, newm2y int
	var newm3x, newm3y int
	var newm4x, newm4y int

	newm1x, newm1y, newm2x, newm2y, newm3x, newm3y, newm4x, newm4y = getnewstatueposition(bottomhalf, righthalf, debglvl, newm1x, newm1y, newm2x, newm2y, newm3x, newm3y, newm4x, newm4y, statuexpos, statueypos, m1)

	m1 = TileIt(m1, newm1x, newm1y)
	writedebugmessage(debglvl, "after tiling m1")

	m2 = TileIt(m2, newm2x, newm2y)
	writedebugmessage(debglvl, "after tiling m2")

	m3 = TileIt(m3, newm3x, newm3y)
	writedebugmessage(debglvl, "after tiling m3")

	m4 = TileIt(m4, newm4x, newm4y)
	writedebugmessage(debglvl, "after tiling m4")

	return m
}

func populatem4(halflen int, m4 [][]int, m [][]int) [][]int {
	for i := 0; i < halflen; i++ {
		m4[i] = m[i+halflen][halflen:]
	}

	return m4
}

func populatem3(halflen int, m3 [][]int, m [][]int) [][]int {
	for i := 0; i < halflen; i++ {
		m3[i] = m[i+halflen][:halflen]
	}

	return m3
}

func populatem2(halflen int, m2 [][]int, m [][]int) [][]int {
	for i := 0; i < halflen; i++ {
		m2[i] = m[i][halflen:]
	}

	return m2
}

func populatem1(halflen int, m1 [][]int, m [][]int) [][]int {

	for i := 0; i < halflen; i++ {

		m1[i] = m[i][:halflen]
	}

	return m1
}

func getnewstatueposition(bottomhalf bool, righthalf bool, debglvl int, newm1x int, newm1y int, newm2x int, newm2y int, newm3x int, newm3y int, newm4x int, newm4y int, statuexpos int, statueypos int, m1 [][]int) (int, int, int, int, int, int, int, int) {
	if bottomhalf && righthalf {
		writedebugmessage(debglvl, "statue in m4")

		newm1x, newm1y, newm2x, newm2y, newm3x, newm3y, newm4x, newm4y = getnewstatuepositionsm4(statuexpos, statueypos, len(m1), debglvl)

	} else if bottomhalf && !righthalf {

		writedebugmessage(debglvl, "statue in m3")

		newm1x, newm1y, newm2x, newm2y, newm3x, newm3y, newm4x, newm4y = getnewstatuepositionsm3(statuexpos, statueypos, len(m1), debglvl)

	} else if !bottomhalf && !righthalf {

		writedebugmessage(debglvl, "statue in m1")

		newm1x, newm1y, newm2x, newm2y, newm3x, newm3y, newm4x, newm4y = getnewstatuepositionsm1(statuexpos, statueypos, len(m1))

	} else if !bottomhalf && righthalf {

		writedebugmessage(debglvl, "statue in m2")

		newm1x, newm1y, newm2x, newm2y, newm3x, newm3y, newm4x, newm4y = getnewstatuepositionsm2(statuexpos, statueypos, len(m1))

	}
	return newm1x, newm1y, newm2x, newm2y, newm3x, newm3y, newm4x, newm4y
}

func getnewstatuepositionsm4(spx, spy, lenofm, dbglvl int) (int, int, int, int, int, int, int, int) {

	var m1x, m1y, m2x, m2y, m3x, m3y, m4x, m4y int

	m1x = lenofm - 1
	m1y = lenofm - 1

	if dbglvl > 0 {
		fmt.Println("m1x = ", m1x, " , m1y = ", m1y)
	}

	m2x = lenofm - 1
	m2y = 0

	if dbglvl > 0 {
		fmt.Println("m2x = ", m2x, " , m2y = ", m2y)
	}

	m3x = 0
	m3y = lenofm - 1

	if dbglvl > 0 {
		fmt.Println("m3x = ", m3x, " , m3y = ", m3y)
	}

	m4x = spx - lenofm
	m4y = spy - lenofm

	if dbglvl > 0 {
		fmt.Println("m4x = ", m4x, " , m4y = ", m4y)
	}

	return m1x, m1y, m2x, m2y, m3x, m3y, m4x, m4y
}

func getnewstatuepositionsm3(spx, spy, lenofm, dbglvl int) (int, int, int, int, int, int, int, int) {

	var m1x, m1y, m2x, m2y, m3x, m3y, m4x, m4y int

	m1x = lenofm - 1
	m1y = lenofm - 1

	if dbglvl > 0 {
		fmt.Println("m1x = ", m1x, " , m1y = ", m1y)
	}

	m2x = lenofm - 1
	m2y = 0

	if dbglvl > 0 {
		fmt.Println("m2x = ", m2x, " , m2y = ", m2y)
	}

	m3x = spx - lenofm
	m3y = spy

	if dbglvl > 0 {
		fmt.Println("m3x = ", m3x, " , m3y = ", m3y)
	}

	m4x = 0
	m4y = 0

	if dbglvl > 0 {
		fmt.Println("m4x = ", m4x, " , m4y = ", m4y)
	}

	return m1x, m1y, m2x, m2y, m3x, m3y, m4x, m4y
}

func getnewstatuepositionsm2(spx, spy, lenofm int) (int, int, int, int, int, int, int, int) {

	var m1x, m1y, m2x, m2y, m3x, m3y, m4x, m4y int

	m1x = lenofm - 1
	m1y = lenofm - 1

	m2x = spx
	m2y = spy - lenofm

	m3x = 0
	m3y = lenofm - 1

	m4x = 0
	m4y = 0

	return m1x, m1y, m2x, m2y, m3x, m3y, m4x, m4y
}

func getnewstatuepositionsm1(spx, spy, lenofm int) (int, int, int, int, int, int, int, int) {

	var m1x, m1y, m2x, m2y, m3x, m3y, m4x, m4y int

	m1x = spx
	m1y = spy

	m2x = lenofm - 1
	m2y = 0

	m3x = 0
	m3y = lenofm - 1

	m4x = 0
	m4y = 0

	return m1x, m1y, m2x, m2y, m3x, m3y, m4x, m4y
}

func findstatuequadrant(statuexpos int, halflen int, statueypos int, debglvl int) (bool, bool) {
	bottomhalf := false
	if statuexpos >= halflen {
		bottomhalf = true
	}

	righthalf := false
	if statueypos >= halflen {
		righthalf = true
	}

	if righthalf {
		writedebugmessage(debglvl, "statue in righthalf")
	} else {
		writedebugmessage(debglvl, "statue in lefthalf")
	}

	if bottomhalf {
		writedebugmessage(debglvl, "statue in bottomhalf")
	} else {
		writedebugmessage(debglvl, "statue in upperhalf")
	}
	return bottomhalf, righthalf
}
