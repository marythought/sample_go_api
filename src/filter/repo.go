package filter

import (
	"fmt"
	"time"
)

var currentId int

var filters Filters

// Give us some seed data
func init() {
	t := time.Now()
	RepoCreateFilter(Filter{Name: "Filter #1", OrgId: 2, CreatedTime: t, ModifiedTime: t})
	RepoCreateFilter(Filter{Name: "Filter #2", OrgId: 2, CreatedTime: t, ModifiedTime: t})
}

func RepoFindFilter(id int) Filter {
	for _, f := range filters {
		if f.FilterId == id {
			return f
		}
	}
	// return empty Filter if not found
	return Filter{}
}

func RepoCreateFilter(f Filter) Filter {
	currentId += 1
	f.FilterId = currentId
	f.CreatedTime = time.Now()
	f.ModifiedTime = time.Now()
	filters = append(filters, f)
	return f
}

func RepoUpdateFilter(f Filter) Filter {
	return f
}

func RepoDestroyFilter(id int) error {
	for i, f := range filters {
		if f.FilterId == id {
			filters = append(filters[:i], filters[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Filter with id of %d to delete", id)
}
