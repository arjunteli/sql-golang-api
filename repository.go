package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

func dbConn() (*sql.DB, error) {

	server := ""
	port := 1433
	user := ""
	password := ""
	database := ""

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	// Open the database connection
	db, err := sql.Open("mssql", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func FetchPersonInfo(personID int) (*PersonInfo, error) {
	// Connect to the DB
	db, err := dbConn()
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
		return nil, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// SQL query
	query := `
		SELECT 
			p.name, ph.number, a.city, a.state, a.street1, a.street2, a.zip_code
		FROM 
			person p
		JOIN 
			phone ph ON p.id = ph.person_id
		JOIN 
			address_join aj ON p.id = aj.person_id
		JOIN 
			address a ON aj.address_id = a.id
		WHERE 
			p.id = ?
	`

	var personInfo PersonInfo

	err = db.QueryRow(query, personID).Scan(&personInfo.Name, &personInfo.PhoneNumber, &personInfo.City, &personInfo.State, &personInfo.Street1, &personInfo.Street2, &personInfo.ZipCode)
	if err != nil {
		return nil, err
	}

	return &personInfo, nil
}

func InsertNewPerson(newPerson CreatePersonRequest) error {
	// Connect to the DB
	db, err := dbConn()
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	insertPersonQuery := `INSERT INTO person(name, age) VALUES(?, ?)`
	res, err := tx.Exec(insertPersonQuery, newPerson.Name, newPerson.Age)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	personID, err := res.LastInsertId()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	insertAddressQuery := `INSERT INTO address(city, state, street1, street2, zip_code) VALUES(?, ?, ?, ?, ?)`
	res, err = tx.Exec(insertAddressQuery, newPerson.City, newPerson.State, newPerson.Street1, newPerson.Street2, newPerson.ZipCode)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	addressID, err := res.LastInsertId()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	insertPhoneQuery := `INSERT INTO phone(number, person_id) VALUES(?, ?)`
	_, err = tx.Exec(insertPhoneQuery, newPerson.PhoneNumber, personID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	insertAddressJoinQuery := `INSERT INTO address_join(person_id, address_id) VALUES(?, ?)`
	_, err = tx.Exec(insertAddressJoinQuery, personID, addressID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
