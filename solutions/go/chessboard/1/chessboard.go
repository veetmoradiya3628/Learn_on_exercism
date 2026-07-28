package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    ans := 0
    for _, val := range cb[file] {
        if val {
            ans++
        }
    }
    return ans
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	ans := 0
    if rank < 1 || rank > 8 {
        return ans
    }

    for _, r := range cb {
        for idx, val := range r {
            if idx + 1 == rank && val {
                ans++
            }
        }
    }
    return ans
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	ans := 0
    for _, r := range cb {
        ans += len(r)
    }
    return ans
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    ans := 0
    for _, r := range cb {
        for _, val := range r {
            if val {
                ans++
            }
        }
    }
    return ans
}
