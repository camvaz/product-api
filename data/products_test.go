package data

import (
	"testing"

	"github.com/camvaz/product-api/models"
)

func TestChecksValidation(t *testing.T){
	p := &models.Product{
		Name: "Vic",
		Price: 1.00,
		SKU: "abs-sda-wda",
	}
	err := p.Validate()
	
	if err != nil {
		t.Fatal(err)
	}
}