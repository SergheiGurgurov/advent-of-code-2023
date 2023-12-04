package main

import util "advent-of-code-2023"

//this is a failed implementation of day4 part 2:
//
//I tried actually creating a linked list that would grow with each new card gained
//scanning it once from left to right
//while inserting items inside of it each time you would gain a new ScratchCard
//unfortunately the list grows insanely fast O(e^n)
//so i stopped the execution of the program after 5 minutes
//the logic is probably correct tho.

func part_2_tooLong(cards *[]ScratchCard) int {
	cardList := util.ToLinkedList(util.Map(*cards, func(card ScratchCard) int {
		return card.cardID
	}))

	count := 0
	maybeHead := cardList.Shift()
	for {
		if !maybeHead.IsSome {
			break
		}
		head := maybeHead.Value
		count++
		newCards := part_2_tooLong_getCardPoints(CardById(cards, head.Value))
		MergeLists(cardList, newCards)
		maybeHead = cardList.Shift()
	}

	cardList.Print()
	return count
}

func part_2_tooLong_getCardPoints(card ScratchCard) *util.LinkedList[int] {
	score := 0
	for _, wn := range card.winningNumbers {
		for _, yn := range card.yourNumbers {
			if wn == yn {
				score++
			}
		}
	}

	ll := util.LinkedList[int]{}

	for i := 1; i <= score; i++ {
		ll.Add(card.cardID + i)
	}

	return &ll
}

func CardById(cards *[]ScratchCard, id int) ScratchCard {
	return (*cards)[id-1]
}

func MergeLists[T comparable](mainList *util.LinkedList[T], toMerge *util.LinkedList[T]) {
	mainListHead := mainList.Head
	toMergeHead := toMerge.Head

	for {
		if toMergeHead == nil {
			return
		}

		if mainListHead.Value == toMergeHead.Value {
			nextToMerge := toMergeHead.Next
			util.InsertNodeAfter(mainListHead, toMergeHead)
			toMergeHead = nextToMerge
		} else {
			mainListHead = mainListHead.Next
		}
	}
}
