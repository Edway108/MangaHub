package progress

import (
	"MangaHub/internal/tcp"
	"MangaHub/internal/udp"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AddToLibrary(db *sql.DB, hub *tcp.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("user_id")
		mangaID := c.Param("manga_id")

		if userID == "" || mangaID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		_, err := db.Exec(`
			INSERT OR IGNORE INTO user_progress
			(user_id, manga_id, current_chapter, status, updated_at)
			VALUES (?, ?, ?, ?, ?)`,
			userID, mangaID, 0, "reading", time.Now(),
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Added to library",
		})
		udp.SendNotification(
			"user " + userID + " added manga " + mangaID,
		)
		fmt.Println("DEBUG user_id =", c.GetString("user_id"))

	}
}
