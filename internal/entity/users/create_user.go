package users

import (
	"fmt"
	"github.com/klassmann/cpfcnpj"
	dto "github.com/odanraujo/financial-organizer-api/internal/dto/request/users"
	"regexp"
	"time"
)

const (
	zipCodeRegexPattern = `^\d{5}-?\d{3}$`
	emailRegexPattern   = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

var (
	zipCodeRegex = regexp.MustCompile(zipCodeRegexPattern)
	emailRegex   = regexp.MustCompile(emailRegexPattern)
)

type CreateUser struct {
	ID            string
	Name          string
	CPF           string
	Email         string
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
		Name:          dto.Name,
		CPF:           dto.CPF,
		Email:         dto.Email,
		BirthDate:     dto.BirthDate,
		Address:       address,
		CurrentSalary: dto.CurrentSalary,
		UsersInvolved: dto.UsersInvolved,
	}
}

func (user CreateUser) IsValidCPF() bool {
	user.CPF = cpfcnpj.Clean(user.CPF)
	return cpfcnpj.ValidateCPF(user.CPF)
}

func (user CreateUser) IsValidZipCodeFormat() bool {
	return zipCodeRegex.MatchString(user.Address.ZipCode)
}

func (user CreateUser) IsValidEmail() bool {
	return emailRegex.MatchString(user.Email)
}

func (user CreateUser) FormatCPF(cpf string) string {
	return fmt.Sprintf("%s.%s.%s-%s", cpf[:3], cpf[3:6], cpf[6:9], cpf[9:11])
}
