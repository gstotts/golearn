package main

import (
	"net/http"

	"github.com/gstotts/golearn/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
