package controllers

import (
	"database/sql"
	"fmt"
	"math/big"
)

type user struct{
	firstname string
	lastname string
	email string
	id big.Int
}

func createUser(userData map[string]interface{}){
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+ "password=%s dbname=%s",
		DB_CONFIG["host"], DB_CONFIG["port"], DB_CONFIG["user"], DB_CONFIG["password"], DB_CONFIG["dbname"])

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if val, ok := userData["given_name"]; !ok {
		userData["first_name"] = ""
		fmt.Println(val)
	}
	if val, ok := userData["family_name"]; !ok {
		userData["family_name"] = ""
		fmt.Println(val)
	}
	if val, ok := userData["sub"]; !ok {
		userData["sub"] = ""
		fmt.Println(val)
	}

	if len(userData) != 0{
		sqlStatement := `
			INSERT INTO users (id, firstname, lastname )
			VALUES ('`+ userData["sub"].(string) +`',
					'`+ userData["given_name"].(string)+`',
					'`+ userData["family_name"].(string)+`');`
		_, err = db.Exec(sqlStatement)
		if err != nil {
			panic(err)
		}
	}
}

func getUserData(userData map[string]interface{}) user {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+ "password=%s dbname=%s",
		DB_CONFIG["host"], DB_CONFIG["port"], DB_CONFIG["user"], DB_CONFIG["password"], DB_CONFIG["dbname"])

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	currentUser := user{}

	sqlStatement := `SELECT firstname, lastname FROM users WHERE id=$1;`
	row:= db.QueryRow(sqlStatement, userData["sub"])
	switch err := row.Scan(&currentUser.firstname, &currentUser.lastname); err{
		case sql.ErrNoRows:
			fmt.Print("")
		case nil:
			fmt.Print( currentUser.firstname, currentUser.lastname)
		default:
			panic(err)
	}

	if err != nil {
		panic(err)
	}
	return currentUser

}