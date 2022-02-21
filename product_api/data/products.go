package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// product structure
type Product struct{
	// making fields lowercase and omiting some fields when parsing to json by using 
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	SKU string `json:"sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}
// creating custom type which is slice of our Productstruct
type Products []*Product
// using json.NewEncoder instead of marshal
func (p*Products) ToJSON(w io.Writer) error{
	// takes in io writer and returns encoder
	e := json.NewEncoder(w)
	// writes json follwed by newline character
	return e.Encode(p)
}
func (p *Product) FromJson(r io.Reader) error{
	e := json.NewDecoder(r)
	return e.Decode(p)
}
// returning our custom type Products
func GetProducts() Products{
		return productList
}
func AddProduct(p *Product){
	p.ID = getNextID()
	productList = append(productList, p)
}
func UpdateProduct(id int, p*Product) error{
	_, pos, err := findProduct(id)
	if err != nil{
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}


func findProduct(id int) (*Product, int, error){
	for i, p := range productList{
		if (p.ID == id){
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func getNextID() int{
	lp := productList[len(productList) - 1]
	return lp.ID + 1
}
var productList = []*Product{
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky coffee",
		Price: 2.45,
		SKU: "abc323",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},&Product{
		ID: 2,
		Name: "Esspresso",
		Description: "Frothy milky coffee",
		Price: 1.45,
		SKU: "abc323",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}