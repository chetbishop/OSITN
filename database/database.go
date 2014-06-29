package database

import (
	"database/sql"
	//"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func CreateDatabase(dbname string) {
	dbloc := "workspaces/" + dbname
	err := os.MkdirAll(dbloc, 0777)
	if err != nil {
		log.Fatal(err)
	}
	DB, err := sql.Open("sqlite3", dbloc+"/"+dbname+".db")
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	sql := `
create table domains (id integer not null primary key, name text unique);
`
	_, err = DB.Exec(sql)
	if err != nil {
		log.Printf("%q: %s\n", err, sql)
		return
	}
}
