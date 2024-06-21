package users

import (
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
	"time"
)

type User struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	CPF           string  `json:"cpf"`
	Email         string  `json:"email"`
	BirthDate     string  `json:"birth_date"`
	Address       Address `json:"Address"`
	CurrentSalary float64 `json:"current_salary"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

func ConverterDomainToResponse(userDomain entity.CreateUser) *User {
	address := Address{
		Street:  userDomain.Address.Street,
		City:    userDomain.Address.City,
		State:   userDomain.Address.State,
		ZipCode: userDomain.Address.ZipCode,
		Country: userDomain.Address.Country,
	}

	return &User{
		ID:            userDomain.ID,
		Name:          userDomain.Name,
		CPF:           userDomain.CPF,
		Email:         userDomain.Email,
		BirthDate:     userDomain.BirthDate.Format(time.RFC3339),
		Address:       address,
		CurrentSalary: userDomain.CurrentSalary,
	}
}
