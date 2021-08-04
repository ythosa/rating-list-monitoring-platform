package dto

type UniversityDirectionsWithRating struct {
	UniversityID   uint                  `json:"university_id"`
	UniversityName string                `json:"university_name"`
	Directions     []DirectionWithRating `json:"directions"`
}
