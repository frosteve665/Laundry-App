package repository

import (
	"database/sql"
	"fmt"
	"laundry/model"
)

func GetAllTransaksi(db *sql.DB) ([]model.TransaksiModel, error) {
	getAllTransaksi := "SELECT * FROM tr_trasaction"

	Transaksis := []model.TransaksiModel{}

	// Execute querry
	rows, err := db.Query(getAllTransaksi)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var t model.TransaksiModel
		err := rows.Scan(&t.Id, &t.IdCust, &t.IdEmp, &t.IdSrv, &t.Jumlah, &t.Total)
		if err != nil {
			return nil, err
		}
		Transaksis = append(Transaksis, t)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return Transaksis, nil
}

func GetIdTransaksi(db *sql.DB, id int) (*model.TransaksiModel, error) {
	getIdTransaksi := "SELECT * FROM tr_trasaction where id=$1"
	t := &model.TransaksiModel{}

	err := db.QueryRow(getIdTransaksi, id).Scan(&t.Id, &t.IdCust, &t.IdEmp, &t.IdSrv, &t.Jumlah, &t.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("id transaksi tidak di temukan")
		} else {
			return nil, err
		}

	}
	return t, err
}

func GetIdCustTransaksi(db *sql.DB, id int) (*model.TransaksiModel, error) {
	getIdCustTransaksi := "SELECT * FROM tr_trasaction where id_cust=$1"
	t := &model.TransaksiModel{}

	err := db.QueryRow(getIdCustTransaksi, id).Scan(&t.Id, &t.IdCust, &t.IdEmp, &t.IdSrv, &t.Jumlah, &t.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("id customer tidak di temukan")
		} else {
			return nil, err
		}

	}
	return t, err
}

func GetIdEmpTransaksi(db *sql.DB, id int) (*model.TransaksiModel, error) {
	getIdEmpTransaksi := "SELECT * FROM tr_trasaction where id_emp=$1"
	t := &model.TransaksiModel{}

	err := db.QueryRow(getIdEmpTransaksi, id).Scan(&t.Id, &t.IdCust, &t.IdEmp, &t.IdSrv, &t.Jumlah, &t.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("id employee tidak di temukan")
		} else {
			return nil, err
		}

	}
	return t, err
}

func GetIdSrvTransaksi(db *sql.DB, id int) (*model.TransaksiModel, error) {
	getIdSrvTransaksi := "SELECT * FROM tr_trasaction where id_srv=$1"
	t := &model.TransaksiModel{}

	err := db.QueryRow(getIdSrvTransaksi, id).Scan(&t.Id, &t.IdCust, &t.IdEmp, &t.IdSrv, &t.Jumlah, &t.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("id employee tidak di temukan")
		} else {
			return nil, err
		}

	}
	return t, err
}

func InsertTransaksi(db *sql.DB, t model.TransaksiModel) error {
	insertTransaksi := "INSERT INTO tr_trasaction (id_cust, id_emp, id_srv, jumlah, total) VALUES ($1,$2,$3,$4,$5)"

	_, err := db.Exec(insertTransaksi, t.IdCust, t.IdEmp, t.IdSrv, t.Jumlah, t.Total)
	if err != nil {
		return err
	}

	return nil
}
