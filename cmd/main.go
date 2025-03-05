package main

import (
	"database/sql"
	"os"

	"github.com/Modalessi/iau_resources/database"
	fanar "github.com/Modalessi/iau_resources/fanar_api"
	"github.com/Modalessi/iau_resources/storage"
	"github.com/Modalessi/iau_resources/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load()

	DB_URL := os.Getenv("DB_URL")
	utils.Assert(DB_URL != "", "could not read DB_URL from .env file, please check that it exist")

	JWT_SECRET := os.Getenv("JWT_SECRET")
	utils.Assert(JWT_SECRET != "", "could not read JWT_SECRET from .env file, please check that it exist")

	ADMIN_EMAIL := os.Getenv("ADMIN_EMAIL")
	utils.Assert(ADMIN_EMAIL != "", "could not read ADMIN_EMAIL from .env file, please check that it exist")

	db, err := sql.Open("postgres", DB_URL)
	utils.ErrorAssert(err, "error connecting to database")

	quries := database.New(db)

	fanarStorage := storage.NewStorage(db, quries)
	fanarServer := fanar.NewFanarServer(":4000", JWT_SECRET, ADMIN_EMAIL, fanarStorage)

	err = fanarServer.Start()
	if err != nil {
		panic("error starting fanar server")
	}
}
