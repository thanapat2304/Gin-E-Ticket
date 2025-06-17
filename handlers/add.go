package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AddHandler(c *gin.Context) {
	username, _ := c.Cookie("session_user")
	currentPath := c.Request.URL.Path
	eticketUsername, eticketRoleID := GetETicketData()

	successMsg, _ := c.Cookie("success_msg")
	if successMsg != "" {
		c.SetCookie("success_msg", "", -1, "/", "", false, true)
	}

	data := gin.H{
		"username":    username,
		"currentPath": currentPath,
		"names":       []string{"คลังสินค้า-แห้ง", "คลังสินค้า-เย็น"},
		"type":        []string{"สติ๊กเกอร์ - Colorpowder", "สติ๊กเกอร์ - Prenature project", "สติ๊กเกอร์ - Puree", "สติ๊กเกอร์ - Tru-blu", "สติ๊กเกอร์ - Tulip Muffin Cup", "สติ๊กเกอร์ อื่นๆ", "แจ้งปัญหา อื่นๆ"},
		"qtys":        []int{1, 2, 3, 4, 5, 10, 20},
		"successMsg":  successMsg,
		"eticketData": gin.H{
			"username": eticketUsername,
			"roleID":   eticketRoleID,
		},
	}

	c.HTML(http.StatusOK, "add.html", data)
}

func generateDocID(typeName string) (string, error) {
	now := time.Now()
	dateStr := now.Format("20060102")

	typePrefix := "OTHER"
	switch {
	case strings.Contains(typeName, "Colorpowder"):
		typePrefix = "CP"
	case strings.Contains(typeName, "Prenature"):
		typePrefix = "PN"
	case strings.Contains(typeName, "Puree"):
		typePrefix = "PR"
	case strings.Contains(typeName, "Tru-blu"):
		typePrefix = "TB"
	case strings.Contains(typeName, "Tulip"):
		typePrefix = "TM"
	}

	// Mocking count for portfolio showcase
	count := 5 // Simulate 5 existing records for today

	seqNum := count + 1
	seqStr := fmt.Sprintf("%03d", seqNum)

	docID := fmt.Sprintf("%s-%s-%s", dateStr, typePrefix, seqStr)
	return docID, nil
}

func AddPostHandler(c *gin.Context) {
	name := c.PostForm("Name")
	department := c.PostForm("department")
	typeName := c.PostForm("type_name")
	qtysNum := c.PostForm("qtys_num")
	note := c.PostForm("Note")

	ipAddress := c.ClientIP()
	logInfo := fmt.Sprintf("IP: %s", ipAddress)

	quantity, err := strconv.Atoi(qtysNum)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid quantity value")
		return
	}

	docID, err := generateDocID(typeName)
	if err != nil {
		fmt.Printf("Error generating document ID: %v\n", err)
		c.String(http.StatusInternalServerError, "Error generating document ID")
		return
	}

	// Mocking database insertion for portfolio showcase
	fmt.Printf("Mocking insert of DocID: %s, Name: %s, Department: %s, Type: %s, Quantity: %d, Note: %s, LogInfo: %s\n",
		docID, name, department, typeName, quantity, note, logInfo)

	successMsg := fmt.Sprintf("บันทึกข้อมูลสำเร็จ รหัสเอกสาร: %s (Mocked)", docID)
	c.SetCookie("success_msg", successMsg, 5, "/", "", false, true)

	c.Redirect(http.StatusSeeOther, "/add")
}
