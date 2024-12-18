package tictactoe

import (
	"fmt"
)

type Player int8
type GameStatus int8

type TicTacToe struct {
	grid          [][]rune
	status        int8
	currentPlayer int8
	availableCell int16
}

const (
	PlayerA Player = iota
	PlayerB
)

const (
	InProgress GameStatus = iota
	Draw
	WinnerA
	WinnerB
)

const (
	EmptySymbol   = ' '
	PlayerASymbol = 'O'
	PlayerBSymbol = 'X'
)

const GridSize = 3

func New() TicTacToe {
	g := make([][]rune, GridSize)
	for i := 0; i < GridSize; i++ {
		g[0] = make([]rune, GridSize)
	}
	return TicTacToe{
		grid:          g,
		status:        int8(InProgress),
		availableCell: 9,
		currentPlayer: int8(PlayerA),
	}
}

func (t *TicTacToe) Play() error {
	for t.status == int8(InProgress) {
		x, y := GetPlayerSymbol()
		err := t.PutSymbol(x, y, t.currentPlayer)
		if err != nil {
			return err
		}

		if t.IsOver(x, y) {
			if t.currentPlayer == int8(PlayerA) {
				t.status = int8(WinnerA)
			} else {
				t.status = int8(WinnerB)
			}
		}

		t.availableCell--
		if t.availableCell == 0 {
			t.status = int8(Draw)
		}

		t.currentPlayer = t.currentPlayer ^ 1
	}
	return nil
}

func GetPlayerSymbol() (int, int) {
	var x, y int
	fmt.Scanf("%d %d", x, y)
	return x, y
}

func (t *TicTacToe) PutSymbol(x int, y int, player int8) error {
	if x < 0 || x >= GridSize {
		return fmt.Errorf("Invalid Position")
	}

	if y < 0 || y >= GridSize {
		return fmt.Errorf("Invalid Position")
	}

	if t.grid[x][y] != EmptySymbol {
		return fmt.Errorf("Invalid Position")
	}

	if player == int8(PlayerA) {
		t.grid[x][y] = PlayerASymbol
	} else {
		t.grid[x][y] = PlayerBSymbol
	}

	return nil
}

func (t *TicTacToe) IsOver(x int, y int) bool {
	// validate rows
	count := 1
	for i := 1; i < GridSize; i++ {
		if t.grid[x][i] == t.grid[x][i-1] {
			count++
		}
	}

	if count == GridSize {
		return true
	}

	count = 1
	// validate columns
	for i := 1; i < GridSize; i++ {
		if t.grid[i][y] == t.grid[i-1][y] {
			count++
		}
	}

	if count == GridSize {
		return true
	}

	// validate diag 1
	count = 1
	for i := 1; i < GridSize; i++ {
		if t.grid[i][i] == t.grid[i-1][i-1] {
			count++
		}
	}

	if count == GridSize {
		return true
	}

	// validate diag 2
	count = 1
	for i := 1; i < GridSize; i++ {
		if t.grid[i][GridSize-i-1] == t.grid[i-1][GridSize-i-2] {
			count++
		}
	}

	if count == GridSize {
		return true
	}

	return false
}
