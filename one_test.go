package adventOfCodeTest

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
	"testing"
)

const RIGHT = 1
const LEFT = -1

const NORTH = 0
const EAST = 1
const SOUTH = 2
const WEST = 3

type state struct {
	direction int
	x         int
	y         int
}

type step struct {
	turn     int
	distance int
}

func newStep(turn int, distance int) step {
	return step{
		turn:     turn,
		distance: distance,
	}
}

func newState() *state {
	return &state{
		direction: NORTH,
		x:         0,
		y:         0,
	}
}

func (state *state) move(step step) *state {
	state.direction = (state.direction + step.turn + 4) % 4
	switch state.direction {
	case NORTH:
		state.y += step.distance
	case SOUTH:
		state.y -= step.distance
	case EAST:
		state.x += step.distance
	case WEST:
		state.x -= step.distance
	}
	return state
}

func (state *state) checkMap(mapToCheck string) int {
	return state.check(strings.Split(mapToCheck, ", "))
}

func (state *state) check(steps []string) int {
	for _, v := range steps {
		distance, _ := strconv.Atoi(string(v[1:len(string(v))]))
		direction := string(v[0])
		var turn int
		if direction == "R" {
			turn = RIGHT
		} else {
			turn = LEFT
		}
		step := newStep(turn, distance)
		state = state.move(step)
	}
	return int(math.Abs(float64(state.x)) + math.Abs(float64(state.y)))
}

func Test_when_r2_l3(t *testing.T) {
	state := newState()
	result := state.check([]string{"R2", "L3"})
	expected := 5
	if result != expected {
		t.Fatalf("The result %v is expected to be %v", result, expected)
	}
}

func Test_when_r2_r2_r2(t *testing.T) {
	state := newState()
	result := state.check([]string{"R2", "R2", "R2"})
	expected := 2
	if result != expected {
		t.Fatalf("The result %v is expected to be %v", result, expected)
	}
}

func Test_when_r5_l5_r5_r3(t *testing.T) {
	state := newState()
	result := state.check([]string{"R5", "L5", "R5", "R3"})
	expected := 12
	if result != expected {
		t.Fatalf("The result %v is expected to be %v", result, expected)
	}
}

func Test_when_map_from_advent_of_code(t *testing.T) {
	state := newState()
	dat, err := ioutil.ReadFile("day1_part1.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	result := state.checkMap(string(dat))
	log.Printf("The answer to advent of code is %v", result)
}
