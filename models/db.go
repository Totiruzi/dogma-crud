package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

var ormObject orm.Ormer

// ConnectToDb initializes the ORM and connection to the postgres Db
func ConnectToDb() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=dbUser password=yourPassword dbname=dbName host=dbHost sslmode=disable")
	orm.RegisterModel(new(Users))
	ormObject = orm.NewOrm()
}

// GetOrrObject getter function for the ORM object with which we an query the database
func GetOrmObject() orm.Ormer {
	return ormObject
}
