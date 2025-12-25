package auth

import (
	"MangaHub/pkg/utils"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest

		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		var userID, passwordHash string
		err = db.QueryRow(
			`SELECT id , password_hash FROM user WHERE username=?`, req.Username).Scan(&userID, &passwordHash)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username "})
			return
		}
		err = bcrypt.CompareHashAndPassword(
			[]byte(passwordHash),
			[]byte(req.Password),
		)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
			return
		}
		token, err := utils.GenerateToken(userID, req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"TOKEN": token,
		})

	}

}
