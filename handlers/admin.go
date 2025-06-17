package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminHandler(c *gin.Context) {
	username, _ := c.Cookie("session_user")
	eticketUsername, eticketRoleID := GetETicketData()

	var (
		records []StickerRecord
		err2    error
	)

	doneRecords := make(chan struct{})

	go func() {
		records, err2 = getStickerRecords(DB)
		close(doneRecords)
	}()

	<-doneRecords

	if err2 != nil {
		c.String(500, fmt.Sprintf("Error retrieving data: %v", err2))
		return
	}

	data := gin.H{
		"username": username,
		"records":  records,
		"eticketData": gin.H{
			"username": eticketUsername,
			"roleID":   eticketRoleID,
		},
	}

	c.HTML(http.StatusOK, "admin.html", data)
}

func UpdateStatusHandler(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")
	reason := c.PostForm("reason")

	// Mocking database update for portfolio showcase
	fmt.Printf("Mocking update for ID: %s, Status: %s, Reason: %s\n", id, status, reason)

	// Simulate successful update
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": fmt.Sprintf("Successfully mocked update for ID %s to status %s", id, status),
	})
}
