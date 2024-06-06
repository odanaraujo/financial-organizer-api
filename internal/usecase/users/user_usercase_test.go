package users

import (
	"context"
	"github.com/golang/mock/gomock"
	dto "github.com/odanraujo/financial-organizer-api/internal/dto/request/users"
	users0 "github.com/odanraujo/financial-organizer-api/internal/dto/response/users"
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

	address := dto.Address{
		Street:  "rua sem nome",
		City:    "Recife",
		State:   "PE",
		ZipCode: "50050202",
		Country: "Brazil",
	}

	user := dto.User{
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
		user dto.User
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
						Return(users0.User{Name: user.Name}, nil)
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

			userRepo := mocks.NewMockUserRepository(ctrl)

			instance := NewUserUsecase(userRepo)

			_, err := instance.CreateUser(test.args.ctx, test.args.user)

			assert.Nil(t, err)
		})
	}
}
