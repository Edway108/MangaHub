package progress

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddToLibrary(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("user_id")
		mangaID := c.Param("mâng_id")

		_, err := db.Exec(
			`INSERT OR IGNORE INTO user_progress(
			user_id , manga_id , current_chapter , status) VALUES (?,?,?,?)`, userID, mangaID, 1, "reading",
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot add to library"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Manga has been added in user library"})

	}
}
