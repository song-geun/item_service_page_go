package main

import (
	"github.com/rs/cors"
	"item_service_page/src/controller"
	"item_service_page/src/repository5"
	"item_service_page/src/service"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	repository5.Repository1()
	service.Productdata()
	service.Product()

	mux := http.NewServeMux()
	mux.HandleFunc("/Product/list", controller.Controller)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":5000", handler)
}
