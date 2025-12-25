package progress

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LibraryItem struct {
	MangaID        string `json:"manga_id"`
	CurrentChapter int    `json:"current_chapter"`
	Status         string `json:"status"`
}

func GetLibrary(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("user_id")

		rows, err := db.Query(
			`SELECT manga_id, current_chapter, status
			 FROM user_progress
			 WHERE user_id = ?`,
			userID,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Query failed"})
			return
		}
		defer rows.Close()

		var library []LibraryItem
		for rows.Next() {
			var item LibraryItem
			rows.Scan(&item.MangaID, &item.CurrentChapter, &item.Status)
			library = append(library, item)
		}

		c.JSON(http.StatusOK, library)
	}
}
