package movies

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Movie struct {
	ID          string `json:"id"`
	Img         string `json:"img"`
	Name        string `json:"name"`
	Year        string `json:"year"`
	Country     string `json:"country"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	Rating      string `json:"rating"`
}

func Router(app *gin.Engine, moviesCollection *mongo.Collection) {

	app.POST("/movie", func(ctx *gin.Context) {
		var movie Movie
		ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, 1048576)
		err := ctx.BindJSON(&movie)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err = moviesCollection.InsertOne(context.TODO(), movie)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Movie created"})
	})
	app.GET("/movie/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		file, err := os.Open("./public/" + id + ".mp4")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()
		ctx.Writer.Header().Set("Content-Type", "video/mp4")
		http.ServeFile(ctx.Writer, ctx.Request, "./public/"+id+".mp4")
	})
}
