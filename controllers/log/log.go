package logcontroller

import (
	"net/http"
	"time"

	"example.com/gin-backend-api/configs"
	"example.com/gin-backend-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertLog(c *gin.Context) {
	var input InputLog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log := models.Log{
		ID:      primitive.NewObjectID(),
		Device:  input.Device,
		LogedAt: time.Now().Local(),
	}

	collection := configs.GetMongoCollection("logs")
	_, err := collection.InsertOne(configs.Ctx, log)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, log)
}

func GetLog(c *gin.Context) {
	cur, err := configs.GetMongoCollection("logs").Find(configs.Ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(configs.Ctx)

	logs := make([]models.Log, 0)
	for cur.Next(configs.Ctx) {
		var log models.Log
		cur.Decode(&log)
		logs = append(logs, log)
	}
	c.JSON(http.StatusOK, logs)
}
