package udemyproblems

import (
	"fmt"
)

var fourspaces = "    "

func abs(v1 int, v2 int) int {

	absolutevalue := 0

	if v1 > v2 {
		absolutevalue = v1 - v2
	} else {
		absolutevalue = v2 - v1
	}
	return absolutevalue
}

func printrow(B [][]int, rownum int) {

	for i := 0; i < len(B); i++ {
		fmt.Print(B[rownum][i], " ")
	}

	fmt.Println("")
}

func printboard(B [][]int) {

	printrow(B, 0)
	printrow(B, 1)
	printrow(B, 2)
	printrow(B, 3)
}

func allowed(board [][]int, newqueenrowpos int, newqueencolpos int) bool {

	answer := true
	answer = checkrowallowed(board, newqueenrowpos, answer)

	if answer {
		answer = checkcolumnallowed(board, newqueencolpos, answer)
	}

	if answer {
		answer = checkcdiagonalallowed(board, newqueenrowpos, newqueencolpos, answer)
	}
	return answer
}

func checkcdiagonalallowed(board [][]int, newqueenrpos int, newqueencpos int, answer bool) bool {
	qcount := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {

			if abs(newqueenrpos, i) == abs(newqueencpos, j) {
				if board[i][j] == 1 {
					qcount++
				}
			}

		}
	}

	if qcount > 1 {
		answer = false
	}

	return answer
}

func checkcolumnallowed(board [][]int, newqueenxpos int, answer bool) bool {
	qcount := 0

	for i := 0; i < len(board); i++ {
		if board[i][newqueenxpos] == 1 {
			qcount++
		}
	}
	if qcount > 1 {
		answer = false
	}
	return answer
}

func checkrowallowed(board [][]int, newqueenrowpos int, answer bool) bool {
	qcount := 0

	//check all coulmns for new queen row position
	for i := 0; i < len(board); i++ {
		if board[newqueenrowpos][i] == 1 {
			qcount++
		}
	}
	if qcount > 1 {
		answer = false
	}
	return answer
}

func placefourthqueen(b [][]int) {

	for r4 := 0; r4 < 4; r4++ {
		//place queen on column 4
		b[r4][3] = 1

		if allowed(b, r4, 3) {
			printboard(b)
			fmt.Println(" ")
		} else {
			b[r4][3] = 0
		}
	}
}

func placethirdqueen(b [][]int) {

	for r3 := 0; r3 < 4; r3++ {
		//place queen on column 3
		b[r3][2] = 1

		if allowed(b, r3, 2) {
			placefourthqueen(b)
		} else {
			b[r3][2] = 0
		}
	}
}

func placesecondqueen(board [][]int) {

	for r := 0; r < len(board); r++ {
		//place q on column 2
		board[r][1] = 1
		if allowed(board, r, 1) {

			placethirdqueen(board)

		} else {
			board[r][1] = 0
		}
	}
}

//Givefourqueensposition gives position of each of the four queens
func Givefourqueensposition() {

	N := 4

	for r := 0; r < N; r++ {
		//place q on col 1
		var board = [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
		
		board[r][0] = 1
		placesecondqueen(board)
		board[r][0] = 0
	}

}
