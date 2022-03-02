package handller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	Router := gin.Default()
	Router.GET("/", Homepage)
	Router.POST("/adduser", AddUser)
	Router.DELETE("/delete/:id", DeleteUser)
	Router.POST("/update", UpdateUser)
	Router.GET("/getuser/:id", GetOneUser)
	Router.GET("/getallusers", GetAllUsers)
	Router.Run(":4040")
}

func Homepage(c *gin.Context) {
	log.Println("intering homepage")
	c.JSON(http.StatusOK, "welcome to homepage")

}
