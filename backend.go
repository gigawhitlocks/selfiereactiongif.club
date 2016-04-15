package main

import (
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"io/ioutil"
)

func getImage(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte(""))
}

func saveImage(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusOK)
	log.Println(req.Body)
	wr.Write([]byte("success"))
}

func deleteImage(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte(""))
}

func indexHandler(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusOK)
	indexfile, _ := ioutil.ReadFile("static/index.html")
	wr.Write(indexfile)
}

func main() {
	router := pat.New()

	router.Get("/images/{id}", getImage)
	router.Post("/images", saveImage)
	router.Delete("/images/{id}", deleteImage)

	router.Get("/scripts", http.HandlerFunc(http.FileServer(http.Dir("./static/scripts"))))
	router.Get("/", indexHandler)

	http.Handle("/", router)
	log.Print("Listening on 127.0.0.1:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
