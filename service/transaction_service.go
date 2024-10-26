package service

import (
	"errors"
	"retail-app/model"
	"retail-app/utils"
)

func CreateTransaction(transaction model.Transaction) error {
	products, err := utils.LoadProducts()
	if err != nil {
		return err
	}

	// Periksa stok barang
	for i, p := range products {
		if p.ID == transaction.ProductID {
			if p.Quantity < transaction.Quantity {
				return errors.New("insufficient stock")
			}
			products[i].Quantity -= transaction.Quantity
			transaction.TotalPrice = float64(transaction.Quantity) * p.Price
			utils.SaveProducts(products)
			break
		}
	}

	// Simpan transaksi
	transactions, err := utils.LoadTransactions()
	if err != nil {
		return err
	}
	transactions = append(transactions, transaction)
	return utils.SaveTransactions(transactions)
}
