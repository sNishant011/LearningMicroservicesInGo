package main

import (
	"context"
	"log"
	"net/http"
	"learningmicroservicesingo/handlers"
	"os"
	"os/signal"
	"time"
)

func main(){
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gg := handlers.NewGoodbye(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gg)
	s := &http.Server{
		Addr: 				":3000",
		Handler:			sm,
		IdleTimeout: 	120*time.Second,
		ReadTimeout: 	1 *time.Second,
		WriteTimeout: 1*time.Second,
	}
	// go func (go routine) to serve connection simultaneously and independently
	go func(){
		err := s.ListenAndServe()
		if err != nil{
			l.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <- sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)
	// creating timeout context
	// wait for tasks to complete untill 30s and shutdown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// graceful shutdown
	s.Shutdown(tc)
}