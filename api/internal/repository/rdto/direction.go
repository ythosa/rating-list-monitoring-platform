package rdto

type Direction struct {
	DirectionID        uint   `db:"direction_id"`
	DirectionName      string `db:"direction_name"`
	DirectionURL       string `db:"direction_url"`
	UniversityID       uint   `db:"university_id"`
	UniversityName     string `db:"university_name"`
	UniversityFullName string `db:"university_full_name"`
}
