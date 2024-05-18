package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"

	"golang.org/x/exp/slices"
)

type Cell struct {
	isMine bool
}

func (cell *Cell) String() string {
	var character string
	if cell.isMine {
		character = "X"
	} else {
		character = "☐"
	}
	return character
}

type Grid struct {
	width  int
	height int
	cells  [][]Cell
}

func (grid *Grid) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("-\t")
	for i := 1; i < grid.width+1; i++ {
		if i < 10 {
			buffer.WriteString(fmt.Sprintf("%d  ", i))
		} else {
			buffer.WriteString(fmt.Sprintf("%d ", i))
		}
	}
	buffer.WriteByte('\n')

	for i, row := range grid.cells {
		buffer.WriteString(fmt.Sprintf("%d\t", i+1))
		for _, cell := range row {
			buffer.WriteString(cell.String())
			buffer.WriteString("  ")
		}
		buffer.WriteByte('\n')
	}

	return buffer.String()
}

func createGrid(width int, height int, bombs int) *Grid {
	cells := make([][]Cell, height)
	for i := range cells {
		cells[i] = make([]Cell, width)
	}

	assignedBombs := 0
	bombLocation := make([]int, bombs)

	for assignedBombs < bombs {
		newLocation := rand.Int() % (width * height)
		if slices.Contains(bombLocation, newLocation) {
			continue
		}
		bombLocation[assignedBombs] = newLocation
		assignedBombs++
	}

	for row := range cells {
		for col := range cells[row] {
			cells[row][col] = Cell{isMine: false}
		}
	}
	return &Grid{cells: cells, width: width, height: height}
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
	fmt.Println(grid.String())

}
