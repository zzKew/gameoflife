package main

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten"
)

const scale int = 2
const width = 800
const height = 400

var black = color.Black
var white = color.White
var grid [width][height]uint8 = [width][height]uint8{}
var buffer [width][height]uint8 = [width][height]uint8{}
var count int = 0

func updatewindow() error {
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			buffer[x][y] = 0
			neighbors := grid[x-1][y-1] + grid[x-1][y+0] + grid[x-1][y+1] + grid[x+0][y-1] + grid[x+0][y+1] + grid[x+1][y-1] + grid[x+1][y+0] + grid[x+1][y+1]

			if grid[x][y] == 0 && neighbors == 3 {
				buffer[x][y] = 1
			} else if neighbors < 2 || neighbors > 3 {
				buffer[x][y] = 0
			} else {
				buffer[x][y] = grid[x][y]
			}
		}
	}

	temp := buffer
	buffer = grid
	grid = temp
	return nil
}

func display(window *ebiten.Image) {
	window.Fill(black)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			for i := 0; i < scale; i++ {
				for j := 0; j < scale; j++ {
					if grid[x][y] == 1 {
						window.Set(x*scale+i, y*scale+j, white)
					}
				}
			}
		}
	}
}

func frame(window *ebiten.Image) error {
	count++
	var err error = nil
	if count == 20 {
		err = updatewindow()
		count = 0
	}
	if !ebiten.IsDrawingSkipped() {
		display(window)
	}

	return err
}

func main() {
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if rand.Float32() < 0.5 {
				grid[x][y] = 1
			}
		}
	}
	ebiten.Run(frame, width, height, 2, "Game of Life")
}
