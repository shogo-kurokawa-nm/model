package main

import(
	"log"
	"net/http"
	"github.com/shogo-kurokawa-nm/model/mvc/controller"
)

func main() {
	// ルートを設定
	http.HandleFunc("/", controller.IndexHandler)
	http.HandleFunc("/increment", controller.IncrementHandler)

	log.Println("Starting server at port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}