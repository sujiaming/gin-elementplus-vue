package controller

import (
	"net/http"

	"wanworld/model"

	"github.com/gin-gonic/gin"
)

// func CreateUser(c *gin.Context) {
// 	var user database.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	database.DB.Create(&user)
// 	c.JSON(http.StatusCreated, user)
// }

// func GetUser(c *gin.Context) {
// 	id := c.Param("id")
// 	var user database.User
// 	if err := database.DB.First(&user, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

// func UpdateUser(c *gin.Context) {
// 	id := c.Param("id")
// 	var user database.User
// 	if err := database.DB.First(&user, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	database.DB.Save(&user)
// 	c.JSON(http.StatusOK, user)
// }

//	func DeleteUser(c *gin.Context) {
//		id := c.Param("id")
//		var user database.User
//		if err := database.DB.First(&user, id).Error; err != nil {
//			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
//			return
//		}
//		database.DB.Delete(&user)
//		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
//	}
func GetUsers(c *gin.Context) {
	repo := &model.UserRepository{}
	users, err := repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
