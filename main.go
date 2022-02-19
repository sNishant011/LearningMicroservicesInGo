package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	
	// http handleFunc method which takes in two parameters
	// first is the path and second is a method to perform
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		// ioutil.ReadAll method returns two things
		// data and error
		data, error := ioutil.ReadAll(r.Body)
		// if error is not empty
		if (error != nil){
			// Error method of http will print the error messages for us
			// takes in response writer, message and error code
			http.Error(rw, "Oppps not found", http.StatusBadRequest)
			// error method doesn't return anything hence adding return statement to terminate the method
			return
		}
		// printing the data in the body
		fmt.Printf("Data: %s\n", data)
		// Fprintf method takes in writer (ResponseWriter in our case) and response back to the request
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

	// serving in localhost:3000 with default handler
	http.ListenAndServe(":3000", nil)
	
}