package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessHandler(c *gin.Context) {
	username, _ := c.Cookie("session_user")
	currentPath := c.Request.URL.Path
	eticketUsername, eticketRoleID := GetETicketData()

	var (
		records []StickerRecordSuccess
		err2    error
	)

	doneRecords := make(chan struct{})

	go func() {
		records, err2 = getStickerSuccess(DB)
		close(doneRecords)
	}()

	<-doneRecords

	if err2 != nil {
		c.String(500, fmt.Sprintf("Error retrieving data: %v", err2))
		return
	}

	data := gin.H{
		"username":    username,
		"currentPath": currentPath,
		"records":     records,
		"eticketData": gin.H{
			"username": eticketUsername,
			"roleID":   eticketRoleID,
		},
	}

	c.HTML(http.StatusOK, "success.html", data)
}

type StickerRecordSuccess struct {
	RecName  string
	Subject  string
	Quantity int
	Details  string
	Status   string
	Time     string
}

func getStickerSuccess(db *sql.DB) ([]StickerRecordSuccess, error) {
	// Mock data for portfolio showcase
	records := []StickerRecordSuccess{
		{
			RecName:  "Mock User 1",
			Subject:  "Mock Subject A",
			Quantity: 10,
			Details:  "Mock details for A",
			Status:   "Complete",
			Time:     "2023-01-15",
		},
		{
			RecName:  "Mock User 2",
			Subject:  "Mock Subject B",
			Quantity: 5,
			Details:  "Mock details for B",
			Status:   "Complete",
			Time:     "2023-01-20",
		},
	}
	return records, nil
}
