package main

import (
	"fmt"
	"net/http"

	"github.com/Totiruzi/dogma-crud/models"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
)

var ORM orm.Ormer

// init executes before the main function is executed
func init() {
	models.ConnectToDb()
	ORM = models.GetOrmObject()
}

func main() {
	router := gin.Default()
	router.POST("/createUser", createUser)
	router.POST("/readUser", readUser)
	router.POST("/updateUser", updateUser)
	router.DELETE("/deleteUser", deleteUser)
	router.Run(":8080")
}

// createUser creates a user
func createUser(c *gin.Context) {
	var newUser models.Users
	c.BindJSON(&newUser)
	_, err := ORM.Insert(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": "Failed to create the user"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"user_id":   newUser.UserId,
			"user_name": newUser.UserName,
			"email":     newUser.Email,
		})
	}
}

func readUser(c *gin.Context) {
	var user []models.Users
	fmt.Println(ORM)
	_, err := ORM.QueryTable("users").All(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": "Failed to read the users"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "users": &user})
	}
}

func updateUser(c *gin.Context) {
	var updateUser models.Users
	c.BindJSON(&updateUser)
	_, err := ORM.QueryTable("users").Filter("email", updateUser.Email).Update(
		orm.Params{
			"user_name": updateUser.UserName,
			"password":  updateUser.Password,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status": http.StatusInternalServerError, "error": "Failed to update user"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	}
}

func deleteUser(c *gin.Context) {
	var delUser models.Users
	c.BindJSON(&delUser)
	_, err := ORM.QueryTable("users").Filter("email", delUser.Email).Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status": http.StatusInternalServerError,
				"error":  "Failed to delete User"})
	}
}
