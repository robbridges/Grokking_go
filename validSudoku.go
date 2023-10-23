package main

func IsValidSudoku(board [][]byte) bool {
	// Check rows
	for i := 0; i < 9; i++ {
		row := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if _, ok := row[board[i][j]]; ok {
					return false
				}
				row[board[i][j]] = true
			}
		}
	}
	
	// Check columns
	for j := 0; j < 9; j++ {
		col := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			if board[i][j] != '.' {
				if _, ok := col[board[i][j]]; ok {
					return false
				}
				col[board[i][j]] = true
			}
		}
	}
	
	// Check 3x3 boxes
	for k := 0; k < 9; k++ {
		box := make(map[byte]bool)
		for i := k/3*3; i < k/3*3+3; i++ {
			for j := k%3*3; j < k%3*3+3; j++ {
				if board[i][j] != '.' {
					if _, ok := box[board[i][j]]; ok {
						return false
					}
					box[board[i][j]] = true
				}
			}
		}
	}
	
	return true
}