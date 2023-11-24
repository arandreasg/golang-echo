package config

import (
	"fmt"
	"pegadaianempat/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "db-go-sql"
)

var (
	db  *gorm.DB
	err error
)

func Connect() {
	psqlInfo := fmt.Sprintf(`
	host=%s 
	port=%d 
	user=%s 	
	password=%s 
	dbname=%s 
	sslmode=disable`,
		host, port, user, password, dbname)

	// db, err := sql.Open("postgres", psqlInfo) // native way
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses konek tu DB\n", psqlInfo)

	db.AutoMigrate(&model.Employee{})
}

func GetDB() *gorm.DB {
	return db
}
