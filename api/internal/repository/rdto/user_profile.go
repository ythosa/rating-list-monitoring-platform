package rdto

type UserProfile struct {
	Username   string `json:"username" db:"username"`
	FirstName  string `json:"first_name" db:"first_name"`
	MiddleName string `json:"middle_name" db:"middle_name"`
	LastName   string `json:"last_name" db:"last_name"`
	Snils      string `json:"snils" db:"snils"`
}
