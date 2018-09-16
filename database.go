package dbinterface

//DbRow database row
type DbRow struct {
	Columns []string
	Row     []string
}

//DbRows array of database rows
type DbRows struct {
	Columns []string
	Rows    [][]string
}

//Database Database
type Database interface {
	Connect() bool
	Test(query string, args ...interface{}) *DbRow
	Insert(query string, args ...interface{}) (bool, int64)
	Update(query string, args ...interface{}) bool
	Get(query string, args ...interface{}) *DbRow
	GetList(query string, args ...interface{}) *DbRows
	Delete(query string, args ...interface{}) bool
	Close() bool
}
