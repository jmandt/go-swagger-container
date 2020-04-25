package daos

import (
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/models"
)

// UserDAO persists user data in database
type UserDAO struct{}

// NewUserDAO creates a new UserDAO
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao *UserDAO) Get(id uint) *models.User {
	var user models.User

	// Query Database here...

	user = models.User{
		Model: models.Model{ID: 1},
		FirstName: "Martin",
		LastName: "Heinz",
		Address: "Not gonna tell you",
		Email: "martin7.heinz@gmail.com"}

	return &user
}
