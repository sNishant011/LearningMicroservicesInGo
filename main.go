package part_1

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		data, error := ioutil.ReadAll(r.Body)
		if (error != nil){
			http.Error(rw, "Oppps not found", http.StatusBadRequest)
			return
		}
		fmt.Printf("Data: %s\n", data)
		fmt.Fprintf(rw, "Your data was: %s", data)
	})
	http.HandleFunc("/home", func(rw http.ResponseWriter, r*http.Request){
		data, error := ioutil.ReadAll(r.Body)
		if (error != nil){
			http.Error(rw, "Oppps not found", http.StatusBadRequest)
			return
		}
		fmt.Printf("Data: %s\n", data)
		fmt.Fprintf(rw, "You are at home")
	})

	http.ListenAndServe(":3000", nil)
	
}