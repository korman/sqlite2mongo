package user_data

import "gopkg.in/mgo.v2/bson"

type CdsgusData struct {
	Id_ bson.ObjectId `bson:"_id"`
	Base BaseInfo `bson:"baseinfo"`
	Contact ContactInfo `bson:"contact"`
	RegTime string `bson:"regTime"`
	OldId string `bson:"oldId"`
}

type BaseInfo struct {
	Name string `bson:"name"`
	CtfTp string `bson:"ctfTp"`
	CtfId string `bson:"ctfId"`
	Gender string `bson:"gender"`
	Birthday string `bson:"birthday"`
	Education string `bson:"education"`
	Nation string `bson:"nation"`
	Nationality string `bson:"nationality"`
	Dirty string `bson:"dirty"`
	Duty string `bson:"duty"`
	Family string `bson:"family"`
}

type ContactInfo struct {
	Address string `bson:"address"`
	Zip string `bson:"zip"`
	Mobile string `bson:"mobile"`
	Tel string `bson:"tel"`
	Fax string `bson:"fax"`
	Email string `bson:"email"`
	Company string `bson:"company"`
}