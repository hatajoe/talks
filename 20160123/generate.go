//go:generate gen -f
//go:generate margen
package models

// +gen * slice:"All,Any,Count,DistinctBy,First,GroupBy[int],GroupBy[string],Shuffle,SortBy,Where"
// +margen
type Member struct {
	ID int64
}
