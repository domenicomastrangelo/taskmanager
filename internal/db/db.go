package db

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDefault() *sql.DB {
	return Connect("", "")
}

func Connect(dbDirPath string, dbFileName string) *sql.DB {
	var dbFilePath string

	if strings.Compare(dbDirPath, "") == 0 || strings.Compare(dbFileName, "") == 0 {
		homeDir, err := os.UserHomeDir()

		if err != nil {
			log.Println(err.Error())
			return nil
		}

		dbFilePath = homeDir + "/.taskmanager/db.sqlite"
	} else {
		dbFilePath = dbDirPath + dbFileName
	}

	if DB == nil {
		var err error

		if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
			os.MkdirAll(dbDirPath, 0700)
			file, err := os.Create(dbFilePath)

			if err != nil {
				log.Println("Cannot create db.sqlite file")
				log.Println(err.Error())
				return nil
			}
			file.Close()
		}

		DB, err = sql.Open("sqlite3", dbFilePath)

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
