package ratingparser

type ParsingResult struct {
	Position         uint
	Score            uint
	PriorityOneUpper uint
	BudgetPlaces     uint
}

var EmptyResult = ParsingResult{}
