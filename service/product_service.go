package service

import (
	"errors"
	"retail-app/model"
	"retail-app/utils"
)

func CreateProduct(product model.Product) error {
    products, err := utils.LoadProducts()
    if err != nil {
        return err
    }
    products = append(products, product)
    return utils.SaveProducts(products)
}

func UpdateProductQuantity(id string, quantity int) error {
    products, err := utils.LoadProducts()
    if err != nil {
        return err
    }
    for i, p := range products {
        if p.ID == id {
            products[i].Quantity += quantity
            return utils.SaveProducts(products)
        }
    }
    return errors.New("product not found")
}

// Fungsi tambahan (ReadProduct, DeleteProduct) dapat ditambahkan sesuai kebutuhan.
