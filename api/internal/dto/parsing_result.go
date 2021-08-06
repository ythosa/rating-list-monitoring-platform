package dto

type ParsingResult struct {
	Position              uint
	Score                 uint
	PriorityOneUpper      uint
	SubmittedConsentUpper uint
	BudgetPlaces          uint
}

var EmptyParsingResult ParsingResult
