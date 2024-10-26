package main

import (
	"fmt"
	"math/rand"
)

const (
	// ?*? map
	DIM = 7
	// conceal ? mines
	MINECNT = 6
)

// two dimension array represent a mine map, include the mine and around count
// 0-8 represent the mine number around
// 9 represent this is a mine
var mineMap [DIM][DIM] int
// two dimension array represent a display map,
// 0 represent the spot still concealed
// 1 represent the spot been digged, revealed
// 2 represent the spot been marked
// 3 represent the spot is mine, will use only game over
var dispMap [DIM][DIM] int
// how many 

func main() {

	fmt.Println("Mine Starting ...")
	
	// 1 init map
	init_map ();
	
	// 2 dig and mark
	dig_mark ();
	
	//show_mineMap();
}

func dig_mark () {
	var c string;
	var x, y, res, digCnt int;
	for {
		// 1 show the current dispMap
		show_dispMap ();
		
		// 2 get user's operation
		c, x, y = get_operation ();
	
		// 3 handle this operation
		if c == "d" {
			res = do_dig (x, y);
			// 3.1 if trigger mine, show the map and break
			if res == 0 {
				fmt.Println("BoooooooooooooooooooooooooooooM!!!")
				show_dispMap()
				return
			}
		} else if c == "m" {
			do_mark (x, y)
		} else if c == "s" {
			reveal_all();
			show_dispMap();
			fmt.Println("Sooooooooooooooooooooooooooooooooooooy")
			return
		} else {
			fmt.Println("Must input 'd' to dig or 'm' to mark or 's' to surrend!")
		}
		
		// 4 verify all un-mine be revealed, success!
		digCnt = dispMap_digCnt ()
		if digCnt == DIM * DIM - MINECNT {
			fmt.Println("WooooooooooooooooooooooooooooooooooooooooooooW")
			return
		}
	} // end loop
}

