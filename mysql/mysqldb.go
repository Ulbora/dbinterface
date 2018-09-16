package mysqldb

import (
	"database/sql"
	di "dbinterface"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

//MySQLDb MySQLDb
type MySQLDb struct {
	Host     string
	User     string
	Password string
	Database string
}

//Connect Connect
func (m *MySQLDb) Connect() bool {
	var rtn = false
	var conStr = m.User + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.Database
	db, err = sql.Open("mysql", conStr)
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		err = db.Ping()
		if err != nil {
			log.Println("Error:", err.Error())
		} else {
			rtn = true
		}
	}
	return rtn
}

//Test Test
func (m *MySQLDb) Test(query string, args ...interface{}) *di.DbRow {
	return m.Get(query, args...)
}

//Insert Insert
func (m *MySQLDb) Insert(query string, args ...interface{}) (bool, int64) {
	var success = false
	var id int64 = -1
	var stmtIns *sql.Stmt
	stmtIns, err = db.Prepare(query)
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		defer stmtIns.Close()
		res, err := stmtIns.Exec(args...)
		if err != nil {
			log.Println("Insert Exec err:", err.Error())
		} else {
			id, err = res.LastInsertId()
			if err != nil {
				log.Println("Error:", err.Error())
			} else {
				success = true
			}
		}
	}
	return success, id
}

//Update Update
func (m *MySQLDb) Update(query string, args ...interface{}) bool {
	var success = false
	var stmtUp *sql.Stmt
	stmtUp, err = db.Prepare(query)
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		defer stmtUp.Close()
		res, err := stmtUp.Exec(args...)
		if err != nil {
			log.Println("Update Exec err:", err.Error())
		} else {
			log.Println("Update Exec success:")
			affectedRows, err := res.RowsAffected()
			if err != nil && affectedRows == 0 {
				log.Println("Error:", err.Error())
			} else {
				success = true
			}
		}
	}
	return success
}

//Get Get
func (m *MySQLDb) Get(query string, args ...interface{}) *di.DbRow {
	var rtn di.DbRow
	stmtGet, err := db.Prepare(query)
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		defer stmtGet.Close()
		rows, err := stmtGet.Query(args...)
		defer rows.Close()
		if err != nil {
			log.Println("Get err: ", err)
		} else {
			columns, err := rows.Columns()
			if err != nil {
				log.Println("Error:", err.Error())
			} else {
				rtn.Columns = columns
				rowValues := make([]sql.RawBytes, len(columns))
				scanArgs := make([]interface{}, len(rowValues))
				for i := range rowValues {
					scanArgs[i] = &rowValues[i]
				}
				for rows.Next() {
					err = rows.Scan(scanArgs...)
					if err != nil {
						log.Println("Error:", err.Error())
					}
					for _, col := range rowValues {
						var value string
						if col == nil {
							value = "NULL"
						} else {
							value = string(col)
						}
						rtn.Row = append(rtn.Row, value)
					}
				}
				if err = rows.Err(); err != nil {
					log.Println("Error:", err.Error())
				}
			}
		}
	}
	return &rtn
}

//GetList GetList
func (m *MySQLDb) GetList(query string, args ...interface{}) *di.DbRows {
	var rtn di.DbRows
	stmtGet, err := db.Prepare(query)
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		defer stmtGet.Close()
		rows, err := stmtGet.Query(args...)
		defer rows.Close()
		if err != nil {
			log.Println("GetList err: ", err)
		} else {
			columns, err := rows.Columns()
			if err != nil {
				log.Println("Error:", err.Error())
			}
			rtn.Columns = columns
			rowValues := make([]sql.RawBytes, len(columns))
			scanArgs := make([]interface{}, len(rowValues))
			for i := range rowValues {
				scanArgs[i] = &rowValues[i]
			}
			for rows.Next() {
				var rowValuesStr []string
				err = rows.Scan(scanArgs...)
				if err != nil {
					log.Println("Error:", err.Error())
				}
				for _, col := range rowValues {
					var value string
					if col == nil {
						value = "NULL"
					} else {
						value = string(col)
					}
					rowValuesStr = append(rowValuesStr, value)
				}
				rtn.Rows = append(rtn.Rows, rowValuesStr)
			}
			if err = rows.Err(); err != nil {
				log.Println("Error:", err.Error())
			}
		}
	}
	return &rtn
}

//Delete Delete
func (m *MySQLDb) Delete(query string, args ...interface{}) bool {
	var success = false
	var stmt *sql.Stmt
	stmt, err = db.Prepare(query)
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		defer stmt.Close()
		res, err := stmt.Exec(args...)
		if err != nil {
			log.Println("Delete Exec err:", err.Error())
		} else {
			affectedRows, err := res.RowsAffected()
			if err != nil {
				log.Println("Error:", err.Error())
			} else {
				//fmt.Println("Delete Exec success:")
				if affectedRows > 0 {
					success = true
				}
			}
		}
	}
	return success
}

//Close Close
func (m *MySQLDb) Close() bool {
	var rtn = false
	err := db.Close()
	if err != nil {
		log.Println("database close error: ", err)
	} else {
		rtn = true
	}
	return rtn
}
