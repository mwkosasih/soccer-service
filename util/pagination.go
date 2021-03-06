package util

import (
	"math"
	"strconv"
)

const (
	defaultPage  = 1
	DefaultCount = 10
	limitPage    = 99999
)

type Pagination struct {
	CurrentPage   int32 `json:"current_page"`
	PageSize      int32 `json:"page_size"`
	TotalPage     int32 `json:"total_page"`
	TotalResult   int32 `json:"total_result"`
	HasPagination bool  `json:"has_pagination"`
}

// NewPagination initiate new pagination obj
func NewPagination(page int32, count int32) *Pagination {
	var p, c int32
	if page == 0 {
		p = defaultPage
	} else if page > 0 {
		p = page
	}

	if count == 0 {
		c = DefaultCount
	} else if count > 0 {
		c = count
	} else if count < 0 {
		c = limitPage
	}

	pagination := &Pagination{
		CurrentPage:   p,
		PageSize:      c,
		TotalResult:   0,
		TotalPage:     0,
		HasPagination: c > 0,
	}

	return pagination
}

// SetTotalPage calculate total page by total count data
func (p *Pagination) SetTotalPage(total int32) *Pagination {
	totalPages := int32(1)

	if p.PageSize == limitPage {
		p.PageSize = total
	} else {
		d := float64(total) / float64(p.PageSize)
		totalPages = int32(math.Ceil(d))
	}

	p.TotalPage = totalPages
	p.TotalResult = total
	return p
}

// Limit get limit
func (p *Pagination) Limit() int32 {
	return p.PageSize
}

// Offset get Offset
func (p *Pagination) Offset() int32 {
	return (p.CurrentPage - 1) * p.PageSize
}

// LimitOffset generate limit and offset for pagination
func (p *Pagination) LimitOffset() string {
	if p.PageSize < 0 {
		return " "
	}
	l := strconv.Itoa(int(p.Limit()))
	o := strconv.Itoa(int(p.Offset()))
	return ` LIMIT ` + l + ` OFFSET ` + o
}
