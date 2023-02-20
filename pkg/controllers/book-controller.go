package controllers

import (
	"encoding/json"
	"fmt"
	"myGolangProjects/golang-bookstore/pkg/models"
	"myGolangProjects/golang-bookstore/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBooks(res http.ResponseWriter, _ *http.Request) {
	books :=  models.GetBooks()
	r, err := json.Marshal(books)
	if err != nil {
		fmt.Println("Error while marshaling")
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r)
}

func GetBookById(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 0)
	if err != nil {
		fmt.Println("Error while parsing id")
	}
	bookById, _ := models.GetBookById(ID)	
	r, e := json.Marshal(bookById)
	if e != nil {
		fmt.Println("Error while marshaling")
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r)
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	updatedBook := &models.Book{}
	utils.ParseBody(req, updatedBook)
	params := mux.Vars(req)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookById, db:=models.GetBookById(ID)
	if updatedBook.Name != "" {
		bookById.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		bookById.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		bookById.Publication = updatedBook.Publication
	}
	db.Save(&bookById)
	r, e := json.Marshal(bookById)
	if e != nil {
		fmt.Println("Error while marshaling")
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r)
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(req , newBook)
	newBook.CreateBook()
	r, e := json.Marshal(newBook)
	if e != nil {
		fmt.Println("Error while marshaling")
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r)

}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 0)
	if err != nil {
		fmt.Println("Error while parsing id")
	}
	deleteBook := models.DeleteBook(ID)
	r, e := json.Marshal(deleteBook)
	if e != nil {
		fmt.Println("Error while marshaling")
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r)
}