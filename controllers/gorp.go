// author roshan
package controllers

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var dbmap = initDb()

//database initialization
func initDb() *gorp.DbMap {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1234"
	dbName := "shardINDEX"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
