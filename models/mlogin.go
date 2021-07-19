package models

import (
	"database/sql"
	"fmt"
	"goecho/db"
)


type User struct {
	Id       int    `json:"id"`
	username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlstatement := "SELECT * FROM tbl_mst_user WHERE username=?"

	err := con.QueryRow(sqlstatement,username).Scan(
		&obj.Id, &obj.username, $pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}
}