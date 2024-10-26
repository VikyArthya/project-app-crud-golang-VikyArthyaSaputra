package utils

import (
	"encoding/json"
	"io/ioutil"
	"retail-app/model"
)

func LoadProducts() ([]model.Product, error) {
	data, err := ioutil.ReadFile("data/products.json")
	if err != nil {
		return nil, err
	}
	var products []model.Product
	json.Unmarshal(data, &products)
	return products, nil
}

func SaveProducts(products []model.Product) error {
	data, err := json.Marshal(products)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("data/products.json", data, 0644)
}

func LoadTransactions() ([]model.Transaction, error) {
	data, err := ioutil.ReadFile("data/transactions.json")
	if err != nil {
		return nil, err
	}
	var transactions []model.Transaction
	json.Unmarshal(data, &transactions)
	return transactions, nil
}

func SaveTransactions(transactions []model.Transaction) error {
	data, err := json.Marshal(transactions)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("data/transactions.json", data, 0644)
}
