package main

import (
	"database/sql"
	"fmt"
	"github.com/korman/sqlite2mongo/user_data"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var dbData = make(chan *user_data.CdsgusData, 40)

func main() {
	for i := 0; i < 20; i++ {
		go writeToDb()
	}

	readFromDb()
}

func readFromDb() {
	db, err := sql.Open("sqlite3", "./2000w.db")
	var count int = 0

	if nil != err {
		fmt.Printf(err.Error())
	} else {
		var cardNo string = ""
		var descriot string = ""
		var district1 string = ""
		var district3 string = ""
		var district4 string = ""
		var district5 string = ""
		var district6 string = ""
		var firstNm string = ""
		var lastNm string = ""
		var taste string = ""
		var cTel string = ""
		var cAddress string = ""
		var cZip string = ""

		rows, err := db.Query("SELECT * FROM cdsgus")

		startTime := time.Now().Unix()

		for rows.Next() {
			data := &user_data.CdsgusData{}

			err = rows.Scan(&data.Base.Name, &cardNo, &descriot, &data.Base.CtfTp, &data.Base.CtfId,
				&data.Base.Gender, &data.Base.Birthday, &data.Contact.Address, &data.Contact.Zip, &data.Base.Dirty,
				&district1, &data.Base.Nationality, &district3, &district4, &district5, &district6, &firstNm, &lastNm,
				&data.Base.Duty, &data.Contact.Mobile, &data.Contact.Tel, &data.Contact.Fax, &data.Contact.Email,
				&data.Base.Nation, &taste, &data.Base.Education, &data.Contact.Company, &cTel, &cAddress, &cZip,
				&data.Base.Family, &data.RegTime, &data.OldId)

			data.Id_ = bson.NewObjectId()

			if err != nil {
				fmt.Println(err)
			}

			if count%1000 == 0 {
				fmt.Printf("当前插入数量:%d\n，已过去%d秒", count, time.Now().Unix()-startTime)
			}

			dbData <- data

			count++
		}
	}

	db.Close()
}

func writeToDb() {
	session, err := mgo.Dial("192.168.1.145")
	defer session.Close()

	if nil != err {
		panic(err)
	}

	for {
		select {
		case data := <-dbData:
			c := session.DB("people_2000w").C("cdsgus")

			err = c.Insert(&data)

			if nil != err {
				panic(err)
			}
		}
	}
}
