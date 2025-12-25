package manga

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Detail(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var mID, mTitle, mAuthor, mGenres, mStatus, mDesc string
		var mTotalchap int

		err := db.QueryRow(
			`SELECT id , title , author , genres , status , total_chapters , description FROM manga WHERE id =?`, id,
		).Scan(&mID, &mTitle, &mAuthor, &mGenres, &mStatus, &mTotalchap, &mDesc)

		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Have no that manga"})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
			return
		}
		c.JSON(200, gin.H{
			"id":             mID,
			"title":          mTitle,
			"author":         mAuthor,
			"genres":         mGenres,
			"status":         mStatus,
			"total_chapters": mTotalchap,
			"description":    mDesc,
		})

	}
}
