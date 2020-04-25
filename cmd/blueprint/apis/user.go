package apis

import (
	"fmt"
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/daos"
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	fmt.Println(c.Param("id"))
	s := services.NewUserService(daos.NewUserDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user := s.Get(uint(id)); user == nil {
		fmt.Println("some error occured on the way. Sorry...")
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(user)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
