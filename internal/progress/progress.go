package progress

import (
	"MangaHub/internal/tcp"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	CurrentChapter int `json:"current_chapter"`
}

func UpdateProgress(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("user_id")
		mangaID := c.Param("manga_id")

		var req UpdateRequest
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request"})
			return
		}
		_, err = db.Exec(
			`UPDATE user_progress
			 SET current_chapter = ?, updated_at = CURRENT_TIMESTAMP
			 WHERE user_id = ? AND manga_id = ?`,
			req.CurrentChapter, userID, mangaID,
		)
		tcp.SendProgressSync(userID, mangaID, req.CurrentChapter)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Cannot update progress",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Progress updated",
		})
	}
}
