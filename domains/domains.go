package domains

import (
	"database/sql"

	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func DomainsAddDB(dbname string, domainname string) {
	var id int
	dbloc := "workspaces/" + dbname

	db, err := sql.Open("sqlite3", dbloc+"/"+dbname+".db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id from domains")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(rows)
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&id)
			fmt.Println(id)
		}
	}

	id = id + 1
	fmt.Println(id)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into domains(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, domainname)
	if err != nil {
		log.Println(err)
	}

	tx.Commit()
}
