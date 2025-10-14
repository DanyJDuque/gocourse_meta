package meta

import (
	"strconv"
)

type (
	Meta struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		PageCount  int `json:"page_count"`
		TotalCount int `json:"tolat_count"`
	}
)

func New(page, perPage, total int, pageLimDef string) (*Meta, error) {

	if perPage <= 0 {
		var err error
		perPage, err = strconv.Atoi(pageLimDef)
		if err != nil {
			return nil, err
		}
	}

	pageCount := 0
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > perPage {
			page = pageCount
		}
	}

	if page < 1 {
		page = 1
	}

	return &Meta{
		Page:       page,
		PerPage:    perPage,
		TotalCount: total,
		PageCount:  pageCount,
	}, nil
}

func (p *Meta) Offset() int {
	return (p.Page - 1) * p.PerPage
}

func (p *Meta) Limit() int {
	return p.PerPage
}
