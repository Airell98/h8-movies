package database

import (
	"database/sql"
	"fmt"
	"h8-movies/infra/config"
	"log"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	appConfig := config.GetAppConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DBHost, appConfig.DBPort, appConfig.DBUser, appConfig.DBPassword, appConfig.DBName,
	)

	db, err = sql.Open(appConfig.DBDialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while trying to validate database arguments:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while trying to connect to database:", err)
	}

}

func handleCreateRequiredTables() {
	userTable := `
		CREATE TABLE IF NOT EXISTS "users" (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL,
			password TEXT NOT NULL,
			createdAt timestamptz DEFAULT now(),
			updatedAt timestamptz DEFAULT now()
		);
	`

	movieTable := `
		CREATE TABLE IF NOT EXISTS "movies" (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			imageUrl TEXT NOT NULL,
			price int NOT NULL,
			userId int NOT NULL,
			createdAt timestamptz DEFAULT now(),
			updatedAt timestamptz DEFAULT now(),
			CONSTRAINT movies_user_id_fk
				FOREIGN KEY(userId)
					REFERENCES users(id)
						ON DELETE SET NULL
		);
	`

	createTableQueries := fmt.Sprintf("%s %s", userTable, movieTable)

	_, err = db.Exec(createTableQueries)

	if err != nil {
		log.Panic("error occured while trying to create required tables:", err)
	}
}

func InitiliazeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTables()
}

func GetDatabaseInstance() *sql.DB {
	return db
}
