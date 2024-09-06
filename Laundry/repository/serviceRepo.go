package repository

import (
	"database/sql"
	"fmt"
	"laundry/model"
)

func GetAllService(db *sql.DB) ([]model.ServiceModel, error) {
	getAllService := "SELECT * FROM mst_Service"

	Services := []model.ServiceModel{}

	// Execute querry
	rows, err := db.Query(getAllService)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var s model.ServiceModel
		err := rows.Scan(&s.Id, &s.Nama, &s.Satuan, &s.Price)
		if err != nil {
			return nil, err
		}
		Services = append(Services, s)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return Services, nil
}

func GetNamaService(db *sql.DB, nama string) (*model.ServiceModel, error) {
	getNamaService := "SELECT * FROM mst_Service where nama=$1"
	Service := &model.ServiceModel{}

	err := db.QueryRow(getNamaService, nama).Scan(&Service.Id, &Service.Nama, &Service.Satuan, &Service.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("nama tidak di temukan")
		} else {
			return nil, err
		}

	}
	return Service, err
}

func GetIdService(db *sql.DB, id int) (*model.ServiceModel, error) {
	getIdService := "SELECT * FROM mst_Service where id=$1"
	Service := &model.ServiceModel{}

	err := db.QueryRow(getIdService, id).Scan(&Service.Id, &Service.Nama, &Service.Satuan, &Service.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("nama tidak di temukan")
		} else {
			return nil, err
		}

	}
	return Service, err
}

func InsertService(db *sql.DB, s model.ServiceModel) error {
	insertService := "INSERT INTO mst_Service (nama, satuan, price) VALUES ($1,$2,$3)"

	_, err := db.Exec(insertService, s.Nama, s.Satuan, s.Price)

	if err != nil {
		return err
	}

	return nil
}

func UpdateService(db *sql.DB, s model.ServiceModel) error {
	UpdateService := "UPDATE mst_Service SET satuan=$2, price=$3 where nama=$1"

	_, err := db.Exec(UpdateService, s.Nama, s.Satuan, s.Price)
	if err != nil {
		return err
	}
	return nil
}

func DeleteService(db *sql.DB, nama string) error {
	deleteService := "DELETE FROM mst_Service where nama=$1"
	_, err := db.Exec(deleteService, nama)
	if err != nil {
		return err
	}
	return nil
}
