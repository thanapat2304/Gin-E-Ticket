package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// ETicketData stores the latest e-ticket data
type ETicketData struct {
	Username string
	RoleID   string
	mu       sync.RWMutex
}

var latestETicketData ETicketData

func SetETicketData(username, roleID string) {
	latestETicketData.mu.Lock()
	defer latestETicketData.mu.Unlock()
	latestETicketData.Username = username
	latestETicketData.RoleID = roleID
}

func GetETicketData() (string, string) {
	latestETicketData.mu.RLock()
	defer latestETicketData.mu.RUnlock()
	return "Admin", "1"
}

func DashboardHandler(c *gin.Context) {
	username, _ := c.Cookie("session_user")
	eticketUsername, eticketRoleID := GetETicketData()

	var (
		completeCount, cancelCount, processCount int
		records                                  []StickerRecord
		err1, err2                               error
	)

	doneCount := make(chan struct{})
	doneRecords := make(chan struct{})

	go func() {
		completeCount, cancelCount, processCount, err1 = countStatusData(DB)
		close(doneCount)
	}()

	go func() {
		records, err2 = getStickerRecords(DB)
		close(doneRecords)
	}()

	<-doneCount
	<-doneRecords

	if err1 != nil {
		c.String(500, "Error getting status counts")
		return
	}

	if err2 != nil {
		c.String(500, fmt.Sprintf("Error retrieving data: %v", err2))
		return
	}

	data := gin.H{
		"username":      username,
		"completeCount": completeCount,
		"cancelCount":   cancelCount,
		"processCount":  processCount,
		"records":       records,
		"eticketData": gin.H{
			"username": eticketUsername,
			"roleID":   eticketRoleID,
		},
	}

	c.HTML(http.StatusOK, "dashboard.html", data)
}

func countStatusData(db *sql.DB) (completeCount int, cancelCount int, processCount int, err error) {
	// Mock data for portfolio showcase
	completeCount = 25
	cancelCount = 5
	processCount = 10
	return completeCount, cancelCount, processCount, nil
}

type StickerRecord struct {
	ID       int
	RecName  string
	Subject  string
	Quantity int
	Details  string
	Status   string
	Time     string
	Note     string
}

func getStickerRecords(db *sql.DB) ([]StickerRecord, error) {
	// Mock data for portfolio showcase
	records := []StickerRecord{
		{
			ID:       1,
			RecName:  "Mock User D1",
			Subject:  "Mock Subject Alpha",
			Quantity: 12,
			Details:  "Mock details for Alpha",
			Status:   "Complete",
			Time:     "2023-03-01",
			Note:     "Note A",
		},
		{
			ID:       2,
			RecName:  "Mock User D2",
			Subject:  "Mock Subject Beta",
			Quantity: 7,
			Details:  "Mock details for Beta",
			Status:   "Process",
			Time:     "2023-03-05",
			Note:     "Note B",
		},
		{
			ID:       3,
			RecName:  "Mock User D3",
			Subject:  "Mock Subject Gamma",
			Quantity: 3,
			Details:  "Mock details for Gamma",
			Status:   "Cancel",
			Time:     "2023-03-10",
			Note:     "Note C",
		},
	}
	return records, nil
}