func dispMap_digCnt () (int) {
	cnt := 0
	for i := 0; i < DIM; i++ {
		for j := 0; j < DIM; j++ {
			if dispMap[i][j] == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func reveal_all () {
	for i := 0; i < DIM; i++ {
		for j := 0; j < DIM; j++ {
			dispMap[i][j] = 1
		}
	}
}

func do_mark (x int, y int) {
	if dispMap[x][y] == 0 {
		dispMap[x][y] = 2;
	} else if dispMap[x][y] == 1 {
		fmt.Println("This cell already revealed, couldn't be marked any more!")
	} else if dispMap[x][y] == 2 {
		// un-mark the marked cell
		dispMap[x][y] = 0;
	} else {
		fmt.Println("!!!IMPOSSIBLE, disMap have cell not equal 0~2 !!!")
	}
}

// if game over, return 0; else, return 1
func do_dig (x int, y int) (int) {
	// 1 judge if the cell already revealed, return if already revealed
	if (dispMap[x][y] == 1) {
		fmt.Println("Already digged this cell!")
		return 1
	}
	// 2 judge if the cell is mine, reveal all mine cells and return 0
	if mineMap[x][y] == 9 {
		for i := 0; i < DIM; i++ {
			for j := 0; j < DIM; j++ {
				if mineMap[i][j] == 9 {
					dispMap[i][j] = 3;
				}
			}
		}
		return 0		
	}
	
	// 3 if the cell number not equal to 0, just reveal the cell
	if mineMap[x][y] != 0 {
		dispMap[x][y] = 1;
		return 1
	} else {
		// 4 else, reveal all conjoined cells equal to 0 and number edge
		reveal_zero_cell(x, y);
		return 1
	}
}

func reveal_zero_cell (x int, y int) {
	// 1 reveal this cell
	dispMap[x][y] = 1;
	// 2 if the 4 neighbors is 0, reveal the neighbors
	if x-1 >=0 && dispMap[x-1][y] == 0 && mineMap[x-1][y] == 0 {
		reveal_zero_cell(x-1, y)
	}
	if y+1 < DIM && dispMap[x][y+1] == 0 && mineMap[x][y+1] == 0 {
		reveal_zero_cell(x, y+1)
	}
	if x+1 < DIM && dispMap[x+1][y] == 0 && mineMap[x+1][y] == 0 {
		reveal_zero_cell(x+1, y)
	}
	if y-1 >= 0 && dispMap[x][y-1] == 0 && mineMap[x][y-1] == 0 {
		reveal_zero_cell(x, y-1)
	}
	// 3 for the zero cell 8 neighbors, should either zero cells or number cells, reveal if under conceal
	// the reason to add mineMap[][]!=0 is not break the recursion
	if x-1 >= 0 && y-1 >= 0 && dispMap[x-1][y-1] == 0 && mineMap[x-1][y-1] != 0 {
		dispMap[x-1][y-1] = 1
	} 
	if x-1 >= 0 && dispMap[x-1][y] == 0 && mineMap[x-1][y] != 0 {
		dispMap[x-1][y] = 1
	} 
	if x-1 >= 0 && y+1 < DIM && dispMap[x-1][y+1] == 0 && mineMap[x-1][y+1] != 0 {
		dispMap[x-1][y+1] = 1
	} 
	if y-1 >= 0 && dispMap[x][y-1] == 0 && mineMap[x][y-1] != 0 {
		dispMap[x][y-1] = 1
	} 
	if y+1 < DIM && dispMap[x][y+1] == 0 && mineMap[x][y+1] != 0 {
		dispMap[x][y+1] = 1
	} 
	if x+1 < DIM && y-1 >= 0 && dispMap[x+1][y-1] == 0 && mineMap[x+1][y-1] != 0 {
		dispMap[x+1][y-1] = 1
	} 
	if x+1 < DIM && dispMap[x+1][y] == 0 && mineMap[x+1][y] != 0 {
		dispMap[x+1][y] = 1
	} 
	if x+1 < DIM && y+1 < DIM && dispMap[x+1][y+1] == 0 && mineMap[x+1][y+1] != 0 {
		dispMap[x+1][y+1] = 1
	}
}

func get_operation () (c string, x int, y int){
	fmt.Println("Please input command, dig, mark or surrend: <d/m/s>:")
	fmt.Scanln(&c)
	if c != "s" {
		fmt.Println("Please input coordinate: <X> <Y>")
		fmt.Scanf("%d %d\n", &x, &y)
	}
	return c, x, y;
}

func show_dispMap () {
	fmt.Println()
	for i := 0; i < DIM; i++ {
		// show X axis index
		if i == 0 {
			for k := 0; k <= DIM; k++ {
				if k == 0 {
					fmt.Printf("\\ ")
				} else {
					fmt.Printf("%d ", k-1)
				}
			}
			fmt.Println();
		}
		// show Y axis index
		fmt.Printf("%d ", i)
		for j := 0; j < DIM; j++ {
			switch dispMap[i][j] {
				case 0:
					fmt.Printf("# ")
				case 1:
					fmt.Printf("%d ", mineMap[i][j])
				case 2:
					fmt.Printf("$ ")
				case 3:
					fmt.Printf("* ")
				default:
			}
		}
		fmt.Println()
	}
}

func init_map () {
	// 1 set all cells in mineMap and dispMap to 0
	for i := 0; i < DIM; i++ {
		for j := 0; j < DIM; j++ {
			mineMap[i][j], dispMap[i][j] = 0, 0;
		}
	}

	// 2 generate the mines coordinate and update the cells in mineMap
	// 2.1 +1 the conjoined cell
	for i := 0; i < MINECNT; i++ {
		x := rand.Intn(DIM)
		y := rand.Intn(DIM)
		// set it to 9 as this is a mine
		if mineMap[x][y] != 9 {
			mineMap[x][y] = 9;
			add_conjoined_cell (x, y)
		} else {
			i--;
		}
	}
}

// cells need updated
// 	x-1, y-1		x-1, y		x-1, y+1
//	x, y-1						x, y+1
//	x+1, y-1		x+1, y		x+1, y+1
func add_conjoined_cell (x int, y int) {
	if x-1 >= 0 && y-1 >= 0 && mineMap[x-1][y-1] != 9 {
		mineMap[x-1][y-1]++
	} 
	if x-1 >= 0 && mineMap[x-1][y] != 9 {
		mineMap[x-1][y]++
	} 
	if x-1 >= 0 && y+1 < DIM && mineMap[x-1][y+1] != 9 {
		mineMap[x-1][y+1]++
	} 
	if y-1 >= 0 && mineMap[x][y-1] != 9 {
		mineMap[x][y-1]++
	} 
	if y+1 < DIM && mineMap[x][y+1] != 9 {
		mineMap[x][y+1]++
	} 
	if x+1 < DIM && y-1 >= 0 && mineMap[x+1][y-1] != 9 {
		mineMap[x+1][y-1]++
	} 
	if x+1 < DIM && mineMap[x+1][y] != 9 {
		mineMap[x+1][y]++
	} 
	if x+1 < DIM && y+1 < DIM && mineMap[x+1][y+1] != 9 {
		mineMap[x+1][y+1]++
	}
}

func show_mineMap () {
	fmt.Println();
	for i := 0; i < DIM; i++ {
		for j := 0; j < DIM; j++ {
			fmt.Printf("%d ", mineMap[i][j])
		}
		fmt.Println()
	}
}