package db

type Pagination struct {
	Offset int
	Limit  int
}

const (
	DefaultPaginationLimit = 25
	MaxPaginationLimit     = 1000
)

func (p *Pagination) Sanitize() {
	if p.Offset < 0 {
		p.Offset = 0
	}

	if p.Limit <= 0 {
		p.Limit = DefaultPaginationLimit
	} else if p.Limit > MaxPaginationLimit {
		p.Limit = MaxPaginationLimit
	}
}
