package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/brunoFpassaglia/go-bookstore/pkg/models"
	"github.com/brunoFpassaglia/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBook()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ShowBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Parse error")
	}
	details, _ := models.ShowBook(ID)
	res, _ := json.Marshal(details)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Parse error")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	vars := mux.Vars(r)
	id := vars["bookId"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("parse error")
	}

	b, db := models.ShowBook(ID)
	if UpdateBook.Name != "" {
		b.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		b.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		b.Publication = UpdateBook.Publication
	}

	db.Save(&b)
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
