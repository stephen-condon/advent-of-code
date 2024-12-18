package solution6

import (
	"scondon87/advent-ofcode/2024/readinput"
	"strings"
)

type program struct {
	theMap   levelMap
	theGuard guard
	done     bool
	count    int
}

func (p *program) advance() {
	currentLoc := p.theGuard.theLocation
	nextLoc := p.theGuard.getAdvanceLocation()
	if p.isDone(nextLoc) {
		p.done = true
		return
	}
	obstructed := p.theMap.isObstructed(nextLoc)
	if obstructed {
		// turn
		p.theGuard.turn()
	} else {
		// advance
		p.theGuard.advance(nextLoc)
		p.update(currentLoc, nextLoc)
	}
}

func (p *program) isDone(nextLoc location) bool {
	return nextLoc.x < 0 || nextLoc.y < 0 || nextLoc.x >= len(p.theMap) || nextLoc.y >= len(p.theMap)
}

func (p *program) update(oldLocation, newLocation location) {
	if p.theMap[newLocation.y][newLocation.x] != "X" {
		p.count++
	}
	p.theMap.update(oldLocation, newLocation)
}

type guard struct {
	theLocation location
	theFacing   guardFacing
}

func (g *guard) getAdvanceLocation() location {
	switch g.theFacing {
	case 1:
		newY := g.theLocation.y - 1
		return location{x: g.theLocation.x, y: newY}
	case 2:
		newX := g.theLocation.x + 1
		return location{x: newX, y: g.theLocation.y}
	case 3:
		newY := g.theLocation.y + 1
		return location{x: g.theLocation.x, y: newY}
	case 4:
		newX := g.theLocation.x - 1
		return location{x: newX, y: g.theLocation.y}
	default:
		// should be unreachable
		return location{}
	}
}

func (g *guard) advance(theLocation location) {
	g.theLocation = theLocation
}

func (g *guard) turn() {
	g.theFacing = g.theFacing.turn()
}

func newGuardFacing(direction int) guardFacing {
	return guardFacing(direction)
}

type levelMap [][]string

func (lm levelMap) findGuard() location {
	for y, line := range lm {
		for x, char := range line {
			if char == "^" {
				return location{x, y}
			}
		}
	}

	return location{}
}

func (lm levelMap) isObstructed(theLocation location) bool {
	return lm[theLocation.y][theLocation.x] == "#"
}

func (lm levelMap) update(oldLocation, newLocation location) {
	lm[oldLocation.y][oldLocation.x] = "X"
	lm[newLocation.y][newLocation.x] = "^"
}

type guardFacing int

/*
	1: Up
	2: Right
	3: Down
	4: Left
*/

func (gf guardFacing) turn() guardFacing {
	if gf != 4 {
		gf++
	} else {
		gf = 1
	}

	return gf
}

type location struct {
	x int
	y int
}

func solution(filename string) int {
	input := readinput.Read(filename)
	var myMap = make(levelMap, len(input))

	for i, line := range input {
		chars := strings.Split(line, "")
		myMap[i] = chars
	}

	theProgram := program{
		theMap: myMap,
		theGuard: guard{
			theLocation: myMap.findGuard(),
			theFacing:   newGuardFacing(1),
		},
		done:  false,
		count: 1, // include start position
	}

	for !theProgram.done {
		theProgram.advance()
	}

	return theProgram.count
}
