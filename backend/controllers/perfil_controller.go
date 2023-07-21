package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/dthlogus/InvGo/configs"
	"github.com/dthlogus/InvGo/models"
	"github.com/dthlogus/InvGo/responses"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var perfilCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func CreatePerfil() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var perfil models.Perfil
		var user models.User
		userId := c.Param("userId")
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(userId)

		if err := c.BindJSON(&perfil); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&perfil); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		newPerfil := models.Perfil{
			Pv:       perfil.Pv,
			Pvp:      perfil.Pvp,
			Dy:       perfil.Dy,
			Lpa:      perfil.Lpa,
			Vpa:      perfil.Vpa,
			Simbolos: perfil.Simbolos,
		}

		user.Perfil = newPerfil

		update := bson.M{"perfil": user.Perfil}
		result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var updateUser models.User
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updateUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updateUser}})
	}
}