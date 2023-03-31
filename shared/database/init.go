package database

import (
	"nwd/shared/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn Connection = Connection{db: &gorm.DB{}}
var connTrx Connection = Connection{db: &gorm.DB{}}

func Init(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect mysql.")
	}

	conn.db = db

	log.GetLog().Info("", "Mysql DB connected")
}

type IConnection interface {
	CommitTransaction()
	RollbackTransaction()
	RollbackOnException()
	Tables(tableName string) *gorm.DB
}

type Connection struct {
	db *gorm.DB
}

func (c Connection) Tables(tableName string) *gorm.DB {
	return c.db.Table(tableName)
}

func (c Connection) RollbackTransaction() {
	c.db.Rollback()
}

func (c Connection) RollbackOnException() {
	if err := recover(); err != nil {
		c.db.Rollback()
		panic(err)
	}
}

func (c Connection) CommitTransaction() {
	c.db.Commit()
}

func Stop() {
	db, _ := conn.db.DB()
	db.Close()
}
