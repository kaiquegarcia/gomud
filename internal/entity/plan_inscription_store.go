package entity

import "time"

type PlanInscriptionStore struct {
	ID            int       `db:"ID;skip-insert;skip-update"`
	PlanID        int       `db:"planID"`
	AuthorID      int       `db:"authorID"`
	Email         string    `db:"email"`
	EmailContent  string    `db:"email_content"`
	DurationPrize string    `db:"duration_prize"`
	Duration      float64   `db:"duration"`
	Charging      int       `db:"charging"`
	Status        int       `db:"status"`
	UpdatedAt     time.Time `db:"status_date" layout:"2006-01-02 15:04:05"`
	CreatedAt     time.Time `db:"registry_date;skip-insert;skip-update" layout:"2006-01-02 15:04:05"`
}
