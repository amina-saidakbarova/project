package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.POST("/signup",controller.SignUp)
	r.POST("/login",controller.Login)
	r.POST("/addService", controller.AddServer)
	r.POST("/addRoadmap", controller.AddMap)
	r.POST("/addMember", controller.AddMember)
}