package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() *sql.DB {
	if DB == nil {
		var err error

		homeDir, err := os.UserHomeDir()

		if err != nil {
			log.Println(err.Error())
			return nil
		}

		if _, err := os.Stat(homeDir + "/.taskmanager/db.sqlite"); os.IsNotExist(err) {
			os.MkdirAll(homeDir+"/.taskmanager", 0700)
			file, err := os.Create(homeDir + "/.taskmanager/db.sqlite")

			if err != nil {
				log.Println("Cannot create db.sqlite file")
				log.Println(err.Error())
				return nil
			}
			file.Close()
		}

		DB, err = sql.Open("sqlite3", homeDir+"/.taskmanager/db.sqlite")

		createDBIfNotExists()

		if err != nil {
			log.Println(err.Error())
			return nil
		}
	}

	return DB
}

func createDBIfNotExists() {
	tasks_table := `CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "title" TEXT,
        "message" TEXT,
        "done" BOOL,
        "created_at" DATETIME);`

	query, err := DB.Prepare(tasks_table)
	if err != nil {
		log.Println(err)
		return
	}
	query.Exec()
}
