package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sut66/team15/controller"
	"github.com/sut66/team15/entity"
	middlewares "github.com/sut66/team15/middleware"
)

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	//Login
	r.POST("/login", controller.Login)
	router := r.Group("")
	{
		router.Use(middlewares.Authorizes()) //wtf "s"
		{
			//Member
			r.POST("/member", controller.CreateMember)
			r.GET("/members", controller.ListMembers)
			r.GET("/member/:id", controller.GetMember)
		}
	}

	// //Room Router
	r.GET("/rooms", controller.Listrooms)
	r.GET("rooms/:id", controller.GetRoom)
	r.POST("/users", controller.CreateRoom)
	r.PATCH("/users", controller.UpdateRoom)
	r.DELETE("/users/:id", controller.DeleteRoom)
	// //Member
	// r.POST("/memberr", controller.CreateMember)
	// r.GET("/members", controller.ListMembers)
	// r.GET("/member/:id", controller.GetMember)
	//request router
		r.POST("/request", controller.CreateBooks_request)
		r.GET("/requests", controller.ListBooks_request)
		r.GET("/request/:id", controller.GetBooks_request)
		r.DELETE("/request/:id", controller.DeleteBooks_request)
		r.PATCH("/request", controller.UpdateBooks_request)
	//requeststatus router
		r.GET("/status", controller.ListBooks_requeststatus)
		r.GET("/status/:id", controller.GetBooks_requeststatus)
	//requeststatus router
		r.GET("/catagory", controller.ListCatagory)
		r.GET("/catagory/:id", controller.GetCatagory)


	//Run the Server
	r.Run()
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
