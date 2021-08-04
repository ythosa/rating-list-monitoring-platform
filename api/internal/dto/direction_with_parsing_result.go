package dto

import "github.com/ythosa/rating-list-monitoring-platform-api/internal/repository/rdto"

type DirectionWithParsingResult struct {
	Direction     rdto.Direction `json:"direction"`
	ParsingResult ParsingResult  `json:"parsing_result"`
}
