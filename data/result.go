package data

type Result struct {
	Connection string
	Query      string
	Columns    []string
	Rows       []Row
}
type Row struct {
	Values []string
}
