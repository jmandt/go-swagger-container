package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmandt/go-swagger-container/cmd/blueprint/daos"
	"github.com/jmandt/go-swagger-container/cmd/blueprint/services"
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
