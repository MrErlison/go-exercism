package chessboard

// Declare a type named File which stores if a square is occupied by a piece -
// this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files,
// accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	_, ok := cb[file]
	if !ok {
		return 0
	}

	var squares int
	for _, square := range cb[file] {
		if square {
			squares++
		}
	}
	return squares
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank <= 0 || rank > 8 {
		return 0
	}

	var squares int
	for file := range cb {
		if cb[file][rank-1] {
			squares++
		}
	}
	return squares
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var squares int
	for _, file := range cb {
		for range file {
			squares++
		}
	}
	return squares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var squares int
	for file := range cb {
		squares += CountInFile(cb, file)
	}
	return squares
}
