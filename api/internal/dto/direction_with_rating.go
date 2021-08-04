package dto

type DirectionWithRating struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Position         uint   `json:"position"`
	Score            uint   `json:"score"`
	PriorityOneUpper uint   `json:"priority_one_upper"`
	BudgetPlaces     uint   `json:"budget_places"`
}

func NewDirectionWithRating(d DirectionWithParsingResult) DirectionWithRating {
	return DirectionWithRating{
		ID:               d.Direction.DirectionID,
		Name:             d.Direction.DirectionName,
		Position:         d.ParsingResult.Position,
		Score:            d.ParsingResult.Score,
		PriorityOneUpper: d.ParsingResult.PriorityOneUpper,
		BudgetPlaces:     d.ParsingResult.BudgetPlaces,
	}
}
