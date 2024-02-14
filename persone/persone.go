package persone

import (
	"time"
)

type Persone struct {
	Name        string
	Email       string
	dateOfBirth time.Time
}

func NewPersone(name, email string, dobYear, dobMonth, dobDay int) Persone {
	return Persone{
		Name:        name,
		Email:       email,
		dateOfBirth: time.Date(dobYear, time.Month(dobMonth), dobDay, 0, 0, 0, 0, time.UTC),
	}
}
