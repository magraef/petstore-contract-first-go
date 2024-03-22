package internal

type CategoryId int64

// Category defines model for Category.
type Category struct {
	Id   *CategoryId
	Name string
}
