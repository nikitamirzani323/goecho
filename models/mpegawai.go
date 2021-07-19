package models

import (
	"goecho/db"
	"net/http"
)

type Mpegawai struct {
	Id        int    `json:"id"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Telephone string `json:"telephone"`
}

func FetchAll_mpegawai() (Response, error) {
	var obj Mpegawai
	var arrobj []Mpegawai
	var res Response

	con := db.CreateCon()

	sql := "SELECT * FROM tbl_trx_pegawai"

	rows, err := con.Query(sql)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telephone)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func Store_mpegawai(nama string, alamat string, telephone string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlstatement := "INSERT tbl_trx_pegawai (nama,alamat, telephone) VALUES (?,?,?)"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(nama, alamat, telephone)

	if err != nil {
		return res, err
	}

	lastInsertid, err := result.LastInsertId()

	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertid,
	}

	return res, nil
}

func Update_mpegawai(id int, nama string, alamat string, telephone string) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlstatement := "UPDATE tbl_trx_pegawai SET nama=?, alamat=?, telephone=? WHERE id=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telephone, id)

	if err != nil {
		return res, err
	}

	rowsAffeted, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rowsAffeted": rowsAffeted,
	}

	return res, nil
}
func Delete_mpegawai(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlstatement := "DELETE FROM tbl_trx_pegawai WHERE ID = ?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)

	if err != nil {
		return res, err
	}

	rowsAffeted, err := result.RowsAffected()
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rowsAffeted": rowsAffeted,
	}

	return res, nil
}
