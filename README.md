# conway-go
Conway-Go is a simple implementation of Conway's Game of Life. While I'm primarily a Back-End developer, I stumbled upon this fascinating game and decided to add it to my portfolio.

## Concept
See [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) to understand the game dynamics. In essence, it's a 2-D world where each cell is either alive or dead, and its state evolves based on the cells surrounding it over time.

### Game Rules:
- Any live cell with fewer than two live neighbours dies, as if by underpopulation.
- Any live cell with two or three live neighbours lives on to the next generation.
- Any live cell with more than three live neighbours dies, as if by overpopulation.
- Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

## Implementation Overview
The screen is represented by a 2-D array of integers, using only ones and zeros. I employ the [ebiten](https://github.com/hajimehoshi/ebiten) library to render a GUI based on the calculated matrix. Each black pixel corresponds to a living cell (1 in the matrix), while white spaces represent dead cells (0 in the matrix).

The game initializes with a random grid. Ebiten updates the game 60 times per second. Each tick updates the grid's state, clears the screen, and represents every living cell with a black pixel.

## Running Game
1. Clone the repo: `it clone https://github.com/FdHerrera/conway-go.git`

2. `cd conway-go`

3. `go run .`


https://github.com/FdHerrera/conway-go/assets/84097984/83b40ed1-5b22-45ee-8287-e5a76cf96619
