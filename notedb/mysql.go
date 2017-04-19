package notedb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	database   string
	password   string
	user       string
	appVersion string
}

func NewMySQLDB(config map[string]string, appVersion string) MySQLDB {
	if _, exists := config["database"]; !exists {
		config["database"] = "mysql"
	}

	if _, exists := config["database"]; !exists {
		config["password"] = "root"
	}

	if _, exists := config["database"]; !exists {
		config["user"] = "root"
	}

	return MySQLDB{
		database:   config["database"],
		password:   config["password"],
		user:       config["user"],
		appVersion: appVersion,
	}
}

const (
	mysqlDatabase   = "note"
	mysqlTable      = "NoteTable"
	mysqlNoteColumn = "Note"
)

var _ NoteDB = MySQLDB{}

func (mysqlDB MySQLDB) createMySQLClient() (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", mysqlDB.user, mysqlDB.password, mysqlDB.database, mysqlDatabase))
}

//NOTE initalize -> check if table exists elxe create table WHERE

// SELECT table_name FROM information_schema.tables where table_schema='mysqlDB';
// if table in rows okay else create new table

func (mysqlDB MySQLDB) GetAllNotes() ([]string, error) {
	db, err := mysqlDB.createMySQLClient()
	if err != nil {
		return []string{}, err
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT %s FROM %s", mysqlNoteColumn, mysqlTable))
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()
	notes := []string{}

	for rows.Next() {
		var note string
		err := rows.Scan(&note)
		if err != nil {
			log.Fatal(err)
		}
		notes = append(notes, note)
	}

	if rows.Err() != nil {
		return notes, rows.Err()
	}

	return notes, nil
}

func (mysqlDB MySQLDB) SaveNote(note string) error {
	db, err := mysqlDB.createMySQLClient()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Query(fmt.Sprintf("INSERT INTO %s VALUES (%s)", mysqlTable, note))
	return err
}

func (mysqlDB MySQLDB) DeleteNote(note string) error {
	db, err := mysqlDB.createMySQLClient()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Query(fmt.Sprintf("DELETE FROM %s WHERE %s=(%s)", mysqlTable, mysqlNoteColumn, note))
	return err
}

func (mysqlDB MySQLDB) GetHealthStatus() map[string]string {
	//Note implement
	return map[string]string{}
}

func (mysqlDB MySQLDB) RegisterMetrics() {
	//Note implement
}

//NOTE create Test Case and container -> https://hub.docker.com/r/mysql/mysql-server/
// http://go-database-sql.org/accessing.html
