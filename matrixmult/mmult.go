/*
Matrix multiplication if i remember how to fucking write the fucking thing
*/

package main

import(
	"fmt"
	//"math"
)

func loadMtx(mtx[10][10]int, n int, m int){
	fmt.Print("Enter Matrix Items to Multiplication = ")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
        fmt.Scan(&mtx[i][j])
		}
	}
}

func printMtx(mtx[10][10]int, n int, m int){
	fmt.Print("Printing out the matrix\n")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
        fmt.Println(mtx[i][j],"\t")
		}
	}
}
func main(){
	var n int;
	var m int;

	fmt.Print("Enter the num of rows and cols: ");
	fmt.Scan(&n,&m);

	var mtx1[10][10]int;
	var mtx2[10][10]int;
	var mtx3[10][10]int;

	loadMtx(mtx1, n,m);
	loadMtx(mtx2, n,m);

	// multiplication YIPPEE
	for row := 0; row < n; row++{
		for col := 0;col <n;col++{
			mtx3[row][col] += mtx1[row][col] * mtx2[col][row]
		}
	}
	printMtx(mtx3, n, m);
}
