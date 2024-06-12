package users

import (
	"context"
	"github.com/golang/mock/gomock"
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
	"github.com/odanraujo/financial-organizer-api/internal/repository/users/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewUserUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	rp := mocks.NewMockUserRepository(ctrl)
	instance := NewUserUsecase(rp)

	assert.NotNil(t, instance)
}

func TestCreateUser(t *testing.T) {

	ctx := context.Background()

	address := entity.Address{
		Street:  "rua sem nome",
		City:    "Recife",
		State:   "PE",
		ZipCode: "50050202",
		Country: "Brazil",
	}

	user := entity.CreateUser{
		ID:        "1",
		Name:      "test",
		CPF:       "99999999999",
		BirthDate: time.Now(),
		Address:   address,
	}

	type fields struct {
		userRepo func(mocks *mocks.MockUserRepository) *mocks.MockUserRepository
	}

	type args struct {
		ctx  context.Context
		user entity.CreateUser
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Should return create user success",
			fields: fields{
				userRepo: func(mocks *mocks.MockUserRepository) *mocks.MockUserRepository {
					mocks.EXPECT().CreateUser(gomock.AssignableToTypeOf(ctx), user).
						Times(1).
						Return(entity.CreateUser{Name: user.Name}, nil)
					return mocks
				},
			},
			args: args{
				ctx:  ctx,
				user: user,
			},
			wantErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Crie um novo mock para cada iteração do loop
			userRepo := mocks.NewMockUserRepository(ctrl)

			// Configure o comportamento esperado do mock, se aplicável
			if test.fields.userRepo != nil {
				userRepo = test.fields.userRepo(userRepo)
			}

			// Crie uma instância do use case com o mock configurado
			instance := NewUserUsecase(userRepo)

			// Chame o método do use case que está sendo testado
			_, err := instance.CreateUser(test.args.ctx, test.args.user)

			// Verifique se o resultado do teste corresponde ao esperado
			assert.Nil(t, err)
		})
	}
}
