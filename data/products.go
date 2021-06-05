package data

import (
	"fmt"
	"time"

	"github.com/camvaz/product-api/models"
)

var ErrProductNotFound = fmt.Errorf("Product not found")

func GetProducts() models.Products {
	return productList
}	

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func findProduct(id int) (*models.Product,int, error) {
	for i, p:= range productList{
		if p.ID == id{
			return p,i,nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func AddProduct(p *models.Product){
	p.ID = getNextID()
	productList = append(productList,p)
}

func UpdateProduct(id int, p *models.Product) error{
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	
	p.ID = id
	productList[pos] = p
	return nil
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
