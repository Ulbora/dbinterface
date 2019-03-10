package mysql

import (
	"database/sql"
	di "github.com/Ulbora/dbinterface"
	//"log"

	_ "github.com/go-sql-driver/mysql"
)

//MyDBMock MyDBMock
type MyDBMock struct {
	Host        string
	User        string
	Password    string
	Database    string
	db          *sql.DB
	err         error
	MockSuccess bool
	MockID      int64
	MockRow     *di.DbRow
	MockRows    *di.DbRows
}

//Connect Connect
func (m *MyDBMock) Connect() bool {
	return m.MockSuccess
}

//Test Test
func (m *MyDBMock) Test(query string, args ...interface{}) *di.DbRow {
	return m.MockRow
}

//Insert Insert
func (m *MyDBMock) Insert(query string, args ...interface{}) (bool, int64) {
	return m.MockSuccess, m.MockID
}

//Update Update
func (m *MyDBMock) Update(query string, args ...interface{}) bool {
	return m.MockSuccess
}

//Get Get
func (m *MyDBMock) Get(query string, args ...interface{}) *di.DbRow {
	return m.MockRow
}

//GetList GetList
func (m *MyDBMock) GetList(query string, args ...interface{}) *di.DbRows {
	return m.MockRows
}

//Delete Delete
func (m *MyDBMock) Delete(query string, args ...interface{}) bool {
	return m.MockSuccess
}

//Close Close
func (m *MyDBMock) Close() bool {
	return m.MockSuccess
}
