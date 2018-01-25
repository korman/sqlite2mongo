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
		rows, err := db.Query("SELECT * FROM cdsgus where CtfId LIKE '370102199%'")
		checkErr(err)

		fmt.Println("hehehe")

		for rows.Next() {
			var name string = ""
			var cardNo string = ""
			var descriot string = ""
			var ctfTp string = ""
			var ctfId string = ""
			var gender string = ""
			var birthday string = ""
			var address string = ""
			var zip string = ""
			var dirty string = ""
			var district1 string = ""
			var district2 string = ""
			var district3 string = ""
			var district4 string = ""
			var district5 string = ""
			var district6 string = ""
			var firstNm string = ""
			var lastNm string = ""
			var duty string = ""
			var mobile string = ""
			var tel string = ""
			var fax string = ""
			var email string = ""
			var nation string = ""
			var taste string = ""
			var education string = ""
			var company string = ""
			var cTel string = ""
			var cAddress string = ""
			var cZip string = ""
			var family string = ""
			var version string = ""
			var id string = ""

			err = rows.Scan(&name, &cardNo, &descriot, &ctfTp,&ctfId,
				&gender,&birthday,&address,&zip,&dirty,&district1,
					&district2,&district3,&district4,&district5,&district6,&firstNm,
						&lastNm,&duty,&mobile,&tel,&fax,&email,
							&nation,&taste,&education,&company,&cTel,&cAddress,
								&cZip,&family,&version,&id)

			if err != nil {
				fmt.Println(err)
			}

			checkErr(err)
			fmt.Println(name + " " + ctfId + "  " + address + "  " + gender + " " + birthday)
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