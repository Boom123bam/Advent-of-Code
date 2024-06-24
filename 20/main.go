package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Tile [][]bool
type Edge []bool
type EdgeMatch struct {
	tileId    int
	dir       int
	needsFlip bool
}
type Image [][]bool

const TILESIZE = 10

const MONSTER_STRING = "                  # \n#    ##    ##    ###\n #  #  #  #  #  #   "

var MONSTER_MATRIX = [][]bool{}

func main() {
	file, err := os.Open("20/input.txt")
	// file, err := os.Open("20/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	tiles := make(map[int]*Tile)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 4 && line[:4] == "Tile" {
			tile_num := extract_nums(line)[0]
			scanner.Scan()
			tile := make(Tile, TILESIZE)
			for i := 0; scanner.Text() != ""; i++ {
				tile[i] = line_to_bool_arr(scanner.Text())
				scanner.Scan()
			}
			tiles[tile_num] = &tile
		}
	}

	p1 := 1
	firstCornerId := -1
	for id := range tiles {
		count := 0
		for i := 0; i < 4; i++ {
			if getMatch(tiles, id, i) != nil {
				count++
			}
		}
		if count == 2 {
			p1 *= id
			if firstCornerId == -1 {
				firstCornerId = id
			}
		}
	}

	fmt.Println(p1)

	corner := tiles[firstCornerId]

	// rotate corner
	for !(getMatch(tiles, firstCornerId, 0) == nil && getMatch(tiles, firstCornerId, 3) == nil) {
		corner.rotateClockwise()
	}

	numTiles := len(tiles)
	size := int(math.Sqrt(float64(numTiles)))

	orderedTilesIds := []int{firstCornerId}
	for i := 1; i < numTiles; i++ {
		if i%size == 0 {
			// first of row
			match := getMatch(tiles, orderedTilesIds[i-size], 2)
			tiles[match.tileId].makeFaceLeft(match.dir)
			if match.needsFlip {
				tiles[match.tileId].flipAlongX()
			}
			tiles[match.tileId].rotateClockwise()
			orderedTilesIds = append(orderedTilesIds, match.tileId)
		} else {
			match := getMatch(tiles, orderedTilesIds[i-1], 1)
			tiles[match.tileId].makeFaceLeft(match.dir)
			if match.needsFlip {
				tiles[match.tileId].flipAlongX()
			}
			orderedTilesIds = append(orderedTilesIds, match.tileId)
		}
	}

	image := Image{}
	for i, tileId := range orderedTilesIds {
		majorRow := i / size
		tile := tiles[tileId]
		for minorRow, tileRow := range (*tile)[1 : TILESIZE-1] {
			totalRow := majorRow*8 + minorRow
			if len(image) <= totalRow {
				image = append(image, []bool{})
			}
			image[totalRow] = append(image[totalRow], tileRow[1:TILESIZE-1]...)
		}
	}

	for _, line := range strings.Split(MONSTER_STRING, "\n") {
		MONSTER_MATRIX = append(MONSTER_MATRIX, line_to_bool_arr(line))
	}

	count := image.orientAndCount()
	imageTileCount := 0
	for _, r := range image {
		for _, c := range r {
			if c {
				imageTileCount++
			}
		}
	}

	monsterTileCount := 0
	for _, r := range MONSTER_MATRIX {
		for _, c := range r {
			if c {
				monsterTileCount++
			}
		}
	}
	fmt.Println(imageTileCount - monsterTileCount*count)
}

func (image Image) orientAndCount() int {
	images := []string{}
	for r := 0; r < 4; r++ {
		count := image.countSeaMonsters()
		if count != 0 {
			return count
		}
		images = append(images, boolMatrixToString(image))
		image.rotateClockwise()
	}
	image.flipAlongX()
	for r := 0; r < 4; r++ {
		count := image.countSeaMonsters()
		if count != 0 {
			return count
		}
		images = append(images, boolMatrixToString(image))
		image.rotateClockwise()
	}
	return -1
}

func (image Image) countSeaMonsters() int {
	count := 0
	for row := 0; row < len(image)-len(MONSTER_MATRIX); row++ {
		for col := 0; col < len(image[row])-len(MONSTER_MATRIX[0]); col++ {
			if image.hasSeaMonsterAt(row, col) {
				count++
			}
		}
	}
	return count
}

