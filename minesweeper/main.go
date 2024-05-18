package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Cell struct {
	isBomb      bool
	isUncovered bool
	symbol      *string
}

func (cell *Cell) String(reveal bool) string {
	var character string

	if cell.symbol != nil {
		return *cell.symbol
	}

	if cell.isBomb && reveal {
		character = "⦿"
	} else if cell.isUncovered {
		character = " "
	} else {
		character = "☐"
	}

	return character
}

type Grid struct {
	width  int
	height int
	bombs  int
	plays  int
	cells  [][]Cell
}

const (
	GRID_PLAY_OUTSIDE_BOUNDARIES     = iota
	GRID_PLAY_CELL_ALREADY_UNCOVERED = iota
	GRID_PLAY_FOUND_BOMB             = iota
	GRID_PLAY_WIN                    = iota
	GRID_PLAY_CONTINUE               = iota
)

func (grid *Grid) Print(reveal bool) {
	fmt.Printf("-\t")
	for i := 1; i < grid.width+1; i++ {
		if i < 10 {
			fmt.Printf("%d  ", i)
		} else {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")

	for i, row := range grid.cells {
		fmt.Printf("%d\t", i+1)
		for _, cell := range row {
			fmt.Printf("%s  ", cell.String(reveal))
		}
		fmt.Printf("\n")
	}
}

func (grid *Grid) UncoverCell(row int, col int) int {
	if !grid.checkBounds(row, col) {
		return 0
	}

	if grid.cells[row][col].isUncovered || grid.cells[row][col].isBomb {
		return 0
	}

	grid.cells[row][col].isUncovered = true

	adjacentBombs := grid.adjacentBombs(row, col)

	returnVal := 1
	if adjacentBombs == 0 {
		returnVal += grid.UncoverCell(row-1, col)
		returnVal += grid.UncoverCell(row+1, col)
		returnVal += grid.UncoverCell(row, col-1)
		returnVal += grid.UncoverCell(row, col+1)
	}

	return returnVal
}

func (grid *Grid) checkBounds(row int, col int) bool {
	return !(row < 0 || row >= grid.height || col < 0 || col >= grid.width)
}

func (grid *Grid) adjacentBombs(row int, col int) int {
	if !grid.checkBounds(row, col) {
		return 0
	}

	adjacentBombs := 0
	adjacentBombs += grid.HasBomb(row-1, col)
	adjacentBombs += grid.HasBomb(row+1, col)
	adjacentBombs += grid.HasBomb(row, col-1)
	adjacentBombs += grid.HasBomb(row, col+1)

	if adjacentBombs > 0 {
		adjacentBombsStr := fmt.Sprintf("%d", adjacentBombs)
		grid.cells[row][col].symbol = &adjacentBombsStr
	}

	return adjacentBombs
}

func (grid *Grid) HasBomb(row int, col int) int {
	if !grid.checkBounds(row, col) {
		return 0
	}
	if grid.cells[row][col].isBomb {
		return 1
	}
	return 0
}

func (grid *Grid) Play(row int, col int) int {

	if row >= grid.height || col >= grid.width {
		return GRID_PLAY_OUTSIDE_BOUNDARIES
	}

	if grid.cells[row][col].isUncovered {
		return GRID_PLAY_CELL_ALREADY_UNCOVERED
	}

	if grid.cells[row][col].isBomb {
		var bombSymbol string = "\x1b[31mX\033[0m"
		grid.cells[row][col].symbol = &bombSymbol
		return GRID_PLAY_FOUND_BOMB
	}

	grid.plays += grid.UncoverCell(row, col)

	if grid.plays == grid.width*grid.height-grid.bombs {
		return GRID_PLAY_WIN
	}

	return GRID_PLAY_CONTINUE
}

func createGrid(width int, height int, bombs int) *Grid {
	cells := make([][]Cell, height)
	for i := range cells {
		cells[i] = make([]Cell, width)
	}

	// initialize grid with blank cells
	for row := range cells {
		for col := range cells[row] {
			cells[row][col] = Cell{isBomb: false, isUncovered: false, symbol: nil}
		}
	}

	assignedBombs := 0

	// generates random bomb locations
	for assignedBombs < bombs {
		newLocation := rand.Int() % (width * height)
		row := int(newLocation / width)
		col := newLocation % height
		if cells[row][col].isBomb {
			continue
		}
		cells[row][col].isBomb = true
		assignedBombs++
	}

	return &Grid{cells: cells, width: width, height: height, bombs: bombs, plays: 0}
}

func main() {
	var gridWidth, gridHeight, bombAmount int

	fmt.Printf("Grid width: ")

	_, err := fmt.Scanf("%d", &gridWidth)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Grid height: ")
	_, err = fmt.Scanf("%d", &gridHeight)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Bombs: ")
	_, err = fmt.Scanf("%d", &bombAmount)
	if err != nil {
		log.Fatal(err)
		return
	}

	if bombAmount >= gridHeight*gridWidth-1 {
		log.Fatal("Can't have more bombs than grid size")
		return
	}

	grid := createGrid(gridWidth, gridHeight, bombAmount)

	start := time.Now()
	for {
		var playRow, playCol int
		grid.Print(false)
		fmt.Printf("Select square (row, col): ")
		fmt.Scanf("%d,%d", &playRow, &playCol)
		playOutcome := grid.Play(playRow-1, playCol-1)
		switch playOutcome {
		case GRID_PLAY_OUTSIDE_BOUNDARIES:
			{
				fmt.Println("\x1b[31mPlease select a cell within grid boundaries\033[0m")
				continue
			}
		case GRID_PLAY_CELL_ALREADY_UNCOVERED:
			{
				fmt.Println("\x1b[31mPlease select a cell that isn't already uncovered\033[0m")
				continue
			}
		case GRID_PLAY_FOUND_BOMB:
			{
				fmt.Println("\x1b[31mYou lose!\033[0m")
				break
			}
		case GRID_PLAY_WIN:
			{
				fmt.Println("\x1b[32mYou win!\033[0m")
				break
			}
		}

		if playOutcome == GRID_PLAY_FOUND_BOMB || playOutcome == GRID_PLAY_WIN {
			break
		}

	}
	end := time.Now()
	grid.Print(true)
	fmt.Printf("Elapsed time: %s\n", end.Sub(start).String())
}
