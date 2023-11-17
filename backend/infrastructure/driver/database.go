package driver

import (
	"database/sql"
	"fmt"
	"os"

	"home/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const TxKey = "transactionObject"
const ErrDuplicateEntryNumber = 1062

func NewDB(conf *config.Config) *gorm.DB {
	var dialector gorm.Dialector
	dbUser := conf.DB_USER
	dbPwd := conf.DB_PWD
	dbName := conf.DB_NAME
	dbPort := conf.DB_PORT
	if config.ExistEnvFile() {
		dbHostName := conf.DB_HOSTNAME
		if dbPwd != "" {
			dbUser += ":" + dbPwd
		}
		dsn := fmt.Sprintf(
			"%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbHostName, dbPort, dbName,
		)
		dialector = mysql.Open(dsn)
	} else {
		dbTCPHost := conf.DB_TCPHOST
		dbURI := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			dbUser, dbPwd, dbTCPHost, dbPort, dbName,
		)
		sqlDB, _ := sql.Open("mysql", dbURI)
		dialector = mysql.New(mysql.Config{Conn: sqlDB})
	}
	db, err := gorm.Open(
		dialector, &gorm.Config{},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return db
}
