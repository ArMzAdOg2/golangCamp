package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var baseURL = "/api/"
var dataList = map[int]Movies{}

type Movies struct {
	ID     int    "json:id"
	Title  string "json:title"
	Status string "json:status"
}

func postInsertMovies(c *gin.Context) {
	movie := Movies{}
	c.ShouldBindJSON(&movie)
	id := len(dataList) + 1
	movie.ID = id
	dataList[id] = movie
	c.JSON(201, gin.H{
		"id":     movie.ID,
		"title":  movie.Title,
		"status": movie.Status,
	})
}

func getMovies(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	movie := dataList[id]
	c.JSON(http.StatusOK, gin.H{
		"id":     movie.ID,
		"title":  movie.Title,
		"status": movie.Status,
	})
}

func getListMovies(c *gin.Context) {
	list := []Movies{}
	for _, item := range dataList {
		list = append(list, item)
	}
	c.JSON(http.StatusOK, list)
}

func putUpdateMovies(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	movie := Movies{}
	c.ShouldBindJSON(&movie)
	movie.ID = id
	dataList[id] = movie
	c.JSON(http.StatusOK, gin.H{
		"id":     movie.ID,
		"title":  movie.Title,
		"status": movie.Status,
	})
}

func delRemoveData(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(dataList, id)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func main() {
	api := gin.Default()
	api.POST(baseURL+"todos", postInsertMovies)
	api.GET(baseURL+"todos/:id", getMovies)
	api.GET(baseURL+"todos", getListMovies)
	api.PUT(baseURL+"todos/:id", putUpdateMovies)
	api.DELETE(baseURL+"todos/:id", delRemoveData)
	api.Run(":1234")
}
