//  
// file        : largest_island_logic.go
// description : this program finds the largest island from the given array
//               in go language
// author      : vinod
// 
// update      : initial code for fixed array
//


package main

import "fmt"

// fixed rows and columns
const MAX_ROWS int = 4
const MAX_COLS int = 6

// 4 neighbours of a given cell
// left, right, up, down
var rowNbr = []int{-1, 1, 0, 0}
var colNbr = []int{0, 0, -1, 1}

// to remember if the column is verified or not
var verified = [MAX_ROWS][MAX_COLS]bool{{false}}

// check if the selected position is ok to include in search logic
func isPositionOk(MainArray [MAX_ROWS][MAX_COLS]int, row int, col int) bool {

    // row and column number are with in range 
	// and value is 1 and not yet verified
	
    //fmt.Printf("in isPositionOk row %d col %d \n", row, col)

	return (row >= 0) && (row < MAX_ROWS) &&
		   (col >= 0) && (col < MAX_COLS) &&
           ((MainArray[row][col]==1) && (verified[row][col] == false))
}
 
 // this function searches the islands with the given row and column
func DFSearch(MainArray [MAX_ROWS][MAX_COLS]int, row int, col int,
	count *int) {

    // Mark this cell as verified
    verified[row][col] = true

    var cell_no int
	
    fmt.Printf("in DFSearch - In count %d \n", *count)
	
    // Recur search for all connected neighbours
    for cell_no = 0; cell_no < 4; cell_no++ {
		if isPositionOk(MainArray, row+rowNbr[cell_no], col+colNbr[cell_no]) {
            // increment region length by one
            *count++
			fmt.Printf("Incrementing count to %d \n", *count)
			DFSearch(MainArray, row+rowNbr[cell_no], col+colNbr[cell_no], count)
        }
    }
}

// core function that finds the largest island 
func findLargestIsland(MainArray [MAX_ROWS][MAX_COLS]int) int {
 
    var i, j, result int

    // Initially all cells are unverified
	for i = 0; i < MAX_ROWS; i++ {
		for j = 0; j < MAX_COLS; j++ {
            verified[i][j] = false
        }
    }
 
    // Initialize result as 0 and travese through
    // all cells of given matrix

    result = 0
	i = 0
	j = 0

	for i = 0; i < MAX_ROWS; i++ {
		for j = 0; j < MAX_COLS; j++ {

            // If a cell is with value 1 and if it is not verified
			if (MainArray[i][j] == 1) && (verified[i][j] == false) {

                // if the cell is not yet verified, then new region found
                var count int
                count = 1
                fmt.Printf("findLargest - Reset count %d \n", count)
				
				// search for other islands
				DFSearch(MainArray, i, j, &count)
                
				fmt.Printf("findLargest - Search Logic count %d \n\n", count)

                // maximum region
				
				// if the new island is bigger than the old one
                if count > result {
                    result = count
				}
            }
        }
    }
    return result
}
 
// Main Function
func main() {
    // fixed array size
	var MainArray = [MAX_ROWS][MAX_COLS]int{
                     {1, 0, 0, 1, 0, 1},
                     {1, 0, 0, 1, 0, 1},
                     {1, 0, 1, 1, 0, 1},
                     {1, 1, 1, 1, 0, 0}}
 
    fmt.Printf("Largest island %d\n", findLargestIsland(MainArray))
     
	return
}