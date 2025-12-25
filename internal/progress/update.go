package progress

import (
	"MangaHub/internal/tcp"
	"MangaHub/internal/udp"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateProgressRequest struct {
	Chapter int `json:"chapter"`
}

func UpdateProgress(db *sql.DB, hub *tcp.Hub, udp *udp.Notifier) gin.HandlerFunc {
	return func(c *gin.Context) {

		userID := c.GetString("user_id")
		mangaID := c.Param("manga_id")
		sessionID := c.GetString("session_id") // middleware set

		var req UpdateProgressRequest
		if err := c.ShouldBindJSON(&req); err != nil || req.Chapter < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter"})
			return
		}

		res, err := db.Exec(`
			UPDATE user_progress
			SET current_chapter = ?, updated_at = ?
			WHERE user_id = ? AND manga_id = ?
		`, req.Chapter, time.Now(), userID, mangaID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
			return
		}

		rows, _ := res.RowsAffected()
		if rows == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Manga not in library"})
			return
		}

		hub.Broadcast(userID, tcp.Message{
			Type:      "progress_update",
			MangaID:   mangaID,
			Chapter:   req.Chapter,
			SessionID: sessionID, // để client khác device mới nhận
		})

		log.Println("[SYNC] progress_update", userID, mangaID, req.Chapter)

		c.JSON(http.StatusOK, gin.H{
			"message": "Progress updated",
		})
		udp.Notify(fmt.Sprintf(
			"USER=%s MANGA=%s CHAPTER=%d",
			userID, mangaID, req.Chapter,
		))

	}
}
