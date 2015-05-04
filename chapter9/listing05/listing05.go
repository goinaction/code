// Package listing05 implements two functions that traverse a
// matrix in different ways.
package listing05

const cols = 64

// Set the size of each row to be 64k.
const rows = 64 * 1024

// matrix represents a set of columns that each exist on
// their own cache line.
var matrix [cols][rows]byte

// init sets ~13% of the matrix to 0XFF.
func init() {
	var ctr int
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if row%8 == 0 {
				matrix[col][row] = 0xFF
				ctr++
			}
		}
	}
	println(ctr, "Elements set out of", cols*rows)
}

// rowTraverse traverses the matrix linearly by each column for each row.
func rowTraverse() int {
	var ctr int

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if matrix[col][row] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}

// colTraverse traverses the matrix linearly by each row for each column.
func colTraverse() int {
	var ctr int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[col][row] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}
