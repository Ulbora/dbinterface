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
	BeginTransaction() Transaction
	Test(query string, args ...interface{}) *DbRow
	Insert(query string, args ...interface{}) (bool, int64)
	Update(query string, args ...interface{}) bool
	Get(query string, args ...interface{}) *DbRow
	GetList(query string, args ...interface{}) *DbRows
	Delete(query string, args ...interface{}) bool
	Close() bool
}

//Transaction transaction
type Transaction interface {
	Insert(query string, args ...interface{}) (bool, int64)
	Update(query string, args ...interface{}) bool
	Delete(query string, args ...interface{}) bool
	Commit() bool
	Rollback() bool
}

//GO111MODULE=on go mod init github.com/Ulbora/dbinterface
