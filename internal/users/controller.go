package users

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(app *gin.Engine, usersCollection *mongo.Collection) {
	app.GET("/users", func(ctx *gin.Context) {
		var users []bson.M

		cursor, err := usersCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		cursor.All(context.TODO(), &users)

		ctx.JSON(200, gin.H{"users": users})
	})
}
