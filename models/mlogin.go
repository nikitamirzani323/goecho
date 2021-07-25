package models

import (
	"database/sql"
	"fmt"
	"goecho/db"
	"goecho/helpers"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlstatement := "SELECT * FROM tbl_mst_user WHERE username=?"

	err := con.QueryRow(sqlstatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)

	if !match {
		fmt.Println("Hash and password doesn't match")
		return false, err
	}

	return true, nil
}
