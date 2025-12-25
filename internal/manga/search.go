package manga

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		keyword := c.Query("search")

		rows, err := db.Query(
			`SELECT	id , title , author , genres , status , total_chapters, description FROM manga WHERE title LIKE ? OR genres LIKE ?`, "%"+keyword+"%", "%"+keyword+"%",
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		}
		defer rows.Close()

		var result []gin.H
		for rows.Next() {
			var id, title, author, genres, status, desc string
			var total int
			rows.Scan(&id, &title, &author, &genres, &status, &total, &desc)

			result = append(result, gin.H{
				"id":             id,
				"title":          title,
				"author":         author,
				"genres":         genres,
				"status":         status,
				"total_chapters": total,
				"description":    desc,
			})

		}
		c.JSON(http.StatusOK, result)

	}
}
