package controller

import (
	"context"
	"fmt"
	"net/http"
	"test/repo/parameters"
	"test/repo/structs"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp(c *gin.Context) {
	var SignUpTemp structs.SignDanneyen
	c.ShouldBindJSON(&SignUpTemp)

	if SignUpTemp.Name == "" || SignUpTemp.Email == "" || SignUpTemp.Login == "" || SignUpTemp.Password == "" {
		c.JSON(404, "Error")
	} else {
		client, ctx := parameters.DBConnection()

		DBConnect := client.Database("AminaDB").Collection("Users")

		Hashed, _ := parameters.HashPassword(SignUpTemp.Password)

		DBConnect.InsertOne(ctx, bson.M{
			"_id":      primitive.NewObjectID().Hex(),
			"name":     SignUpTemp.Name,
			"email":    SignUpTemp.Email,
			"login":    SignUpTemp.Login,
			"password": Hashed,
			"permission": "Client",
		})
	}
	c.JSON(200, "Success")
}

func Login(c *gin.Context) {
	var LoginTemp structs.SignDanneyen
	c.ShouldBindJSON(&LoginTemp)

	if LoginTemp.Login == "" || LoginTemp.Password == "" {
		c.JSON(404, "Error")
	} else {
		client, ctx := parameters.DBConnection()

		DBConnect := client.Database("AminaDB").Collection("Users")

		result := DBConnect.FindOne(ctx, bson.M{
			"login": LoginTemp.Login,
		})

		var userdata structs.SignDanneyen
		result.Decode(&userdata)
		isValidPass := parameters.CompareHashPasswords(userdata.Password, LoginTemp.Password)
		fmt.Println(isValidPass)

		if isValidPass {
			http.SetCookie(c.Writer, &http.Cookie{
				Name:    "SiteCookie",
				Value:   userdata.Id,
				Expires: time.Now().Add(60 * time.Second),
			})
			c.JSON(200, "success")
		} else {
			c.JSON(404, "Wrong login or password")
		}
	}
}

func CreateAdmin() {
	Client, ctx := parameters.DBConnection()

	createdb := Client.Database("AminaDB").Collection("Users")

	result := createdb.FindOne(ctx, bson.M{
		"permission": "Admin",
	})

	var Signup structs.SignDanneyen
	result.Decode(&Signup)

	if Signup.Id == "" {
		HashedPassword, _ := parameters.HashPassword("amina1234")
		var InsertResult, InsertError = createdb.InsertOne(context.TODO(), bson.M{
			"_id":        primitive.NewObjectID().Hex(),
			"login":      "Amina",
			"password":   HashedPassword,
			"permission": "Admin",
		})

		if InsertError != nil {
			fmt.Printf("InsertError: %v\n", InsertError)
		} else {
			fmt.Printf("InsertResult: %v\n", InsertResult)
		}
	} else {
		fmt.Println("Aram yast")
	}
}

func AddServer(c *gin.Context) {
	var CookieData, CookieError = c.Request.Cookie("SiteCookie")
	fmt.Printf("CookieData: %v\n", CookieData)
	if CookieError != nil {
		c.JSON(404, "ErrorCookie")
	} else {
		client, ctx := parameters.DBConnection()
		DBConnect := client.Database("AminaDB").Collection("Users")

		result := DBConnect.FindOne(ctx, bson.M{
			"_id": CookieData.Value,
		})
		var ServerTemp structs.SignDanneyen
		result.Decode(&ServerTemp)
		if ServerTemp.Permission == "Admin" {
			var DanneyenTemp structs.AddServiceData
			c.ShouldBindJSON(&DanneyenTemp)

			client, _ := parameters.DBConnection()
			DBConnect := client.Database("AminaDB").Collection("AddService")

			var InsertResult, InsertError = DBConnect.InsertOne(context.TODO(), bson.M{
				"_id":          primitive.NewObjectID().Hex(),
				"name":         DanneyenTemp.Name,
				"course_count": DanneyenTemp.CourseCount,
			})

			if InsertError != nil {
				fmt.Printf("InsertError: %v\n", InsertError)
			} else {
				fmt.Printf("InsertResult: %v\n", InsertResult)
				c.JSON(200,"Success")
			}
		} else {
			c.JSON(404, "Error")
			fmt.Printf("ServerTemp: %v\n", ServerTemp)
		}
	}
}
func GetServer(c *gin.Context) {
	var ServeSlice = []structs.AddServiceData{}

	var CookieData, CookieError = c.Request.Cookie("SiteCookie")
	fmt.Printf("CookieData: %v\n", CookieData)
	if CookieError != nil {
		c.JSON(404, "ErrorCookie")
	} else {
		client, ctx := parameters.DBConnection()
		DBConnect := client.Database("AminaDB").Collection("Users")

		result := DBConnect.FindOne(ctx, bson.M{
			"_id": CookieData.Value,
		})
		var ServerTemp structs.SignDanneyen
		result.Decode(&ServerTemp)
		if ServerTemp.Permission == "Client" {
			client, ctx = parameters.DBConnection()
			DBConnectService := client.Database("AminaDB").Collection("AddService")
			result, _ := DBConnectService.Find(ctx, bson.M{})
			for result.Next(ctx) {
				var DBTemp structs.AddServiceData
				result.Decode(&DBTemp)

				ServeSlice = append(ServeSlice, DBTemp)

			}
			c.JSON(200, ServeSlice)
		} else {
			c.JSON(404, "Error")
			fmt.Printf("ServerTemp: %v\n", ServerTemp)
		}
	}
}

func AddMap(c *gin.Context) {
	var CookieData, CookieError = c.Request.Cookie("SiteCookie")
	fmt.Printf("CookieData: %v\n", CookieData)
	if CookieError != nil {
		c.JSON(404, "ErrorCookie")
	} else {
		client, ctx := parameters.DBConnection()
		DBConnect := client.Database("AminaDB").Collection("Users")

		result := DBConnect.FindOne(ctx, bson.M{
			"_id": CookieData.Value,
		})
		var ServerTemp structs.SignDanneyen
		result.Decode(&ServerTemp)
		if ServerTemp.Permission == "Admin" {
			var MapTemp structs.AddMapData
			c.ShouldBindJSON(&MapTemp)

			client, _ := parameters.DBConnection()
			DBConnect := client.Database("AminaDB").Collection("RoadMap")

			var InsertResult, InsertError = DBConnect.InsertOne(context.TODO(), bson.M{
				"_id":         primitive.NewObjectID().Hex(),
				"year":        MapTemp.Year,
				"description": MapTemp.Description,
			})

			if InsertError != nil {
				fmt.Printf("InsertError: %v\n", InsertError)
				c.JSON(200,"Success")
			} else {
				fmt.Printf("InsertResult: %v\n", InsertResult)
			}
		} else {
			c.JSON(404, "Error")
			fmt.Printf("ServerTemp: %v\n", ServerTemp)
		}
	}
}
func GetRoadmap(c *gin.Context) {
	var RoadSlice = []structs.AddMapData{}

	var CookieData, CookieError = c.Request.Cookie("SiteCookie")
	fmt.Printf("CookieData: %v\n", CookieData)
	if CookieError != nil {
		c.JSON(404, "ErrorCookie")
	} else {
		client, ctx := parameters.DBConnection()
		DBConnect := client.Database("AminaDB").Collection("Users")

		result := DBConnect.FindOne(ctx, bson.M{
			"_id": CookieData.Value,
		})
		var MapTemp structs.SignDanneyen
		result.Decode(&MapTemp)
		if MapTemp.Permission == "Client" {
			client, ctx = parameters.DBConnection()
			DBConnectService := client.Database("AminaDB").Collection("RoadMap")
			result, _ := DBConnectService.Find(ctx, bson.M{})
			for result.Next(ctx) {
				var RoadTemp structs.AddMapData
				result.Decode(&RoadTemp)

				RoadSlice = append(RoadSlice, RoadTemp)

			}
			c.JSON(200, RoadSlice)
		} else {
			c.JSON(404, "Error")
			fmt.Printf("MapTemp: %v\n", MapTemp)
		}
	}
}

func AddMember(c *gin.Context) {
	var CookieData, CookieError = c.Request.Cookie("SiteCookie")
	fmt.Printf("CookieData: %v\n", CookieData)
	if CookieError != nil {
		c.JSON(404, "ErrorCookie")
	} else {
		client, ctx := parameters.DBConnection()
		DBConnect := client.Database("AminaDB").Collection("Users")

		result := DBConnect.FindOne(ctx, bson.M{
			"_id": CookieData.Value,
		})
		var ServerTemp structs.SignDanneyen
		result.Decode(&ServerTemp)
		if ServerTemp.Permission == "Admin" {
			var TeamTemp structs.AddMemberData
			c.ShouldBindJSON(&TeamTemp)

			client, _ := parameters.DBConnection()
			DBConnect := client.Database("AminaDB").Collection("TeamMembers")

			var InsertResult, InsertError = DBConnect.InsertOne(context.TODO(), bson.M{
				"_id":      primitive.NewObjectID().Hex(),
				"name":     TeamTemp.Name,
				"position": TeamTemp.Position,
			})

			if InsertError != nil {
				fmt.Printf("InsertError: %v\n", InsertError)
			} else {
				fmt.Printf("InsertResult: %v\n", InsertResult)
				c.JSON(200,"Success")
			}
		} else {
			c.JSON(404, "Error")
			fmt.Printf("ServerTemp: %v\n", ServerTemp)
		}
	}
}
func GetTeam(c *gin.Context) {
	var MemberSlice = []structs.AddMemberData{}

	var CookieData, CookieError = c.Request.Cookie("SiteCookie")
	fmt.Printf("CookieData: %v\n", CookieData)
	if CookieError != nil {
		c.JSON(404, "ErrorCookie")
	} else {
		client, ctx := parameters.DBConnection()
		DBConnect := client.Database("AminaDB").Collection("Users")

		result := DBConnect.FindOne(ctx, bson.M{
			"_id": CookieData.Value,
		})
		var MapTemp structs.SignDanneyen
		result.Decode(&MapTemp)
		if MapTemp.Permission == "Client" {
			client, ctx = parameters.DBConnection()
			DBConnectService := client.Database("AminaDB").Collection("TeamMembers")
			result, _ := DBConnectService.Find(ctx, bson.M{})
			for result.Next(ctx) {
				var TeamTemp structs.AddMemberData
				result.Decode(&TeamTemp)

				MemberSlice = append(MemberSlice, TeamTemp)

			}
			c.JSON(200, MemberSlice)
		} else {
			c.JSON(404, "Error")
			fmt.Printf("MapTemp: %v\n", MapTemp)
		}
	}

}
