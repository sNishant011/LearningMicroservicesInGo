package handlers

import (
	"fmt"
	"io/ioutil"
	"learningmicroservicesingo/product_api/data"
	"log"
	"net/http"
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
		p.postProducts(rw, r)
		return
	}
	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
func (p*Products) getProducts(rw http.ResponseWriter, h *http.Request){
	// getting the data
	lp := data.GetProducts()
	// parsing struct to jso n
	err := lp.ToJSON(rw)
	if err != nil{
		http.Error(rw, "Unable to parse JSON", http.StatusInternalServerError)
	}
}
func (p*Products) postProducts(rw http.ResponseWriter, r*http.Request){
	d := r.Body
	body, _ := ioutil.ReadAll(d)
	// if err != nil{
	// 	return http.Error(rw, "Not Found", http.StatusBadRequest)
	// }
	fmt.Printf("Data: %s", body)
}