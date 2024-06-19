package handler

import (
	"TrainingGo/Gin-Crud/entity"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	users  []entity.User
	nextID int = 1
)

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello gin",
	})
}

func PostHandler(c *gin.Context) {
	var json struct {
		Message string `json:"message"`
	}
	if err := c.ShouldBindJSON(&json); err == nil {
		c.JSON(200, gin.H{"message": json.Message})
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}
}

// membuat user baru
func CreateUsers(c *gin.Context) {
	var newUser entity.User
	if err := c.ShouldBindJSON(&newUser); err == nil {
		newUser.ID = int(nextID + 1)
		nextID++
		newUser.CreatedAt = time.Now()
		newUser.UpdatedAt = time.Now()
		users = append(users, newUser)
		c.JSON(http.StatusCreated, users)
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}
}

// get all users
func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// get user by id
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id) //ini untuk convert
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	for _, user := range users {
		if user.ID == userID {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// get user by name
func GetUserByName(c *gin.Context) {
	name := c.Param("name")
	for _, user := range users {
		if user.Name == name {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
