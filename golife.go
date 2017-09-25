package main

import (
	"fmt"
	"bytes"
	"sync"
	"time"
)

type Grid struct {
	squares    [][]bool
	width, height int
}

// Checks to see if a given square is alive.  The check is wrapped so that width+1 = 0 and likewise -1 = width-1
func (g *Grid) Alive(x, y int) bool {
	x += g.width
	x %= g.width
	y += g.height
	y %= g.height
	return g.squares[x][y]
}

// Tells you if the square will be alive next generation
func (g *Grid) Next(x,y int) bool {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && g.Alive(x+i, y+j) {
				alive++
			}
		}
	}

	// The cell is dead unless there are exactly 3 neighbors or it is already alive and has 2 neighbors
	return alive == 3 || alive == 2 && g.Alive(x, y)
}

// Create an Empty Grid
func EmptyGrid(width, height int) *Grid {
	squares := make([][]bool, height)
	for i := range squares {
		squares[i] = make([]bool, width)
	}
	return &Grid{squares: squares, width: width, height: height}
}

// initialize a Glider in the middle of an otherwise empty grid
func GliderGrid(width, height int) *Grid {
	grid := EmptyGrid(width, height)
	if width > 4 && height > 4 {
		var center_x int = width / 2
		var center_y int = height / 2
		grid.squares[center_x][center_y-1] = true
		grid.squares[center_x+1][center_y] = true
		grid.squares[center_x-1][center_y+1] = true
		grid.squares[center_x][center_y+1] = true
		grid.squares[center_x+1][center_y+1] = true
	}

	return grid
}

// Generate a new grid for the next generation
func nextGen(g *Grid) *Grid {
	newGrid := EmptyGrid(g.width, g.height)
	// Create a wait group so that the program will wait for the entire grid to be calculated before printing
	var wg sync.WaitGroup
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			wg.Add(1)
			newGrid.threadAssignment(g, x, y, &wg)
		}
	}
	wg.Wait()
	return newGrid
}

// Go requires all go routines to be a separate function
func (g *Grid) threadAssignment(prev *Grid, x,y int, wg *sync.WaitGroup) {
	g.squares[x][y] = prev.Next(x,y)
	wg.Done()
	return
}

// Print the Grid
func (g *Grid) String() string {
	var buf bytes.Buffer
	for i := 0; i < g.width+2; i++ {
		buf.WriteByte('-')
	}
	buf.WriteByte('\n')

	for i := 0; i < g.width; i++ {
		buf.WriteByte('|')
		for j := 0; j < g.height; j++ {
			b := byte(' ')
			if g.squares[j][i] {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('|')
		buf.WriteByte('\n')
	}

	for i := 0; i < g.width+2; i++ {
		buf.WriteByte('-')
	}
	buf.WriteByte('\n')
	return buf.String()
}

func main() {
	g := GliderGrid(25, 25)
	fmt.Print("\x0c", g)
	genCount := 0

	for genCount < 1000 {
		fmt.Println("Generating next generation in 5 seconds")
		time.Sleep(5 * time.Second)
		g = nextGen(g)
		fmt.Print("\x0c", g)
		genCount += 1
	}
}
