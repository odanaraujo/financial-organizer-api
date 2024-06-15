package users

import (
	"time"
)

type User struct {
	Name          string    `json:"name" binding:"required,min=3,max=80"`
	CPF           string    `json:"cpf" binding:"required"`
	Email         string    `json:"email" binding:"required,email"`
	BirthDate     time.Time `json:"birth_date" example:"2020-01-20T23:00:00 -03:00" binding:"required"`
	Address       Address   `json:"Address" binding:"required"`
	CurrentSalary float64   `json:"current_salary"`
	SharedAccount bool      `json:"shared_account"`
	UsersInvolved []string  `json:"users_involved"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}
