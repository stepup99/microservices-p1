package app

import (
	"fmt"
	"net/http"

	"github.com/stepup99/microservices-p1/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	fmt.Println("starting server port 8080 ............")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
