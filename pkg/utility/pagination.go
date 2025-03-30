package utility

import (
	"Go-Starter-Template/internal/utils/pagination"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"reflect"
	"strings"
)

var (
	ErrSortBy           = errors.New("invalid sort (must be 'asc' or 'desc')")
	ErrInvalidTypeModel = errors.New("invalid type model")
	ErrInvalidField     = errors.New("invalid filter or sort field")
)

type MetaService struct {
	Filter map[string]string
	Sorter map[string]string
}

type Option func(*MetaService)

// WithFilters now accepts a *fiber.Ctx and MetaService options.
func WithFilters(db *gorm.DB, m *pagination.Meta, opts ...Option) *gorm.DB {
	metaService := MetaService{
		Filter: make(map[string]string),
		Sorter: make(map[string]string),
	}

	for _, opt := range opts {
		opt(&metaService)
	}

	return metaService.buildFilter(db, m)
}

// AddModels adds filter and sorter mappings for struct fields.
func AddModels(model any, tablePrefix string) Option {
	return func(ms *MetaService) {
		v := reflect.TypeOf(model)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		if v.Kind() == reflect.Struct {
			for i := 0; i < v.NumField(); i++ {
				field := v.Field(i)
				jsonTag := field.Tag.Get("json")
				if jsonTag == "" {
					jsonTag = field.Name
				}

				fullField := fmt.Sprintf("%s.%s", tablePrefix, jsonTag)

				if field.Anonymous {
					addEmbeddedFields(field.Type, ms, tablePrefix)
				} else {
					switch field.Type.Kind() {
					case reflect.String:
						ms.Filter[jsonTag] = fmt.Sprintf("%s ILIKE ?", fullField)
						ms.Sorter[jsonTag] = fullField
					default:
						ms.Filter[jsonTag] = fmt.Sprintf("%s = ?", fullField)
						ms.Sorter[jsonTag] = fullField
					}
				}
			}
		}
	}
}

// AddCustomField allows custom filters and sorters to be added.
func AddCustomField(field string, filterquery string, alias ...string) Option {
	return func(ms *MetaService) {
		ms.Filter[field] = filterquery
		ms.Sorter[field] = field
		if len(alias) > 0 {
			ms.Sorter[field] = alias[0]
		}
	}
}

// buildFilter applies filters and sorting based on pagination metadata.
func (ms *MetaService) buildFilter(db *gorm.DB, meta *pagination.Meta) *gorm.DB {
	query := db

	// Handle filters
	filterBy := strings.Split(meta.FilterBy, ",")
	filters := strings.Split(meta.Filter, ",")

	for i, field := range filterBy {
		if i >= len(filters) {
			break
		} else if filters[i] == "" {
			continue
		}

		if condition, ok := ms.Filter[field]; ok {
			var filterValue string
			if i < len(filters) {
				filterValue = filters[i]
			}

			// Handle string-based filtering
			if strings.Contains(strings.ToUpper(condition), "ILIKE") {
				if filterValue != "" {
					filterValue = fmt.Sprintf("%%%s%%", strings.ToLower(filterValue))
				} else {
					filterValue = "%%"
				}
			}

			questionMarks := strings.Count(condition, "?")

			filterValues := make([]interface{}, questionMarks)
			for j := range filterValues {
				filterValues[j] = filterValue
			}

			query = query.Where(condition, filterValues...)
			if err := query.Error; err != nil {
				return query
			}
		} else if field != "" {
			query.Error = ErrInvalidField
			return query
		}
	}

	// Handle sorting
	if meta.SortBy != "" {
		if _, ok := ms.Sorter[meta.SortBy]; !ok {
			query.Error = ErrInvalidTypeModel
			return query
		}

		if meta.Sort != "asc" && meta.Sort != "desc" {
			query.Error = ErrSortBy
			return query
		}

		order := fmt.Sprintf("%s %s", ms.Sorter[meta.SortBy], meta.Sort)
		query = query.Order(order)
		if err := query.Error; err != nil {
			return query
		}
	}

	// Count total records for pagination
	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		log.Printf("error counting total count: %v", err)
		return query
	}

	meta.Count(int(totalCount))

	// Apply pagination
	skip, limit := meta.GetSkipAndLimit()
	query = query.Scopes(paginate(skip, limit))
	if err := query.Error; err != nil {
		log.Printf("error paging query: %v", err)
		return query
	}

	return query
}

// paginate applies the skip and limit for pagination.
func paginate(page, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(page).Limit(perPage)
	}
}

// addEmbeddedFields handles embedded structs in models.
func addEmbeddedFields(embedType reflect.Type, ms *MetaService, tablePrefix string) {
	if embedType.Kind() == reflect.Ptr {
		embedType = embedType.Elem()
	}
	if embedType.Kind() == reflect.Struct {
		for i := 0; i < embedType.NumField(); i++ {
			field := embedType.Field(i)
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				jsonTag = field.Name
			}

			fullField := fmt.Sprintf("%s.%s", tablePrefix, jsonTag)

			switch field.Type.Kind() {
			case reflect.String:
				ms.Filter[jsonTag] = fmt.Sprintf("%s ILIKE ?", fullField)
				ms.Sorter[jsonTag] = fullField
			default:
				ms.Filter[jsonTag] = fmt.Sprintf("%s = ?", fullField)
				ms.Sorter[jsonTag] = fullField
			}
		}
	}
}
