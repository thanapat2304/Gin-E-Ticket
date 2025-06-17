package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CancelHandler(c *gin.Context) {
	username, _ := c.Cookie("session_user")
	currentPath := c.Request.URL.Path
	eticketUsername, eticketRoleID := GetETicketData()

	var (
		records []StickerRecordCancel
		err2    error
	)

	doneRecords := make(chan struct{})

	go func() {
		records, err2 = getStickerCancel(DB)
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

	c.HTML(http.StatusOK, "cancel.html", data)
}

type StickerRecordCancel struct {
	RecName  string
	Subject  string
	Quantity int
	Details  string
	Status   string
	Time     string
	Note     string
}

func getStickerCancel(db *sql.DB) ([]StickerRecordCancel, error) {
	// Mock data for portfolio showcase
	records := []StickerRecordCancel{
		{
			RecName:  "Mock User C1",
			Subject:  "Mock Subject PQR",
			Quantity: 2,
			Details:  "Mock details for PQR",
			Status:   "Cancel",
			Time:     "2023-04-01",
			Note:     "Cancelled by user",
		},
		{
			RecName:  "Mock User C2",
			Subject:  "Mock Subject STU",
			Quantity: 1,
			Details:  "Mock details for STU",
			Status:   "Cancel",
			Time:     "2023-04-03",
			Note:     "Out of stock",
		},
	}
	return records, nil
}
