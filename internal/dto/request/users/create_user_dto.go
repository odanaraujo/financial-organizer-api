package users

import (
	"github.com/klassmann/cpfcnpj"
	"time"
)

type User struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	CPF           string    `json:"cpf"`
	BirthDate     time.Time `json:"birth_date"`
	Address       Address   `json:"Address"`
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

func (user User) IsValidCPF() bool {
	user.CPF = cpfcnpj.Clean(user.CPF)
	return cpfcnpj.ValidateCPF(user.CPF)
}
