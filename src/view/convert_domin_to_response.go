package view

import (
	"github.com/pedro-phd/first-api-go/src/controller/model/response"
	"github.com/pedro-phd/first-api-go/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
