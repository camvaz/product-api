package data

import (
	"time"

	"github.com/camvaz/product-api/models"
)

func GetProducts() models.Products {
	return productList
}

var productList = models.Products{
	&models.Product{
		ID:1,
		Name:"Latte",
		Description: "Frothy milky coffee",
		Price: 2.45,
		SKU: "abc123",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
		DeletedOn: time.Now().UTC().String(),
	},
	&models.Product{
		ID:2,
		Name:"Espresso",
		Description: "Short strong coffee without milk",
		Price: 1.99,
		SKU: "fjd34",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
		DeletedOn: time.Now().UTC().String(),
	},
}
