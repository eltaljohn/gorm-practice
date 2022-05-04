package storage

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Driver of storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New create the connection with db
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

func newPostgresDB() {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			"postgres",
			"password",
			"127.0.0.1",
			"5432",
			"godb",
		)
		var err error
		db, err = gorm.Open(postgres.Open(dsn))
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("connected to postgres")
	})
}

func newMySQLDB() {
	once.Do(func() {
		var err error
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true&allowNativePasswords=true&parseTime=true",
			"root",
			"password",
			"127.0.0.1",
			"3306",
			"godb",
		)
		db, err = gorm.Open(mysql.Open(dsn))
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("connected to mySQL")
	})
}

// DB returns a unique instance of db
func DB() *gorm.DB {
	return db
}
