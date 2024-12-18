// Code generated by "stringer -type=GameStatus"; DO NOT EDIT.

package tictactoe

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InProgress-0]
	_ = x[Draw-1]
	_ = x[WinnerA-2]
	_ = x[WinnerB-3]
}

const _GameStatus_name = "InProgressDrawWinnerAWinnerB"

var _GameStatus_index = [...]uint8{0, 10, 14, 21, 28}

func (i GameStatus) String() string {
	if i < 0 || i >= GameStatus(len(_GameStatus_index)-1) {
		return "GameStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _GameStatus_name[_GameStatus_index[i]:_GameStatus_index[i+1]]
}