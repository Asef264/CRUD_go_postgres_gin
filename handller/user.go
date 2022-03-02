package handller

import (
	"log"
	"net/http"
	"sample/constants"
	"sample/models"
	"sample/repository"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Println("error on binding user details")
		panic(err)
	}

	_, err = repository.DBP.Exec(constants.INSERT_USER, user.Name, user.Username)
	if err != nil {
		log.Println("error on adding new user to the table")
		c.JSON(http.StatusInternalServerError, "user did not add")
		panic(err)
	}
	log.Println("user added successfully")
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user added successfully",
		"data":    user,
	})

}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	_, err := repository.DBP.Exec(constants.DELETE_USER, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "some thing went wrong")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "the user successfully deleted",
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "error on binding data")
		return
	}
	_, err = repository.DBP.Exec(constants.UPDATE_USER, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error on some place in exec query")
		return
	}
	c.JSON(http.StatusOK, "the user successfully updated")

}

func GetOneUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	res := repository.DBP.QueryRow(constants.GET_USER, id)
	err := res.Scan(&user.Id, &user.Name, &user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "an error on writing data ")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "you got the user",
		"data":    user,
	})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	var user models.User

	res, err := repository.DBP.Query(constants.GET_USERS)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error on reading the query")
		return
	}
	for res.Next() {
		err := res.Scan(&user.Id, &user.Name, &user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "error on writing data on the user")
			return
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)
}
