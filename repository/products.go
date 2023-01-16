package repository

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
)

type ProductRepository struct {
	db db.DB
}

func NewProductRepository(db db.DB) ProductRepository {
	return ProductRepository{db}
}

func (u *ProductRepository) ReadProducts() ([]model.Product, error) {
	records, Alert := u.db.Load("products")

	var listProducts []model.Product
	Alert = json.Unmarshal([]byte(records), &listProducts)

	if (Alert != nil){
		Alert := model.ErrorResponse{
			Error: "Data Error",
		}
		fmt.Sprintln(Alert)
	}else{
		
	}
	return listProducts, nil // TODO: replace this
}

func (u *ProductRepository) ResetProducts() error {
	err := u.db.Reset("products", []byte("[]"))
	if err != nil {
		return err
	}

	return nil
}
