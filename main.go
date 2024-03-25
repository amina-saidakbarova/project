package main

import (
	"test/repo/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
    go controller.CreateAdmin()
	r.POST("/signup",controller.SignUp)
	r.POST("/login",controller.Login)
	r.POST("/addService", controller.AddServer)
	r.GET("/getservice", controller.GetServer)
	r.POST("/addRoadmap", controller.AddMap)
	r.GET("/getmap", controller.GetRoadmap)
	r.POST("/addMember", controller.AddMember)
	r.GET("/TeamMember", controller.GetTeam)

	r.Run(":2020")
}