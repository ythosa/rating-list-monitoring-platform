package dto

type DirectionWithRating struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Position         uint   `json:"position"`
	Score            uint   `json:"score"`
	PriorityOneUpper uint   `json:"priority_one_upper"`
	BudgetPlaces     uint   `json:"budget_places"`
}
