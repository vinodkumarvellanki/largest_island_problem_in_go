//  
// file        : largest_island_logic.go
// description : this program finds the largest island from the user given array
//               of 0s and 1s in go language. array size can be upto max 100 rows and columns.
//				 user can select required rows and columns.
// author      : Vinod
// 
// update      : initial code for fixed array
//               added dynamic array input
//               added validations for input values 
//


package main

import "fmt"

// fixed rows and columns
const MAX_ROWS int = 100
const MAX_COLS int = 100

// global max_row
var max_row int
var colPerRow [100] int

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
	
	fmt.Printf("in DFSearch - Incomig count %d \n", *count)
	
    // Recur search for all connected neighbours
    for cell_no = 0; cell_no < 4; cell_no++ {
		if isPositionOk(MainArray, row+rowNbr[cell_no], col+colNbr[cell_no]) {
            // increment region length by one
            *count++
			fmt.Printf("     Incrementing island count to %d \n", *count)
			DFSearch(MainArray, row+rowNbr[cell_no], col+colNbr[cell_no], count)
        }
    }
}

// core function that finds the largest island 
func findLargestIsland(MainArray [MAX_ROWS][MAX_COLS]int) int {
 
    var row_count, column, result int

    // Initially all cells are unverified
	for row_count = 0; row_count < max_row; row_count++ {
		for column = 0; column < colPerRow[row_count]; column++ {
            verified[max_row][column] = false
        }
    }
 
    // Initialize result as 0 and travese through
    // all cells of given matrix

	result = 0
	row_count = 0
	column = 0

	for row_count = 0; row_count < max_row; row_count++ {
		for column = 0; column < colPerRow[row_count]; column++ {

			// If a cell is with value 1 and if it is not verified
			if (MainArray[row_count][column] == 1) && (verified[row_count][column] == false) {

                // if the cell is not yet verified, then new region found
                var count int
                count = 1
                fmt.Printf("findLargest - New island count %d \n", count)
				
				// search for other islands
				DFSearch(MainArray, row_count, column, &count)
                
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

	var MainArray [100][100]int
	var inputVal int
	var max_col int = 0
	max_row =0
	// get max rows from user; upto 100
	for {
		fmt.Print("\nEnter the Max No of Rows[Range 0 to 100]: ")
		fmt.Scan(&max_row)
	    if (max_row >= 1 && max_row <= 100) { break }
	}

	// get the each column elements from users
	// again columns are up to max 100
	for row_count :=0; row_count < max_row; row_count++ {
	    fmt.Print("Row No: ", row_count)
		fmt.Print(" Enter the Max Column Elements[Range 0 to 100]: ")
	    max_col = 0
		for {
			fmt.Scan(&max_col)
		    if (max_col <=0 || max_col > 100) { 
				fmt.Print(" Error !! Enter in the Range 0 to 100: ")
				continue 
			} else { break }
        }			
		// just remember the columns per each row
		// so that it can be passed to next functions
		colPerRow[row_count] = max_col;
		fmt.Print("\nEnter the column elements:[0 or 1]:\n")
		for col:= 0; col < max_col; col++ {
            for {
				fmt.Scan(&inputVal)
				if inputVal != 0 && inputVal!=1 {
					fmt.Print("\nError !! Enter [0 or 1]:\n")
					continue
				} else {
					MainArray[row_count][col] = inputVal
					break
				}  // if 0 and
			}
		}  // for columns
	}  // for rows

	fmt.Print("\nGiven Array Values:")
	for row_count:=0; row_count < max_row; row_count++ {
		fmt.Print("\n")
		for col:=0; col < colPerRow[row_count]; col++ {
			fmt.Print(MainArray[row_count][col]," ")
		}
	}
	fmt.Print("\n\n")

/*
    // fixed array size
	var MainArray = [MAX_ROWS][MAX_COLS]int{
                     {1, 0, 0, 1, 0, 1},
                     {1, 0, 0, 1, 0, 1},
                     {1, 0, 1, 1, 0, 1},
                     {1, 1, 1, 1, 0, 0}}
*/
 
    fmt.Printf("Largest island %d\n", findLargestIsland(MainArray))
     
	return
}