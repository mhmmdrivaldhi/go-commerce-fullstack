package controllers

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Everyone, Welcome to Go-Commerce!"))
}