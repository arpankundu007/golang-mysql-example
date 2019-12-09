package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"mobile-specs-golang/constants"
)

func GetDB() *sql.DB {
	driver, dataSource := getDBConfig()
	dbSource, err := sql.Open(driver, dataSource)
	if err != nil {
		panic(err)
	}

	_, err = dbSource.Exec("CREATE DATABASE IF NOT EXISTS " + constants.DbName) //Create main if not exists
	if err != nil {
		panic(err)
	}

	defer dbSource.Close()

	db, err := sql.Open(driver, dataSource+constants.DbName) //Open the DB and return it
	if err != nil {
		panic(err)
	}
	return db
}

func getDBConfig() (string, string){
	return constants.DbDriver, constants.DbDataSource
}

func CreateTableIfNotExists(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `" + constants.TableName + "`(" +
		" `id` varchar (50) NOT NULL," +
		" `brand` varchar (30) NOT NULL, " +
		" `model` varchar (30) NOT NULL," +
		" `processor` varchar (30) NOT NULL, " +
		" `ram` varchar (30) NOT NULL," +
		" `storage` varchar (30) NOT NULL, " +
		" `createdAt` varchar (30) NOT NULL, " +
		" `updatedAt` varchar (30) NOT NULL, " +
		"  PRIMARY KEY (`id`))" +
		"  ENGINE=InnoDB " +
		"  DEFAULT CHARSET=latin1")
	if err != nil {
		panic(err.Error())
	}
}
