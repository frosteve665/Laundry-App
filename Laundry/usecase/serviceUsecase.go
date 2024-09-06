package usecase

import (
	"database/sql"
	"fmt"
	"laundry/model"
	"laundry/repository"
)

func GetAllService(db *sql.DB) ([]model.ServiceModel, error) {
	return repository.GetAllService(db)
}

func GetNamaService(db *sql.DB, nama string) (*model.ServiceModel, error) {
	return repository.GetNamaService(db, nama)
}

func InsertService(db *sql.DB, s model.ServiceModel) error {
	_, err := repository.GetNamaService(db, s.Nama)
	if err == nil {
		return fmt.Errorf("nama sudah ada")
	}
	if len(s.Satuan) > 6 {
		return fmt.Errorf("satuan hanya boleh 6 huruf")
	}
	return repository.InsertService(db, s)
}

func UpdateService(db *sql.DB, s model.ServiceModel) error {
	_, err := repository.GetNamaService(db, s.Nama)
	if err != nil {
		return err
	}
	if len(s.Satuan) > 6 {
		return fmt.Errorf("satuan hanya boleh 6 huruf")
	}
	return repository.UpdateService(db, s)
}

func DeleteService(db *sql.DB, nama string) error {
	_, err := repository.GetNamaService(db, nama)
	if err != nil {
		return err
	}
	return repository.DeleteService(db, nama)
}
