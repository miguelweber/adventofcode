package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"
)

const (
	HighCard     Rank = iota
	OnePair      Rank = iota
	TwoPair      Rank = iota
	ThreeOfAKind Rank = iota
	FullHouse    Rank = iota
	FourOfAKind  Rank = iota
	FiveOfAKind  Rank = iota
)

type Card = int8
type Rank = int8
type Hand struct {
	cards [5]Card
	rank  Rank
	bid   int32
}

func Btoi(bs []byte) int {
	result := 0
	for _, b := range bs {
		result = result*10 + int(b) - '0'
	}
	return result
}

func GetCard(char byte) Card {
	if char >= '2' && char <= '9' {
		return int8(char) - '2' + 1
	}
	switch char {
	case 'T':
		return 9
	case 'J':
		return 0
	case 'Q':
		return 10
	case 'K':
		return 11
	case 'A':
		return 12
	}
	return -1
}

func GetRank(cards [5]Card) Rank {
	cardCount := [13]int8{}

	for _, card := range cards {
		cardCount[card]++
	}
	maxI, max := 0, int8(0)
	for i, card := range cardCount[1:] {
		if card > max {
			max = card
			maxI = i + 1
		}
	}
	if maxI == 0 {
		return FiveOfAKind
	}
	cardCount[maxI] += cardCount[0]
	pairCount := 0
	has3Equal := false

	for _, count := range cardCount[1:] {
		switch count {
		case 2:
			pairCount++
		case 3:
			has3Equal = true
		case 4:
			return FourOfAKind
		case 5:
			return FiveOfAKind
		}
	}

	switch {
	case has3Equal && pairCount > 0:
		return FullHouse
	case has3Equal:
		return ThreeOfAKind
	case pairCount == 2:
		return TwoPair
	case pairCount == 1:
		return OnePair
	default:
		return HighCard
	}
}

func main() {
	filename := flag.String("i", "input", "input file")
	flag.Parse()
	buf, _ := os.ReadFile(*filename)

	hands := []Hand{}

	for len(buf) > 0 {
		end := bytes.IndexByte(buf, '\n')
		if end < 0 {
			break
		}

		hand := Hand{bid: int32(Btoi(buf[bytes.IndexByte(buf, ' ')+1 : end]))}
		for i := 0; i < 5; i++ {
			hand.cards[i] = GetCard(buf[i])
		}
		hand.rank = GetRank(hand.cards)
		hands = append(hands, hand)

		buf = buf[end+1:]
	}
	buf = nil // GC it

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].rank == hands[j].rank {
			for card := 0; card < 5; card++ {
				cardA, cardB := hands[i].cards[card], hands[j].cards[card]
				if cardA == cardB {
					continue
				}
				return cardA < cardB
			}
		}
		return hands[i].rank < hands[j].rank
	})

	result := 0
	for i, hand := range hands {
		result += (i + 1) * int(hand.bid)
	}

	fmt.Println(result)
}
