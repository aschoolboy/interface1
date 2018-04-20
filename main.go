package main

import (
	_ "github.com/go-sql-driver/mysql"
	"interface1/service/deleteOne"
	"interface1/service/query"
	"interface1/service/typeQuery"
	"log"
	"net/http"
)

func main() {
	StartServer()
}

func StartServer() {

	http.HandleFunc("/query", query.Query)
	http.HandleFunc("/typeQuery", typeQuery.TypeQuery)
	http.HandleFunc("/deleteOne", deleteOne.DeleteOne)

	log.Println("Now listening1...")
	http.ListenAndServe(":8080", nil)

}
