package handlers

import (
	"learningmicroservicesingo/product_api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}
func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}
func (p *Products) ServeHTTP(rw http.ResponseWriter, r*http.Request){
	if r.Method == http.MethodGet{
		p.getProducts(rw, r)
		return
	}	else if r.Method == http.MethodPost{
		p.addProduct(rw, r)
		return
	} else if r.Method == http.MethodPut{
		// expects product id
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1{
			http.Error(rw, "Invalid URI", http.StatusNotFound)
			return
		}
		if len(g[0]) != 2{
			http.Error(rw, "Invalid URI", http.StatusNotFound)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil{
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.updateProducts(id, rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
func (p Products) getProducts(rw http.ResponseWriter, h *http.Request){
	// getting the data
	lp := data.GetProducts()
	// parsing struct to jso n
	err := lp.ToJSON(rw)

	if err != nil{
		http.Error(rw, "Unable to parse JSON", http.StatusInternalServerError)
		return
	}
}
func (p*Products) addProduct(rw http.ResponseWriter, r*http.Request){
	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil{
		http.Error(rw, "Unable to parse to JSON", http.StatusBadRequest)
		return
	}
	data.AddProduct(prod)
	p.l.Printf("Prod: %#v", prod)
}
func (p Products) updateProducts(id int, rw http.ResponseWriter, r*http.Request){
	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil{
		http.Error(rw, "Unable to parse json", http.StatusBadRequest)
		return
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound{
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil{
		http.Error(rw, "PNF", http.StatusInternalServerError)
		return
	}

}