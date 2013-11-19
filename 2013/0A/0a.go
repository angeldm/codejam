package main

import (
	. "github.com/angeldm/codejam"
	"strings"
)

const (
	XWON         = "X won"
	OWON         = "O won"
	DRAW         = "Draw"
	NOTCOMPLETED = "Game has not completed"
	FOUR_X       = "XXXX"
	FOUR_O       = "OOOO"
	THREE_X      = "XXX"
	THRE_O       = "OOO"
	DOT          = "."
	T            = "T"
)

func main() {
	Start(solve)
}

func solve() string {
	horizontal := make([]string, 4)
	vertical := make([]string, 4)
	diagonal := make([]string, 6)
	hasdot := false

	horizontal[0] = ReadString()
	horizontal[1] = ReadString()
	horizontal[2] = ReadString()
	horizontal[3] = ReadString()
	ReadString()

	vertical[0] = string(horizontal[0][0]) + string(horizontal[1][0]) + string(horizontal[2][0]) + string(horizontal[3][0])
	vertical[1] = string(horizontal[0][1]) + string(horizontal[1][1]) + string(horizontal[2][1]) + string(horizontal[3][1])
	vertical[2] = string(horizontal[0][2]) + string(horizontal[1][2]) + string(horizontal[2][2]) + string(horizontal[3][2])
	vertical[3] = string(horizontal[0][3]) + string(horizontal[1][3]) + string(horizontal[2][3]) + string(horizontal[3][3])

	diagonal[0] = string(horizontal[2][0]) + string(horizontal[1][1]) + string(horizontal[0][2])
	diagonal[1] = string(horizontal[3][0]) + string(horizontal[2][1]) + string(horizontal[1][2]) + string(horizontal[0][3])
	diagonal[2] = string(horizontal[3][1]) + string(horizontal[2][2]) + string(horizontal[1][3])

	diagonal[3] = string(horizontal[0][1]) + string(horizontal[1][2]) + string(horizontal[2][3])
	diagonal[4] = string(horizontal[0][0]) + string(horizontal[1][1]) + string(horizontal[2][2]) + string(horizontal[3][3])
	diagonal[5] = string(horizontal[1][0]) + string(horizontal[2][1]) + string(horizontal[3][2])

	for i := 0; i < 4; i++ {
		X_PATTERN := FOUR_X
		O_PATTERN := FOUR_O
		if strings.Contains(horizontal[i], T) {
			X_PATTERN = "XXX"
			O_PATTERN = "OOO"
		}

		if strings.Contains(horizontal[i], X_PATTERN) {
			return XWON
		}
		if strings.Contains(horizontal[i], O_PATTERN) {
			return OWON
		}
		hasdot = strings.Contains(horizontal[i], DOT)
	}

	for i := 0; i < 4; i++ {
		X_PATTERN := FOUR_X
		O_PATTERN := FOUR_O
		if strings.Contains(vertical[i], T) {
			X_PATTERN = "XXX"
			O_PATTERN = "OOO"
		}

		if strings.Contains(vertical[i], X_PATTERN) {
			return XWON
		}
		if strings.Contains(vertical[i], O_PATTERN) {
			return OWON
		}
		hasdot = strings.Contains(vertical[i], DOT)
	}

	for i := 0; i < 6; i++ {
		X_PATTERN := FOUR_X
		O_PATTERN := FOUR_O
		if strings.Contains(diagonal[i], T) {
			X_PATTERN = "XXX"
			O_PATTERN = "OOO"
		}
		if strings.Contains(diagonal[i], X_PATTERN) {
			return XWON
		}
		if strings.Contains(diagonal[i], O_PATTERN) {
			return OWON
		}
		hasdot = strings.Contains(diagonal[i], DOT)
	}

	if hasdot == true {
		return NOTCOMPLETED
	} else {
		return DRAW
	}

	return ""
}
