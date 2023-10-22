package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create order")
	w.WriteHeader(http.StatusOK)
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List order")
	w.WriteHeader(http.StatusOK)
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get order by ID")
	w.WriteHeader(http.StatusOK)
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update order by ID")
	w.WriteHeader(http.StatusOK)
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete order by ID")
	w.WriteHeader(http.StatusOK)
}
