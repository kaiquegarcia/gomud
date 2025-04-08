package entity

import "time"

type Institution struct {
	ID          int       `db:"ID;skip-insert;skip-update"`
	AuthorID    int       `db:"userID"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"registry_date;skip-insert;skip-update" layout:"2006-01-02 15:04:05"`
}
