package websocket

import (
	"context"
	"log"
	"net/http"
	"real-chat/database"
	"real-chat/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	Service *service.UserService
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

/*
func (h *UserHandler) SignUp(c *gin.Context) {
	var ctx, cancel = context.WithTimeot
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validate

	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, Gin.H{"error": validationErr.Error()})
		return
	}

	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	User.Updated_at, _ = time.Parse(time.RCF3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.UserID = user.ID.Hex()

}

*/

func (h *UserHandler) GetUsers(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	//conv string to intigers to access record
	recordsPerPage, err := strconv.Atoi(c.Query("recordsPerPage"))
	if err != nil || recordsPerPage < 1 {
		recordsPerPage = 10 //default 10 record if it has not been inputted
	}
	page, err1 := strconv.Atoi(c.Query("page"))
	if err1 != nil || page < 1 {
		page = 1
	}

	//skip and limit logic
	startIndex := (page - 1) * recordsPerPage
	startIndex, err = strconv.Atoi("startIndex")

	//agg stages, research aggregation stages, dollar sign in project is an operator
	matchStage := bson.D{{Key: "$match", Value: bson.D{}}}
	projectStage := bson.D{
		{Key: "$project", Value: bson.D{ //project allows you to match as many items
			{Key: "_id", Value: 0},
			{Key: "total_count", Value: 1},
			{Key: "user_items", Value: bson.D{{Key: "$slice", Value: []interface{}{"$data", recordsPerPage, startIndex}}}},
		}},
	}

	result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, projectStage})
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
	}

	var allUsers []bson.M
	if err = result.All(ctx, &allUsers); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, allUsers[0])

}

//either err
//or ideally want to return all the users based on the various query param

/*
func (h *UserHandler) GetUsers(*gin.Context){
} ADMIN FEATURE TO CREATE
*/