func (image Image) hasSeaMonsterAt(row, col int) bool {
	for rowOffset := 0; rowOffset < len(MONSTER_MATRIX); rowOffset++ {
		for colOffset := 0; colOffset < len(MONSTER_MATRIX[rowOffset]); colOffset++ {
			expectTile := MONSTER_MATRIX[rowOffset][colOffset]
			if expectTile && !image[row+rowOffset][col+colOffset] {
				return false
			}
		}
	}
	return true
}

func getMatch(tilesMap map[int]*Tile, targetTileId, targetTileDir int) *EdgeMatch {
	targetEdge := tilesMap[targetTileId].getEdge(targetTileDir)
	for id, tile := range tilesMap {
		for dir := 0; dir < 4; dir++ {
			if id == targetTileId && dir == targetTileDir {
				continue
			}
			edge := tile.getEdge(dir)
			if edge.equals(targetEdge) {
				return &EdgeMatch{
					tileId:    id,
					dir:       dir,
					needsFlip: true,
				}
			} else if edge.flipped().equals(targetEdge) {
				return &EdgeMatch{
					tileId:    id,
					dir:       dir,
					needsFlip: false,
				}
			}
		}
	}
	return nil
}

func (e Edge) flipped() Edge {
	result := make(Edge, TILESIZE)
	for i, item := range e {
		result[TILESIZE-1-i] = item
	}
	return result
}

func (e1 Edge) equals(e2 Edge) bool {
	for i, item := range e1 {
		if item != e2[i] {
			return false
		}
	}
	return true
}

func (t *Tile) makeFaceLeft(dir int) {
	for i := 0; i < 3-dir; i++ {
		t.rotateClockwise()
	}
}

func (t Tile) getEdge(dir int) Edge {
	result := make(Edge, TILESIZE)
	switch dir {
	case 0:
		// north
		return t[0]
	case 1:
		// east
		for i, row := range t {
			result[i] = row[TILESIZE-1]
		}
	case 2:
		// south
		for i, j := 0, TILESIZE-1; i < TILESIZE; i, j = i+1, j-1 {
			result[i] = t[TILESIZE-1][j]
		}
	case 3:
		// west
		for i, j := 0, TILESIZE-1; i < TILESIZE; i, j = i+1, j-1 {
			result[i] = t[j][0]
		}
	default:
		panic(fmt.Sprint("invalid edge: ", dir))
	}
	return result
}

func (e Edge) String() string {
	sb := strings.Builder{}
	for _, c := range e {
		if c {
			sb.WriteByte('#')
		} else {
			sb.WriteByte('.')
		}
	}
	return sb.String()
}

func (i *Image) rotateClockwise() {
	newImage := Image(rotatedClockwise([][]bool(*i)))
	i = &newImage
}

func (i *Image) flipAlongX() {
	newImage := Image(flippedAlongX([][]bool(*i)))
	i = &newImage
}

func (i *Image) flipAlongY() {
	i.rotateClockwise()
	i.flipAlongX()
	i.rotateClockwise()
	i.rotateClockwise()
	i.rotateClockwise()
}

func (t *Tile) rotateClockwise() {
	newTile := Tile(rotatedClockwise([][]bool(*t)))
	t = &newTile
}

func (t *Tile) flipAlongX() {
	newTile := Tile(flippedAlongX([][]bool(*t)))
	t = &newTile
}

func rotatedClockwise(matrix [][]bool) [][]bool {
	// Transpose the matrix
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Reverse the rows
	for i := 0; i < len(matrix); i++ {
		for j, k := 0, len(matrix[i])-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
	return matrix

}

func flippedAlongX(matrix [][]bool) [][]bool {
	for i, j := 0, len(matrix)-1; i < len(matrix)/2; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}
	return matrix
}

func (t Tile) String() string {
	sb := strings.Builder{}
	sb.WriteByte('\n')
	for _, row := range t {
		for _, col := range row {
			if col {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func boolMatrixToString(matrix [][]bool) string {
	sb := strings.Builder{}
	sb.WriteByte('\n')
	for _, row := range matrix {
		for _, col := range row {
			if col {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func line_to_bool_arr(line string) []bool {
	result := make([]bool, len(line))
	for i, char := range line {
		if char == '#' {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	return result
}

func extract_nums(line string) []int {
	result := []int{}
	sum := 0
	prev_num := false
	for _, char := range line {
		num, err := strconv.Atoi(string(char))
		if err == nil {
			sum = sum*TILESIZE + num
			prev_num = true
			continue
		}
		if prev_num {
			result = append(result, sum)
		}
		prev_num = false
		sum = 0
	}
	if prev_num {
		result = append(result, sum)
	}
	return result
}
