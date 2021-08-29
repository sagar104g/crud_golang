package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	initSetup "github.com/sagar104g/crud_golang/init"
	model "github.com/sagar104g/crud_golang/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func addBook(c *gin.Context) {
	books := &model.Books{}

	if err := c.ShouldBindJSON(books); err != nil {
		respond(c, http.StatusBadRequest, nil, fmt.Errorf("Internal server error"))
		return
	}

	booksCollection := initSetup.GetCollection("books", "history")

	exp := booksCollection.FindOne(c.Request.Context(), bson.M{"author": books.Author, "title": books.Title})
	if exp.Err() == nil {
		respond(c, http.StatusConflict, nil, fmt.Errorf("Entry already exists"))
		return
	}

	books.ID = primitive.NewObjectID()
	res, err := booksCollection.InsertOne(c.Request.Context(), books)
	if err != nil {
		respond(c, http.StatusInternalServerError, nil, err)
		return
	}

	respond(c, http.StatusCreated, res, nil)
}

func getBook(c *gin.Context) {

	books := []*model.Books{}
	booksCollection := initSetup.GetCollection("books", "history")

	cursor, err := booksCollection.Find(c.Request.Context(), bson.M{})
	if err != nil {
		log.Println(err)
		respond(c, http.StatusInternalServerError, nil, fmt.Errorf("Internal server error"))
		return
	}

	for cursor.Next(c.Request.Context()) {
		book := &model.Books{}
		err = cursor.Decode(book)
		books = append(books, book)
	}

	if err != nil {
		respond(c, http.StatusInternalServerError, nil, fmt.Errorf("Internal server error"))
		return
	}

	respond(c, http.StatusOK, books, nil)
}

func updateBook(c *gin.Context) {
	books := &model.Books{}

	if err := c.ShouldBindJSON(books); err != nil {
		respond(c, http.StatusBadRequest, nil, fmt.Errorf("Internal server error"))
		return
	}

	booksCollection := initSetup.GetCollection("books", "history")
	book := booksCollection.FindOneAndUpdate(c.Request.Context(), bson.M{"author": books.Author, "title": books.Title}, bson.M{"$set": books})
	if book.Err() != nil {
		respond(c, http.StatusInternalServerError, nil, fmt.Errorf("Failed to update book"))
		return
	}

	respond(c, http.StatusOK, books, nil)
}

func deleteBook(c *gin.Context) {

	expMap := make(map[string]string)

	if err := json.NewDecoder(c.Request.Body).Decode(&expMap); err != nil {
		respond(c, http.StatusBadRequest, nil, fmt.Errorf("Failed to update book"))
		return
	}
	if id, ok := expMap["_id"]; !ok {
		respond(c, http.StatusBadRequest, nil, fmt.Errorf("_id cannot be empty"))
		return
	} else {
		booksCollection := initSetup.GetCollection("books", "history")

		dID, _ := primitive.ObjectIDFromHex(id)
		exp := booksCollection.FindOneAndDelete(c.Request.Context(), bson.M{"_id": dID})
		if exp.Err() != nil {
			log.Println(exp.Err())
			respond(c, http.StatusInternalServerError, nil, fmt.Errorf("Failed to delete book"))
			return
		}
	}
	respond(c, http.StatusOK, nil, nil)
}

func lost(c *gin.Context) {
	respond(c, http.StatusNotFound, nil, nil)
}
