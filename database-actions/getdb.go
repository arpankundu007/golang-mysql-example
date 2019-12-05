package database_actions

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"mobile-specs-golang/constants"
)

func GetDB(DbName string) *sql.DB {
	dbSource, err := sql.Open("mysql", "root:1q2w3e4r@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}

	_, err = dbSource.Exec("CREATE DATABASE IF NOT EXISTS " + DbName) //Create main if not exists
	if err != nil {
		panic(err)
	}

	defer dbSource.Close()

	db, err := sql.Open("mysql", "root:1q2w3e4r@tcp(127.0.0.1:3306)/"+DbName) //Open the main and return it
	if err != nil {
		panic(err)
	}
	return db
}

func CreateTableIfNotExists(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS ` " + constants.TableName + "`(" +
		" `id` varchar (50) NOT NULL," +
		" `brand` varchar (30) NOT NULL, " +
		" `model` varchar (30) NOT NULL," +
		" `processor` varchar (30) NOT NULL, " +
		" `ram` varchar (30) NOT NULL," +
		" `storage` varchar (30) NOT NULL, " +
		"  PRIMARY KEY (`id`))" +
		"  ENGINE=InnoDB " +
		"  DEFAULT CHARSET=latin1")
	if err != nil {
		panic(err.Error())
	}
}
