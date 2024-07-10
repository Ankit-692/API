package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ankit-692/API/router"
)

func main() {

	r:= router.Router()

	fmt.Println("Mongodb API")
	log.Fatal(http.ListenAndServe(":4000",r))
}
