package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var (
	DB    *sql.DB
	store *sessions.CookieStore
)

func InitSessionStore(secretKey string) {
	store = sessions.NewCookieStore([]byte(secretKey))
	// ตัวอย่างตั้งค่าความปลอดภัยเพิ่ม
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   30 * 60, // 30 นาที
		HttpOnly: true,
		Secure:   false, // true ถ้าใช้ HTTPS
		SameSite: http.SameSiteLaxMode,
	}
}

func SetDB(database *sql.DB) {
	DB = database
}

func LogoutHandler(c *gin.Context) {
	// Redirect ไปที่ main application
	c.Redirect(http.StatusFound, "http://192.168.2.201:8080")
}
