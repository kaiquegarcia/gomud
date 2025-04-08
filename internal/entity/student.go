package entity

import "time"

type Student struct {
	ID            int       `db:"ID;skip-insert;skip-update"`
	InstitutionID int       `db:"institutionID"`
	Login         string    `db:"login"`
	Email         string    `db:"email"`
	Name          string    `db:"name"`
	Password      string    `db:"password"`
	Iam           int       `db:"Iam"`
	Status        int       `db:"status"`
	UpdatedAt     time.Time `db:"status_date" layout:"2006-01-02 15:04:05"`
	CreatedAt     time.Time `db:"registry_date;skip-insert;skip-update" layout:"2006-01-02 15:04:05"`
}
