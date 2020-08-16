package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)
type webhandle struct {

}
func(webhandle) ServeHTTP(writer http.ResponseWriter, request *http.Request){

	writer.Write([]byte("web1"))
}
type webhandle2 struct {

}
func(webhandle2) ServeHTTP(writer http.ResponseWriter, request *http.Request){

	writer.Write([]byte("web2"))
}

func main(){
	c:=make(chan os.Signal)
	go(func(){

		http.ListenAndServe(":9091",webhandle{})

	})()
	go(func(){
		http.ListenAndServe(":9092",webhandle2{})

	})()
	signal.Notify(c,os.Interrupt)
	s:=<-c
	log.Println(s)
}
