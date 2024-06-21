package update

import (
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
	"time"
)

type UserRequest struct {
	Name          *string         `json:"name,omitempty"`
	CPF           *string         `json:"cpf,omitempty"`
	Email         *string         `json:"email,omitempty"`
	BirthDate     *time.Time      `json:"birth_date,omitempty"`
	Address       *AddressRequest `json:"Address,omitempty"`
	CurrentSalary *float64        `json:"current_salary,omitempty"`
}

type AddressRequest struct {
	Street  *string `json:"street,omitempty"`
	City    *string `json:"city,omitempty"`
	State   *string `json:"state,omitempty"`
	ZipCode *string `json:"zip_code,omitempty"`
	Country *string `json:"country,omitempty"`
}

func (request UserRequest) FillModel(user entity.CreateUser) entity.CreateUser {

	if request.Name != nil {
		user.Name = *request.Name
	}

	if request.CPF != nil {
		user.CPF = *request.CPF
	}

	if request.Email != nil {
		user.Email = *request.Email
	}

	if request.BirthDate != nil {
		request.BirthDate = request.BirthDate
	}

	if request.CurrentSalary != nil {
		user.CurrentSalary = *request.CurrentSalary
	}

	if request.Address != nil {
		request.Address.fillModel(user.Address)
	}

	return user
}

func (request AddressRequest) fillModel(address entity.Address) {

	if request.Street != nil {
		address.Street = *request.Street
	}

	if request.State != nil {
		address.State = *request.State
	}

	if request.Country != nil {
		address.State = *request.State
	}

	if request.ZipCode != nil {
		address.ZipCode = *request.ZipCode
	}

	if request.City != nil {
		address.City = *request.City
	}
}
