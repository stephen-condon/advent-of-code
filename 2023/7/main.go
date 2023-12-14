package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	result2 := solve("input.txt")

	log.Printf(`Result 2: %v`, result2)
}

func solve(filename string) int {
	data := readInput(filename)

	result2 := processGame(data)

	return result2
}

func processGame(data []string) int {
	var gameHands GameHands

	for _, line := range data {
		hand, bid, _ := strings.Cut(line, " ")

		handCounts := buildHandCounts(hand)
		handType := evaluateTricks(handCounts)

		bidValue, _ := strconv.Atoi(bid)

		gameHand := GameHand{
			rawHand:    hand,
			handCounts: handCounts,
			handType:   handType,
			bid:        bidValue,
		}

		gameHands = append(gameHands, gameHand)
	}

	gameHands.sort()
	result := gameHands.calculateGameScore()

	return result
}

func buildHandCounts(hand string) HandCounts {
	handCounts := HandCounts{}
	splitHand := strings.Split(hand, "")
	for _, card := range splitHand {
		_, success := handCounts[card]
		if success {
			handCounts[card] = handCounts[card] + 1
		} else {
			handCounts[card] = 1
		}
	}
	return handCounts
}

func evaluateTricks(handCounts HandCounts) int {
	hasCount := []bool{false, false, false, false, false} // 5, 4, 3, 2, 2nd 2
	secondKey := ""
	largestCount := 0
	log.Println(handCounts)
	for key, count := range handCounts {
		log.Printf(`key: %v; count: %v`, key, count)
		if key != "J" {
			if count > largestCount {
				if count == 3 && largestCount == 2 {
					secondKey = " "
				}
				largestCount = count
			} else if count == 2 && largestCount == 3 {
				secondKey = " "
			} else if count == largestCount && count == 2 {
				secondKey = key
			}
		}
	}
	log.Println(secondKey)
	newCount := largestCount + handCounts["J"]
	if newCount == 5 {
		hasCount[0] = true
	} else if newCount == 4 {
		hasCount[1] = true
	} else if newCount == 3 {
		hasCount[2] = true
		if secondKey != "" {
			hasCount[3] = true
		}
	} else if newCount == 2 {
		if secondKey != "" {
			hasCount[4] = true
		} else {
			hasCount[3] = true
		}
	}
	log.Println(hasCount)

	return assignRankValue(hasCount)
}

func assignRankValue(hasCount []bool) int {
	rank := 0
	// 50, 40, 32, 30, 22, 20 | 5 cards, 4 cards, full house, 3 cards, 2 pairs, 1 pair
	if hasCount[0] {
		rank = 50
	} else if hasCount[1] {
		rank = 40
	} else if hasCount[2] {
		if hasCount[3] {
			rank = 32
		} else {
			rank = 30
		}
	} else if hasCount[3] {
		rank = 20
	} else if hasCount[4] {
		rank = 22
	}

	return rank
}

func readInput(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	var dataSlice []string

	for scanner.Scan() {
		text := scanner.Text()
		dataSlice = append(dataSlice, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dataSlice
}

type GameHands []GameHand

func (g GameHands) sort() {
	slices.SortFunc(g, compareGameHand)
}

func (g GameHands) calculateGameScore() int {
	total := 0
	for index, hand := range g {
		log.Println(hand)
		handScore := (index + 1) * hand.bid
		total = total + handScore
	}

	return total
}

type GameHand struct {
	rawHand    string
	handCounts HandCounts
	handType   int
	bid        int
}

type HandCounts map[string]int

func compareGameHand(a GameHand, b GameHand) int {
	if a.handType == b.handType {
		return compareCards(a.rawHand, b.rawHand)
	}
	return a.handType - b.handType
}

func compareCards(a string, b string) int {
	aRaw := strings.Split(a, "")
	bRaw := strings.Split(b, "")
	for index := range aRaw {
		for _, card := range getCardOrder() {
			if aRaw[index] == card && bRaw[index] != card {
				return 1
			} else if aRaw[index] != card && bRaw[index] == card {
				return -1
			}
		}
	}

	return 0
}

func getCardOrder() []string {
	return []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}
}
