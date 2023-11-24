package config

import (
	"fmt"
	"os"
	"pegadaianempat/model"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")

	db  *gorm.DB
	err error
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var (
		host     = os.Getenv("PGHOST")
		port     = os.Getenv("PGPORT")
		user     = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
		dbname   = os.Getenv("PGDATABASE")
	)

	psqlInfo := fmt.Sprintf(`
	host=%s 
	port=%s
	user=%s `+`
	password=%s 
	dbname=%s 
	sslmode=disable`,
		host, port, user, password, dbname)

	//_ = psqlInfo

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
