package main

import (
	"os"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	. "github.com/Rojuinex/go-plugin-sql/dataprovider"
)

type SqlProvider struct {
	db *sql.DB
}

var ProviderImplementation SqlProvider

func getEnvDefault(envName, defaultValue string) string {
	envVar, ok := os.LookupEnv(envName)

	if !ok {
		return defaultValue
	}

	return envVar
}

func (p *SqlProvider) openDB() error {
	MYSQL_USER := getEnvDefault("MYSQL_USER", "root")
	MYSQL_PASS := getEnvDefault("MYSQL_PASS", "")
	MYSQL_HOST := getEnvDefault("MYSQL_HOST", "localhost")
	MYSQL_PORT := getEnvDefault("MYSQL_PORT", "3306")
	MYSQL_DB   := getEnvDefault("MYSQL_DB", "testdb")

	dbString := MYSQL_USER + ":" + MYSQL_PASS + "@tcp(" + MYSQL_HOST + ":" + MYSQL_PORT + ")/" + MYSQL_DB

	db, err := sql.Open("mysql", dbString)

	if err != nil {
		return err
	}

	err = db.Ping()

	if err != nil {
		return err
	}

	p.db = db
	return nil
}

func (p *SqlProvider) GetData() (records []DataType, err error) {
	if p.db == nil {
		err = p.openDB()

		if err != nil {
			return
		}
	}

	rows, err := p.db.Query("SELECT ID, Name FROM testtable")

	if err != nil {
		return
	}
	defer rows.Close()

	var id int32
	var name string

	for rows.Next() {
		err = rows.Scan(&id, &name)

		if err != nil {
			return
		}

		records = append(records, DataType{
			ID: id,
			Name: name,
		})
	}

	err = rows.Err()

	return
}
