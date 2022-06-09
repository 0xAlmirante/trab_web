package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "poo-avancado-web/pkg/models"
    "poo-avancado-web/pkg/models/mysql"
    "strconv"
)

func CreateCar(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.Write([]byte("CreateCar tem que ser POST"))
        return
    }

    bodyDecoder := json.NewDecoder(r.Body)

    var carro models.Car
    bodyDecoder.Decode(&carro)

    if carro.Marca == "" || carro.Nome == "" || carro.Placa == "" {
        w.Write([]byte("Nome, Marca e Placa do carro precisa ser informado"))
        return
    }

    if carro.Preco <= 0 || carro.Vendido {
        w.Write([]byte("Preco deve ser positivo e a flag Vendido false"))
        return
    }

    id, err := mysql.Insert(carro.Marca, carro.Nome, carro.Placa, carro.Preco, carro.Vendido)

    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    if id > 0 {
        w.Write([]byte(fmt.Sprintf("Carro cadastrado id: %d", id)))
    } else {
        w.Write([]byte("Carro com essa placa ja ta cadastrado"))
    }
}

func ShowCar(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.Write([]byte("ShowCar tem que ser GET"))
        return
    }

    var carros []models.Car
    rows, err := mysql.List()

    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    for rows.Next() {
        var carro models.Car
        rows.Scan(&carro.Id, &carro.Marca, &carro.Nome, &carro.Placa, &carro.Preco, &carro.Vendido)
        carros = append(carros, carro)
    }

    if len(carros) > 0 {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(carros)
    } else {
        w.Write([]byte("Não exista carro cadastrado"))
    }
}