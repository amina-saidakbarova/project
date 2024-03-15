package controller

import (
	"context"
	"fmt"
	"test/repo/structs"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var SignUpTemp structs.SignDanneyen
	c.ShouldBindJSON(&SignUpTemp)

	if SignUpTemp.Name == "" || SignUpTemp.Email == "" || SignUpTemp.Login == "" || SignUpTemp.Password == "" {
		c.JSON(404, "Error")
	} else {
		client, ctx := DBConnection()

		DBConnect := client.Database("AminaDBcoin").Collection("Users")
		
		Hashed, _ := HashPassword(SignUpTemp.Password)


		DBConnect.InsertOne(ctx, bson.M{
			"name":     SignUpTemp.Name,
			"email":    SignUpTemp.Email,
			"login":    SignUpTemp.Login,
			"password": Hashed,
		})
	}
}
func Login(c *gin.Context){
	








}
func DBConnection() (*mongo.Client, context.Context) {
	url := options.Client().ApplyURI("mongodb://192.168.43.246:27017")
	NewCtx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	Client, err := mongo.Connect(NewCtx, url)
	if err != nil {
		fmt.Printf("errors: %v\n", err)
	}
	return Client, NewCtx
}
func HashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}
func CompareHashPasswords(HashedPasswordFromDB, PasswordToCampare string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(HashedPasswordFromDB), []byte(PasswordToCampare))
	return err == nil
}
