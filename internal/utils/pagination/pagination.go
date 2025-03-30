package pagination

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// Meta structure holds pagination data
type Meta struct {
	Take      int    `json:"take"`
	Page      int    `json:"page"`
	TotalData int    `json:"total_data"`
	TotalPage int    `json:"total_page"`
	Sort      string `json:"sort"`
	SortBy    string `json:"sort_by"`
	Filter    string `json:"filter,omitempty"`
	FilterBy  string `json:"filter_by,omitempty"`
}

// New creates and initializes a Meta object with default pagination settings.
func New(c *fiber.Ctx) Meta {
	meta := Meta{
		Take:   10,
		Page:   0,
		Sort:   "asc",
		SortBy: "id",
	}

	meta.Page = ToInt(c.Query("page"))
	meta.Take = DefaultTake(ToInt(c.Query("take")))
	sort := c.Query("sort")
	sortby := c.Query("sort_by")
	filter := c.Query("filter")
	filterby := c.Query("filter_by")

	if sort != "" {
		meta.Sort = sort
	}

	if sortby != "" {
		meta.SortBy = sortby
	}

	if filter != "" {
		meta.Filter = filter
	}

	if filterby != "" {
		meta.FilterBy = filterby
	}

	return meta
}

// Count calculates the total number of pages based on the total data count.
func (m *Meta) Count(totaldata int) {
	m.TotalData = totaldata
	m.TotalPage = (totaldata + m.Take - 1) / m.Take
}

// GetSkipAndLimit calculates the offset (skip) and limit values for pagination.
func (m *Meta) GetSkipAndLimit() (int, int) {
	switch {
	case m.Page <= 0:
		m.Page = 1
		return 0, m.Take
	default:
		return ((m.Page - 1) * m.Take), m.Take
	}
}

// DefaultTake returns a default value for the take parameter if the provided value is less than or equal to 0.
func DefaultTake(i int) int {
	if i <= 0 {
		return 10
	}
	return i
}

// ToInt converts a string to an integer.
func ToInt(i string) int {
	res, err := strconv.Atoi(i)
	if err != nil {
		return 0
	}
	return res
}
