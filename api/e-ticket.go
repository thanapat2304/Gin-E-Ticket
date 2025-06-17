package api

import (
	"e-ticket/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ETicketRequest struct {
	Username string `json:"username"`
	RoleID   string `json:"role_id"`
}

func ReceiveETicketHandler(c *gin.Context) {
	var request ETicketRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request format",
		})
		return
	}

	handlers.SetETicketData(request.Username, request.RoleID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data received successfully",
	})
}
