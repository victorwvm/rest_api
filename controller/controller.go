package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorwvm/rest_api/models"
)

var films = []models.Film{
	{ID: "1", Title: "Star Wars: Episode III  Revenge of the Sith 2005", ReleaseDate: 2005, Platform: "Disney+", Rating: 5.0},
	{ID: "2", Title: "Before Sunset", ReleaseDate: 2004, Platform: "HBO+", Rating: 4.9},
	{ID: "3", Title: "Past Lives", ReleaseDate: 2023, Platform: "Telecine+", Rating: 5.0},
	{ID: "4", Title: "Aftersun", ReleaseDate: 2022, Platform: "Mubi+", Rating: 4.8},
}

func GetFilms(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, films)
}
