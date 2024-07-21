package response

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/utils"
)

type userResponse struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func NewUserResponse(u *models.User) *userResponse {
	r := new(userResponse)
	r.Name = u.Name
	r.Email = u.Email
	r.Token = utils.GenerateJWT(u.ID)
	return r
}
