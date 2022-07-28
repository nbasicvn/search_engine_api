package main

import (
	"fmt"
	"github.com/blevesearch/bleve/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	mapping := bleve.NewIndexMapping()
	_, err = bleve.New("index", mapping)
	index, err := bleve.Open("index")
	if err != nil {
		fmt.Println(err)
	}

	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		text := c.Query("text")
		id := c.Query("id")
		err = index.Index(id, text)

		c.JSON(http.StatusOK, gin.H{
			"a": "b",
		})
	})

	r.GET("/remove", func(c *gin.Context) {
		err = index.Delete("2")
		c.JSON(http.StatusOK, gin.H{
			"a": "b",
		})
	})

	r.GET("/reindex", func(c *gin.Context) {
		text := c.Query("text")
		id := c.Query("id")

		err = index.Delete(id)
		c.JSON(http.StatusOK, gin.H{
			"a": "b",
		})

		err = index.Index(id, text)

		c.JSON(http.StatusOK, gin.H{
			"a": "b",
		})
	})

	r.GET("/query", func(c *gin.Context) {
		key := c.Query("key")
		query := bleve.NewMatchQuery(key)
		search := bleve.NewSearchRequest(query)
		searchResults, _ := index.Search(search)
		c.JSON(http.StatusOK, searchResults)
	})

	_ = r.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}
