package main
import "fmt"

func main(){
	x := 7
	table := [][]int{{1,5,9},{2,6,5,13},{5,3,7,4}}
	run1(x, table)
	run2(x, table)
}

func run1(x int, table [][]int){
	found := false
	for row := 0; row < len(table);row++{
		for col:= 0; col < len(table[row]); col++ {
			if table[row][col] == x {
				found = true
				fmt.Printf("found %d(row:%d, col:%d)\n", x, row, col)
				break
			}
		}
		if found{
			break
		}
	}
}

func run2(x int, table[][]int){
	LOOP:
	for row:= 0; row < len(table); row++ {
		for col := 0; col < len(table[row]);col++ {
			if table[row][col] == x {
				fmt.Printf("found %d(row: %d, col: %d)\n", x, row, col)
				break LOOP
			}
		}
	}
}
