package main

import(
	"log"
	"net/http"
	"model/controller"
)

func main() {
	
	// ルートを設定
	http.HandleFunc("/", controller.IndexHandler)
	http.HandleFunc("/increment", controller.IncrementHandler)

	log.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", nil)
	
}