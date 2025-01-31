package solution25

import (
	"scondon87/advent-ofcode/2024/readinput"
	"strings"
)

type keyLock struct {
	keyLockType string
	heights     []int
	grid        [][]string
}

func solution(filename string) int {
	sum := 0
	keys, locks := parse(filename)
	// fmt.Print(keys, locks)

	for _, key := range keys {
		for _, lock := range locks {
			match := 0
			for i := range key {
				if key[i]+lock[i] <= 5 {
					match++
				}
			}
			if match == 5 {
				sum++
			}
		}
	}

	return sum
}

func parse(filename string) ([][]int, [][]int) {
	input := readinput.Read(filename)
	key := make([][]int, 0)
	lock := make([][]int, 0)

	device := keyLock{}
	for _, line := range input {
		if len(line) != 0 {
			if len(device.grid) == 0 {
				if line == "#####" {
					device.keyLockType = "lock"
				} else if line == "....." {
					device.keyLockType = "key"
				}
			}
			splitLine := strings.Split(line, "")
			device.grid = append(device.grid, splitLine)
		} else {
			// calc partial
			device.calcHeights()
			// commit to data
			if device.keyLockType == "key" {
				key = append(key, device.heights)
			} else {
				lock = append(lock, device.heights)
			}
			// reset partial
			device = keyLock{}
		}
	}

	// handle last set
	device.calcHeights()
	if device.keyLockType == "key" {
		key = append(key, device.heights)
	} else {
		lock = append(lock, device.heights)
	}

	return key, lock
}

func (kl *keyLock) calcHeights() {
	kl.heights = make([]int, 5)
	for _, line := range kl.grid {
		for i, char := range line {
			if char == "#" {
				kl.heights[i] = kl.heights[i] + 1
			}
		}
	}

	// ignore base of key/lock
	for i := range kl.heights {
		kl.heights[i] = kl.heights[i] - 1
	}
}
