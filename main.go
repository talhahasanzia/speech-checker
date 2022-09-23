package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Word struct {
	Name string `json:"name"`
	P1   string `json:"p1"`
	P2   string `json:"p2"`
	P3   string `json:"p3"`
}

var words []Word

func main() {

	words = getWords()

	router := gin.Default()
	router.GET("/pos", getPos)

	router.Run("localhost:8070")
}

func getPos(c *gin.Context) {
	query := c.Query("word")

	fmt.Println("received query:", query)

	found := false
	for _, word := range words {
		if word.Name == query {
			found = true

			fmt.Println("found:", word.Name)

			c.IndentedJSON(http.StatusOK, word)

			break
		}
	}

	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "word not found"})
	}

}
