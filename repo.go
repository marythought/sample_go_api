package main

import "fmt"

var currentId int

var filters Filters

// Give us some seed data
func init() {
	RepoCreateFilter(Filter{Name: "Write presentation"})
	RepoCreateFilter(Filter{Name: "Host meetup"})
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
	filters = append(filters, f)
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
