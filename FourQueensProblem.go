package udemyproblems

import (
	"fmt"
)

func abs(v1 int, v2 int) int {

	absolutevalue := 0

	if v1 > v2 {
		absolutevalue = v1 - v2
	} else {
		absolutevalue = v2 - v1
	}
	return absolutevalue
}

func printrow(B [4][4]int, rownum int) {

	for i := 0; i < len(B); i++ {
		fmt.Print(B[rownum][i], " ")
	}

	fmt.Println("")
}

func printboard(B [4][4]int) {

	printrow(B, 0)
	printrow(B, 1)
	printrow(B, 2)
	printrow(B, 3)
}

func allowed(x1 int, y1 int, x2 int, y2 int) bool {

	answer := true

	if x1 == x2 {
		answer = false
	}

	if y1 == y2 {
		answer = false
	}

	if abs(x2, x1) == 1 && abs(y2, y1) == 1 {
		answer = false
	}

	return answer
}

func fn0(board [4][4]int, r2 int, r int) {
	board[r2][1] = 1
	if allowed(r, 0, r2, 1) {
		printboard(board)
		fmt.Println(" ")
	} else {
		board[r2][1] = 0
	}
}

//Givefourqueensposition gives position of each of the four queens
func Givefourqueensposition() {

	N := 4

	var board = [4][4]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}

	for r := 0; r < N; r++ {
		fmt.Println("----")
		board[r][0] = 1
		for r2 := 0; r2 < N; r2++ {
			fn0(board, r2, r)
		}
		board[r][0] = 0
	}

}
