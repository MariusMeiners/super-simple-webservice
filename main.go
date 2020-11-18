package main

import (
	"net/http"
	"super-simple-webservice/controllers"
)


func main() {

controllers.RegisterControllers()
http.ListenAndServe(":3000", nil)

}