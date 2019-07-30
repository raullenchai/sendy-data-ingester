package mysql

import (
	"database/sql"
	//"log"
)

var Db *sql.DB = new(sql.DB)

func Open() bool {
	//var err error
	//Db, err = sql.Open("mysql", keys.MysqlUid+":"+keys.MysqlPwd+"@tcp("+keys.MysqlHost+":"+keys.MysqlPort+")/"+keys.MysqlDbName+"?charset=utf8")
	//if err != nil {
	//	log.Println("mysql", "open err", err)
	//	return false
	//}
	//
	//Db.SetMaxOpenConns(100)
	//Db.SetMaxIdleConns(100)
	//
	//err = Db.Ping()
	//if err != nil {
	//	log.Println("mysql", "open and ping err", err)
	//	return false
	//}
	//
	//return true
}
