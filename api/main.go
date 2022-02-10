package main


import (
	"log"
	"my-driveapp/pkg"
	"net/http"
)


func main() {
	s := pkg.New()
	log.Fatal(http.ListenAndServe(":8080", s.Run()))
}