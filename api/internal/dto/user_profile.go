package dto

type UserProfile struct {
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Snils      string `json:"snils"`
}
