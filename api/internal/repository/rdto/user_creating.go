package rdto

type UserCreating struct {
	Username   string `db:"username"`
	Password   string `db:"password"`
	FirstName  string `db:"first_name"`
	MiddleName string `db:"middle_name"`
	LastName   string `db:"last_name"`
	Snils      string `db:"snils"`
}
