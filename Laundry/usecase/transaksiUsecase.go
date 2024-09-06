package usecase

import (
	"database/sql"
	"laundry/model"
	"laundry/repository"
)

func GetAllTransaksi(db *sql.DB) ([]model.TransaksiModel, error) {
	return repository.GetAllTransaksi(db)
}

func GetIdTransaksi(db *sql.DB, id int) (*model.TransaksiModel, error) {
	return repository.GetIdTransaksi(db, id)
}

func GetIdCustTransaksi(db *sql.DB, id int) (*model.TransaksiModel, error) {
	_, err := repository.GetIdCustomer(db, id)
	if err != nil {
		return nil, err
	}
	return repository.GetIdCustTransaksi(db, id)
}

func GetIdEmpTransaksi(db *sql.DB, id int) (*model.TransaksiModel, error) {
	_, err := repository.GetIdEmployee(db, id)
	if err != nil {
		return nil, err
	}
	return repository.GetIdEmpTransaksi(db, id)
}

func GetIdSrvTransaksi(db *sql.DB, id int) (*model.TransaksiModel, error) {
	_, err := repository.GetIdService(db, id)
	if err != nil {
		return nil, err
	}
	return repository.GetIdSrvTransaksi(db, id)
}

func InsertTransaksi(db *sql.DB, t model.TransaksiModel) error {
	_, err := repository.GetIdEmployee(db, t.IdEmp)
	if err != nil {
		return err
	}
	_, err = repository.GetIdService(db, t.IdSrv)
	if err != nil {
		return err
	}
	value, err := repository.GetIdService(db, t.IdSrv)
	if err != nil {
		return err
	}
	t.Total = t.Jumlah * value.Price
	return repository.InsertTransaksi(db, t)
}
