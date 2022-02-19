package handlers
// making our custom handler
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
// struct for loggin data
type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
	return &Hello{l}
}
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r* http.Request){
	h.l.Println("Hello")
	data, error := ioutil.ReadAll(r.Body)
	if (error != nil){
		http.Error(rw, "Oppps not found", http.StatusBadRequest)
		return
	}	
	fmt.Printf("Data: %s\n", data)
	fmt.Fprintf(rw, "Your data was: %s", data)
}