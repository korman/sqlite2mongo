package main

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./2000w.db")

	if nil != err {
		fmt.Printf(err.Error())
	} else {
		rows, err := db.Query("SELECT * FROM cdsgus where Name = '姜川'")
		checkErr(err)

		for rows.Next() {
			var name string = ""
			var cardNo string = ""
			var descriot string = ""
			err = rows.Scan(&name, &cardNo, &descriot)
			checkErr(err)
			fmt.Println(name)
			fmt.Println(cardNo)
			fmt.Println(descriot)
		}
	}

	//插入数据
	//stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	//checkErr(err)
	//
	//res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	//checkErr(err)
	//
	//id, err := res.LastInsertId()
	//checkErr(err)
	//
	//fmt.Println(id)
	////更新数据
	//stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	//checkErr(err)
	//
	//res, err = stmt.Exec("astaxieupdate", id)
	//checkErr(err)
	//
	//affect, err := res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println(affect)

	//查询数据

	//删除数据
	//stmt, err = db.Prepare("delete from userinfo where uid=?")
	//checkErr(err)
	//
	//res, err = stmt.Exec(id)
	//checkErr(err)
	//
	//affect, err = res.RowsAffected()
	//checkErr(err)

	db.Close()
}

func checkErr(err error) {

}