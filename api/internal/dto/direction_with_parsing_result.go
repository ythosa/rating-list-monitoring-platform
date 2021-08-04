package dto

import "github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository/rdto"

type DirectionWithParsingResult struct {
	Direction     rdto.Direction `json:"direction"`
	ParsingResult ParsingResult  `json:"parsing_result"`
}
