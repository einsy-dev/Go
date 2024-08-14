package users

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(app *gin.Engine, usersCollection *mongo.Collection) {
	app.GET("/users", func(c *gin.Context) {
		var users []bson.M

		cursor, err := usersCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		cursor.All(context.TODO(), &users)

		c.JSON(200, gin.H{"users": users})
	})
}
