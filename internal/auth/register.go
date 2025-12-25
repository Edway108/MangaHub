package auth

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest

		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Request"})
			return

		}
		if len(req.Username) < 6 || len(req.Password) < 6 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Username or Password too short"})
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to hash the password"})
			return
		}

		//if the password has been successfully hashed
		_, err = db.Exec(
			`INSERT INTO user(id , username , password_hash) VALUES (?,?,?)`, uuid.New().String(), req.Username, string(hash),
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to register or username already exists"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Register success"})

	}
}
