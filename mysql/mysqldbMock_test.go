package mysql

import (
	"fmt"
	"strconv"
	//"database/sql"
	di "github.com/Ulbora/dbinterface"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var dbim di.Database

func TestMyDBMock_Connect(t *testing.T) {
	var mdb MyDBMock
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin1"
	mdb.Database = "testdb1"
	mdb.MockSuccess = true
	dbim = &mdb
	suc := dbim.Connect()
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_Test(t *testing.T) {
	var mdb MyDBMock
	var rtn bool
	var rtnRow di.DbRow
	rtnRow.Row = []string{"1", "test2"}
	mdb.MockRow = &rtnRow
	dbim = &mdb
	var q = "select count(*) from test "
	var a []interface{}
	rowPtr := dbim.Test(q, a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		fmt.Print("Mock Records found during test ")
		fmt.Println(int64Val)
		if err2 != nil {
			fmt.Print(err2)
		}
		if int64Val >= 0 {
			rtn = true
		}
	}
	if !rtn {
		t.Fail()
	}
}

func TestMyDBMock_Insert(t *testing.T) {
	var mdb MyDBMock
	mdb.MockSuccess = true
	mdb.MockID = 1
	dbim = &mdb
	var q = "insert into test (name, address) values(?, ?)"
	var a []interface{}
	a = append(a, "test insert 1", "123 main st")
	suc, id := dbim.Insert(q, a...)
	if !suc || id < 1 {
		t.Fail()
	} else {
		iid1 = id
		fmt.Println("mock inserted id: ", id)
	}
}

func TestMyDBMock_Update(t *testing.T) {
	var mdb MyDBMock
	mdb.MockSuccess = true
	dbim = &mdb
	var q = "update test set name = ? , address = ? where id = ? "
	var a []interface{}
	a = append(a, "test insert 2", "123456 main st", iid1)
	suc := dbim.Update(q, a...)
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_Get(t *testing.T) {
	var mdb MyDBMock
	var rtnRow di.DbRow
	rtnRow.Row = []string{"2", "test2"}
	mdb.MockRow = &rtnRow
	dbim = &mdb
	var rtn bool
	var inint int64 = 2
	var q = "select * from test where id = ? "
	var a []interface{}
	a = append(a, inint)
	rowPtr := dbim.Get(q, a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get ")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if inint != int64Val {
			fmt.Print(" Mock Get ")
			fmt.Print(inint)
			fmt.Print(" != ")
			fmt.Println(int64Val)
			t.Fail()
		} else {
			fmt.Print("Mock found id")
			fmt.Print(" = ")
			fmt.Println(int64Val)
			rtn = true
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
	if !rtn {
		t.Fail()
	}
}

func TestMyDBMock_GetList(t *testing.T) {
	var mdb MyDBMock
	var rtnRows di.DbRows
	var r1 = []string{"1", "test1"}
	var r2 = []string{"2", "test2"}
	var val [][]string
	val = append(val, r1)
	val = append(val, r2)
	rtnRows.Rows = val
	mdb.MockRows = &rtnRows
	dbim = &mdb
	var suc bool
	var q = "select * from test where address = ? "
	var a []interface{}
	a = append(a, "123456 main st")
	rowsPtr := dbim.GetList(q, a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Println("Mock rows found: ", foundRows)
		//fmt.Println("GetList results: --------------------------")
		fmt.Println("Mock rows found count: ", len(foundRows))
		if len(foundRows) > 0 {
			suc = true
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_Delete(t *testing.T) {
	var mdb MyDBMock
	mdb.MockSuccess = true
	dbim = &mdb
	var inint int64 = 2
	var q = "delete from test1 where id = ? "
	var a []interface{}
	a = append(a, inint)
	suc := dbim.Delete(q, a...)
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_Close(t *testing.T) {
	var mdb MyDBMock
	mdb.MockSuccess = true
	dbim = &mdb
	suc := dbim.Close()
	if !suc {
		t.Fail()
	}
}
