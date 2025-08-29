package pagination

import (
	"math"
)

// Config is used to calculate offset and limit
type Config struct {
	Page   int
	Offset int
	Limit  int
}

// Paginator represents pagination info in API response
type Paginator struct {
	TotalItems   int  `json:"total_items"`
	TotalPages   int  `json:"total_pages"`
	ItemFrom     int  `json:"item_from"`
	ItemTo       int  `json:"item_to"`
	CurrentPage  int  `json:"current_page"`
	Limit        int  `json:"limit"`
	NextPage     *int `json:"next_page,omitempty"`
	PreviousPage *int `json:"previous_page,omitempty"`
}

// GetPaginationConfig returns offset and limit based on page
func GetPaginationConfig(page int, limit int) Config {
	if page <= 0 {
		page = 1 // default page
	}
	if limit <= 0 {
		limit = 30 // default limit
	}
	offset := (page - 1) * limit

	return Config{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
}

// BuildPaginator constructs paginator info
func BuildPaginator(total, limit, offset int) *Paginator {
	currentPage := offset/limit + 1
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	itemFrom := validateItem(offset+1, total)
	itemTo := validateItem(offset+limit, total)

	var nextPage *int
	var prevPage *int
	if currentPage < totalPages {
		n := currentPage + 1
		nextPage = &n
	}
	if currentPage > 1 {
		p := currentPage - 1
		prevPage = &p
	}

	return &Paginator{
		TotalItems:   total,
		TotalPages:   totalPages,
		ItemFrom:     itemFrom,
		ItemTo:       itemTo,
		CurrentPage:  currentPage,
		Limit:        limit,
		NextPage:     nextPage,
		PreviousPage: prevPage,
	}
}

// validateItem ensures the item index is within total items
func validateItem(number, total int) int {
	if number <= 0 {
		return 0
	}
	if number > total {
		return total
	}
	return number
}
