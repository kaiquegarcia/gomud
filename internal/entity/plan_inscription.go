package entity

import "time"

type PlanInscription struct {
	ID           int       `db:"ID;skip-insert;skip-update"`
	AuthorID     int       `db:"authorID"`
	PlanID       int       `db:"planID"`
	StudentID    int       `db:"studentID"`
	VoucherID    int       `db:"voucherID"`
	ChargeID     int       `db:"chargeID"`
	PlanCost     float64   `db:"plan_cost"`
	PlanDuration string    `db:"plan_duration"`
	Discount     float64   `db:"discount"`
	DiscountTXT  string    `db:"discount_txt"`
	TotalCost    float64   `db:"total_cost"`
	StartedAt    time.Time `db:"start_date" layout:"2006-01-02 15:04:05"`
	FinishedAt   time.Time `db:"end_date" layout:"2006-01-02 15:04:05"`
	Status       int       `db:"status"`
	UpdatedAt    time.Time `db:"status_date" layout:"2006-01-02 15:04:05"`
	CreatedAt    time.Time `db:"registry_date;skip-insert;skip-update" layout:"2006-01-02 15:04:05"`
}
