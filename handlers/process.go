package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProcessHandler(c *gin.Context) {
	username, _ := c.Cookie("session_user")
	currentPath := c.Request.URL.Path
	eticketUsername, eticketRoleID := GetETicketData()

	var (
		records []StickerRecordProcess
		err2    error
	)

	doneRecords := make(chan struct{})

	go func() {
		records, err2 = getStickerProcess(DB)
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

	c.HTML(http.StatusOK, "process.html", data)
}

type StickerRecordProcess struct {
	RecName  string
	Subject  string
	Quantity int
	Details  string
	Status   string
	Time     string
}

func getStickerProcess(db *sql.DB) ([]StickerRecordProcess, error) {
	// Mock data for portfolio showcase
	records := []StickerRecordProcess{
		{
			RecName:  "Mock User P1",
			Subject:  "Mock Subject X",
			Quantity: 8,
			Details:  "Mock details for X",
			Status:   "Process",
			Time:     "2023-02-01",
		},
		{
			RecName:  "Mock User P2",
			Subject:  "Mock Subject Y",
			Quantity: 3,
			Details:  "Mock details for Y",
			Status:   "Process",
			Time:     "2023-02-05",
		},
	}
	return records, nil
}
