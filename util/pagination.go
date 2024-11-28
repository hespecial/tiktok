package util

type Pagination struct {
	Page   int `form:"page" json:"page"`
	Limit  int `form:"size" json:"size"`
	Offset int
}

func (p *Pagination) Format() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 10
	}
	p.Offset = (p.Page - 1) * p.Limit
}
