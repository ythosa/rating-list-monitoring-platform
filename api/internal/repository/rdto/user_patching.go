package rdto

type UserPatching struct {
	FirstName  *string `db:"first_name"`
	MiddleName *string `db:"middle_name"`
	LastName   *string `db:"last_name"`
	Snils      *string `db:"snils"`
}
