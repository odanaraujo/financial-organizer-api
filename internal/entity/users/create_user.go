package users

import (
	dto "github.com/odanraujo/financial-organizer-api/internal/dto/request/users"
	"time"
)

type CreateUser struct {
	ID            string
	Name          string
	CPF           string
	BirthDate     time.Time
	Address       Address
	CurrentSalary float64
	UsersInvolved []string
}

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
	Country string
}

func NewCreateUser(dto dto.User) CreateUser {
	address := Address{
		Street:  dto.Address.Street,
		City:    dto.Address.City,
		State:   dto.Address.State,
		ZipCode: dto.Address.ZipCode,
		Country: dto.Address.Country,
	}
	return CreateUser{
		ID:            dto.ID,
		Name:          dto.Name,
		CPF:           dto.CPF,
		BirthDate:     dto.BirthDate,
		Address:       address,
		CurrentSalary: dto.CurrentSalary,
		UsersInvolved: dto.UsersInvolved,
	}
}
