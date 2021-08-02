package dto

type ParsingResult struct {
	Position         uint
	Score            uint
	PriorityOneUpper uint
	BudgetPlaces     uint
}

var EmptyParsingResult ParsingResult
