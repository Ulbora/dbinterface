package mysql

import (
	"fmt"
	di "github.com/Ulbora/dbinterface"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var dbi di.Database

var dbi2 di.Database

var iid1 int64

func TestMySQLDb_Connectfail(t *testing.T) {
	var mdb MyDB
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin1"
	mdb.Database = "testdb1"
	dbi = &mdb
	suc := dbi.Connect()
	if suc {
		t.Fail()
	}

}

func TestMySQLDb_Connect(t *testing.T) {
	var mdb MyDB
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "testdb"
	dbi = &mdb
	suc := dbi.Connect()
	if !suc {
		t.Fail()
	}
}

func TestMySQLDb_Connect2(t *testing.T) {
	var mdb MyDB
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "ffl_list_10012018"
	dbi2 = &mdb
	suc := dbi2.Connect()
	if !suc {
		t.Fail()
	}
}

func TestMySQLDb_Test(t *testing.T) {
	var rtn bool
	var q = "select count(*) from test "
	var a []interface{}
	rowPtr := dbi.Test(q, a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		fmt.Print("Records found during test ")
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

func TestMySQLDb_Insertfail(t *testing.T) {
	var q = "insert into test2 (name, address) values(?, ?)"
	var a []interface{}
	a = append(a, "test insert 1", "123 main st")
	suc, id := dbi.Insert(q, a...)
	if suc {
		t.Fail()
	} else {
		fmt.Println("inserted id: ", id)
	}
}

func TestMySQLDb_Insert(t *testing.T) {
	var q = "insert into test (name, address) values(?, ?)"
	var a []interface{}
	a = append(a, "test insert 1", "123 main st")
	suc, id := dbi.Insert(q, a...)
	if !suc || id < 1 {
		t.Fail()
	} else {
		iid1 = id
		fmt.Println("inserted id: ", id)
	}
}

func TestMySQLDb_Updatefail(t *testing.T) {
	var q = "update test1 set name = ? , address = ? where id = ? "
	var a []interface{}
	a = append(a, "test insert 2", "123456 main st", iid1)
	suc := dbi.Update(q, a...)
	if suc {
		t.Fail()
	}
}

func TestMySQLDb_Update(t *testing.T) {
	var q = "update test set name = ? , address = ? where id = ? "
	var a []interface{}
	a = append(a, "test insert 2", "123456 main st", iid1)
	suc := dbi.Update(q, a...)
	if !suc {
		t.Fail()
	}
}

func TestMySQLDb_Get(t *testing.T) {
	var rtn bool
	var q = "select * from test where id = ? "
	var a []interface{}
	a = append(a, iid1)
	rowPtr := dbi.Get(q, a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get ")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if iid1 != int64Val {
			fmt.Print(iid1)
			fmt.Print(" != ")
			fmt.Println(int64Val)
			t.Fail()
		} else {
			fmt.Print("found id")
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

func TestMySQLDb_GetList(t *testing.T) {
	var suc bool
	var q = "select * from test where address = ? "
	var a []interface{}
	a = append(a, "123456 main st")
	rowsPtr := dbi.GetList(q, a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Println("rows found: ", foundRows)
		//fmt.Println("GetList results: --------------------------")
		fmt.Println("rows found count: ", len(foundRows))
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

func TestMySQLDb_Deletefail(t *testing.T) {
	var q = "delete from test1 where id = ? "
	var a []interface{}
	a = append(a, iid1)
	suc := dbi.Delete(q, a...)
	if suc {
		t.Fail()
	}
}

func TestMySQLDb_Delete(t *testing.T) {
	var q = "delete from test where id = ? "
	var a []interface{}
	a = append(a, iid1)
	suc := dbi.Delete(q, a...)
	if !suc {
		t.Fail()
	}
}

func TestMySQLDb_Close(t *testing.T) {
	suc := dbi.Close()
	if !suc {
		t.Fail()
	}
}
