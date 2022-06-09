package main

import "net/http"

func Rotas() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/createCar", CreateCar)
	mux.HandleFunc("/showCar", ShowCar)
	mux.HandleFunc("/updateCar", UpdateCar)
	mux.HandleFunc("/deleteCar", DeleteCar)
	return mux
}
