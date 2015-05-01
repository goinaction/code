// Package listing05 implements two functions that traverse a
// matrix in different ways.
package listing05

import "fmt"

const rows = 64

// Set the size of each row to 64k
const cols = 64 * 1024

// matrix represents a set of columns that each exist on
// their own cache line.
var matrix [rows][cols]byte

// init sets ~13% of the matrix to 0XFF.
func init() {
	var ctr int
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if col%8 == 0 {
				matrix[row][col] = 0xFF
				ctr++
			}
		}
	}
	fmt.Println(ctr, "Elements set out of", cols*rows)
}

// rowTraverse traverses the matrix linearly by each column for each row.
func rowTraverse() int {
	var ctr int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}

// colTraverse traverses the matrix linearly by each row for each column.
func colTraverse() int {
	var ctr int

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}
