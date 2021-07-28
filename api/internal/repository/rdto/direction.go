package rdto

type Direction struct {
	DirectionID    uint   `json:"direction_id" db:"direction_id"`
	DirectionName  string `json:"direction_name" db:"direction_name"`
	UniversityID   uint   `json:"university_id" db:"university_id"`
	UniversityName string `json:"university_name" db:"university_name"`
}
