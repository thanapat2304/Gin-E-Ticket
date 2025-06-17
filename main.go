package main

import (
	"e-ticket/api"
	"e-ticket/database"
	"e-ticket/handlers"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	// defer database.DB.Close() // Not needed for mock data

	handlers.SetDB(database.DB)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./static")

	r.GET("/", homeHandler)
	r.GET("/logout", handlers.LogoutHandler)
	r.GET("/dashboard", handlers.DashboardHandler)
	r.GET("/success", handlers.SuccessHandler)
	r.GET("/process", handlers.ProcessHandler)
	r.GET("/cancel", handlers.CancelHandler)
	r.GET("/add", handlers.AddHandler)
	r.POST("/add", handlers.AddPostHandler)
	r.GET("/admin", handlers.AdminHandler)
	r.POST("/update-status", handlers.UpdateStatusHandler)
	r.POST("/api/e-ticket", api.ReceiveETicketHandler)

	log.Println("Starting server on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}

func homeHandler(c *gin.Context) {
	userCookie, err := c.Cookie("session_user")
	if err != nil || userCookie == "" {
		eticketUsername, eticketRoleID := handlers.GetETicketData()

		if eticketUsername == "" {
			c.Redirect(http.StatusFound, "/dashboard")
			return
		}

		c.SetCookie("session_user", eticketUsername, 3600, "/", "", false, true)
		c.SetCookie("session_role_id", eticketRoleID, 3600, "/", "", false, true)

		c.Redirect(http.StatusFound, "/dashboard")
		return
	}

	c.Redirect(http.StatusFound, "/dashboard")
}
